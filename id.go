package uid64

import (
	"errors"
	"unsafe"
)

// A UID is a 64 bit (8byte) Unique IDentifier
//   - 6 byte Unix-Milli Timestamp
//   - 1 byte Entropy from /dev/urandom
//   - 1 byte for Counter & GeneratorID (6bit counter & 2bit id)
type UID [8]byte

// Initialize UID from timestamp, random uint8, and genID, and counter.
// For ordinary case, you don't have to call this method,
// call Generater.Gen()/GenDanger() instead.
func InitUID(timestamp int64, entropy, generatorID, counter uint8) (UID, error) {
	if generatorID > 0b11 {
		return UID{}, errors.New("generatorID must be less than 4")
	}
	if counter > 0b0011_1111 {
		return UID{}, errors.New("counter must be less than 64")
	}
	return initUID(timestamp, entropy, generatorID, counter), nil
}

// ToInt return int64 which is interpreted [8]byte as integer.
// This is used when it insert UID into sql DB.
// For ordinary, you don't have to call this method directly.
func (uid UID) ToInt() int64 {
	return *(*int64)(unsafe.Pointer(&uid))
}

// FromInt restores UID from int64.
// This is used when it select UID from sql DB.
func FromInt(i int64) UID {
	return UID(*(*[8]byte)(unsafe.Pointer(&i)))
}

// String retrun Base36 representation.
func (uid UID) String() string {
	return toBase36(uid)
}

// Parse restores UID from Base36 string.
func Parse(str string) (UID, error) {
	return fromBase36(str)
}

// Timestamp returns 32 bit timestamp field value as int64, same to time.Unix().Milli.
func (uid UID) Timestamp() int64 {
	return int64(bytesToUint(uid[:6]))
}

// Entropy returns 8bit random field value.
func (uid UID) Entropy() uint8 {
	return uint8(bytesToUint(uid[6:7]))
}

// Counter returns 6bit counter field as uint8.
func (uid UID) Counter() uint8 {
	return uint8(bytesToUint(uid[7:])) >> 2

}

// GeneratorID returns 2bit generator-id field as unint8
func (uid UID) GeneratorID() uint8 {
	return uint8(bytesToUint(uid[7:])) & 0b11
}

func bytesToUint(buf []byte) uint64 {
	l := len(buf) - 1
	var u64 uint64 = 0
	for i, b := range buf {
		u64 += uint64(b) << ((l - i) * 8)
	}
	return u64
}

func (uid *UID) setTimestamp(timestamp int64) {
	uid[0] = byte(timestamp >> 40)
	uid[1] = byte(timestamp >> 32)
	uid[2] = byte(timestamp >> 24)
	uid[3] = byte(timestamp >> 16)
	uid[4] = byte(timestamp >> 8)
	uid[5] = byte(timestamp)
}

func (uid *UID) setEntropy(ent uint8) {
	uid[6] = ent
}

func (uid *UID) setCounterAndGenID(cnt, genid uint8) {
	uid[7] = (cnt << 2) + (genid & 0b11)
}

func initUID(timestamp int64, entropy, generatorID, counter uint8) UID {
	uid := UID{}
	(&uid).setTimestamp(timestamp)
	(&uid).setEntropy(entropy)
	(&uid).setCounterAndGenID(counter, generatorID)
	return uid
}

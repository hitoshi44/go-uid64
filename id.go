package uid64

import "unsafe"

// A UID is a 64 bit (8byte) Unique IDentifier
//   - 5 byte Unix Timestamp
//   - 1 byte Entropy from /dev/urandom
//   - 1 byte for counter per second
//   - 1 byte for Generator IDentifier
type UID [8]byte

func InitUID(timestamp int64, entropy, counter, generatorID uint8) UID {
	uid := UID{}
	(&uid).setTimestamp(timestamp)
	(&uid).setEntropy(entropy)
	(&uid).setCounter(counter)
	(&uid).setGeneratorID(generatorID)
	return uid
}

func (uid UID) ToInt() int64 {
	return *(*int64)(unsafe.Pointer(&uid))
}

func FromInt(i int64) UID {
	return UID(*(*[8]byte)(unsafe.Pointer(&i)))
}

func (uid UID) String() string {
	return toBase36(uid)
}

func Parse(str string) (UID, error) {
	return fromBase36(str)
}

// Timestamp returns 32 bit timestamp field value as int64, same to time.Unix().
func (uid UID) Timestamp() int64 {
	return int64(bytesToUint(uid[:5]))
}

func (uid UID) Entropy() uint8 {
	return uint8(bytesToUint(uid[5:6]))
}

func (uid UID) Counter() uint8 {
	return uint8(bytesToUint(uid[6:7]))

}

func (uid UID) GeneratorID() uint8 {
	return uint8(bytesToUint(uid[7:]))
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
	uid[0] = byte(timestamp >> 32)
	uid[1] = byte(timestamp >> 24)
	uid[2] = byte(timestamp >> 16)
	uid[3] = byte(timestamp >> 8)
	uid[4] = byte(timestamp)
}

func (uid *UID) setEntropy(ent uint8) {
	uid[5] = ent
}

func (uid *UID) setCounter(cnt uint8) {
	uid[6] = cnt
}

func (uid *UID) setGeneratorID(genid uint8) {
	uid[7] = genid
}

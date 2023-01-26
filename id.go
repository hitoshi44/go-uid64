package uid64

import (
	"errors"
)

const MaxTimestamp int64 = 0x8ffffff_ffffffff

// A UID is a 64 bit (8byte) Unique IDentifier
//   - 6 byte Unix-Milli Timestamp
//   - 1 byte Entropy from /dev/urandom
//   - 1 byte for Counter & GeneratorID (6bit counter & 2bit id)
type UID uint64

// Initialize UID from timestamp, random uint8, and genID, and counter.
// For ordinary case, you don't have to call this method,
// call Generater.Gen()/GenDanger() instead.
func InitUID(timestamp int64, entropy, generatorID, counter uint8) (UID, error) {
	if timestamp > MaxTimestamp {
		return 0, errors.New("timestamp should be less than 0x0fffffff_ffffffff")
	}
	if generatorID > 0b11 {
		return 0, errors.New("generatorID must be less than 4")
	}
	if counter > 0b0011_1111 {
		return 0, errors.New("counter must be less than 64")
	}
	return initUID(timestamp, entropy, generatorID, counter), nil
}

// String retrun Base36 string representation.
func (uid UID) String() string {
	return toBase36(uid)
}

// Parse restores UID from Base36 string.
func Parse(str string) (UID, error) {
	u, err := fromBase36(str)
	return UID(u), err
}

// ToInt return int64 (casting from uint64).
// This is used when it insert UID into sql DB.
// For ordinary, you don't have to call this method directly.
func (uid UID) ToInt() int64 {
	return int64(uid)
}

// FromInt restores UID from int64.
// This is used when it select UID from sql DB.
func FromInt(i int64) (UID, error) {
	// TODO check timestamp range
	return UID(i), nil
}

// Timestamp returns 48 bit timestamp field value as int64, same to time.Unix().Milli.
func (uid UID) Timestamp() int64 {
	return int64((uid & 0xffffffff_ffff0000) >> 16)
}

// Entropy returns 8bit random field value.
func (uid UID) Entropy() uint8 {
	return uint8((uid & 0xff00) >> 8)
}

// Counter returns 6bit counter field.
func (uid UID) Counter() uint8 {
	return uint8(uid & 0b0011_1111)
}

// GeneratorID returns 2bit generator-id field,
func (uid UID) GeneratorID() uint8 {
	return uint8(uid&0b1100_0000) >> 6
}

func initUID(timestamp int64, entropy, generatorID, counter uint8) UID {
	var uid uint64 = 0
	uid += (uint64(timestamp) << 16)
	uid += (uint64(entropy) << 8)
	uid += uint64(generatorID << 6)
	uid += uint64(counter)
	return UID(uid)
}

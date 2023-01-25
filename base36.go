package uid64

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// This file is for base36 enc/dec stuff.
// For now, it just use golang std's facility to do that.
// Improvement is needed.

// TODO: refactor to more efficient way.
func toBase36(uid UID) string {
	buf := make([]byte, 13)
	copy(buf, "0000000000000")
	// encode into base36 with uint repr.
	i := uid.ToInt()
	u := *(*uint64)(unsafe.Pointer(&i))
	b36 := strings.ToUpper(strconv.FormatUint(u, 36))
	// padding 0
	l := len(b36)
	copy(buf[13-l:], b36)

	return string(buf)
}

func fromBase36(str string) (UID, error) {
	if len(str) != 13 {
		return UID{}, fmt.Errorf("UID string must be 13 digts: %s len(%d)", str, len(str))
	}
	u, err := strconv.ParseUint(str, 36, 64)
	if err != nil {
		return UID{}, err
	}
	bytes := *(*[8]byte)(unsafe.Pointer(&u))
	return UID(bytes), nil
}

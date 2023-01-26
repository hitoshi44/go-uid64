package uid64

import (
	"database/sql/driver"
	"fmt"
)

// Implementations some interfaces for your convinience.
//  - sql.Valuer, sql.Scanner
//  - sort.Interface

// Implementation for sql.driver interfaces.
func (uid UID) Value() (driver.Value, error) {
	return uid.ToInt(), nil
}
func (uid *UID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		// followed github.com/google/uuid
		return nil
	case int64:
		res, err := FromInt(src)
		if err != nil {
			return err
		}
		*uid = res
	case string:
		parsed, err := Parse(src)
		if err != nil {
			return err
		}
		*uid = parsed
	default:
		return fmt.Errorf("Scan: unable to scan type %T into UID64", src)
	}
	return nil
}

// Implementation sort.Interface for convinience
type UID64Slice []UID

func (x UID64Slice) Len() int           { return len(x) }
func (x UID64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x UID64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

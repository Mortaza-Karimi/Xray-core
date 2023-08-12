package task

import "github.com/Mortaza-Karimi/Xray-core/blob/main/common"

// Close returns a func() that closes v.
func Close(v interface{}) func() error {
	return func() error {
		return common.Close(v)
	}
}

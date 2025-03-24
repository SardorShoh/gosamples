package file

import "os"

func Exists(path string) bool {
	if _, err := os.Lstat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

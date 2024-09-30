package pwd

import "os"

func Pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return wd
}

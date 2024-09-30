package ls

import (
	"CmdTool/pwd"
	"os"
)

func List() (files []string) {
	wd := pwd.Pwd()
	dirlist, err := os.ReadDir(wd)
	if err != nil {
		return []string{}
	}
	for _, file := range dirlist {
		if file.IsDir() {
			files = append(files, file.Name()+"/")
		} else {
			files = append(files, file.Name())
		}

	}
	return
}

package cat

import (
	"CmdTool/ls"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	mp map[string]int // 1表示是file  2表示是dir
)

func init() {
	mp = make(map[string]int)
	fileLists := ls.List()
	for _, file := range fileLists {
		if strings.HasSuffix(file, "/") {
			mp[strings.TrimSuffix(file, "/")] = 2
		} else {
			mp[file] = 1
		}
	}
	// fmt.Println("---------", mp)  // map[cat:2 go.mod:1 helper:2 ls:2 main.go:1 mkdir:2 pwd:2 timer:2 touch:2]
}

func Cat(writer *bufio.Writer, file string) (err string) { // 展开file的内容并逐行打印

	v, ok := mp[file]
	if !ok { // 如果不存在file的话，返回[ERROR] cat: %s: No such file or directory
		return fmt.Sprintf("[ERROR] cat: %s: No such file or directory", file)
	}
	if v == 2 { // 如果此file是dir的话，返回错误信息  [ERROR] cat: %s/: Is a directory
		return fmt.Sprintf("[ERROR] cat: %s/: Is a directory", file)
	}

	// 具体按行展开file中的内容
	content, errInfo := os.ReadFile(file)
	if errInfo != nil {
		return fmt.Sprintf("[ERROR] cat: %s/: cannot read this file", file)
	}
	fmt.Fprint(writer, string(content))

	return ""
}

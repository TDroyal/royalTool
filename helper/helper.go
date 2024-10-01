package helper

import (
	"bufio"
	"fmt"
)

type Helper struct {
	name  string //命令的名称 -v or version
	tips  string // 命令的作用 get version
	usage string //命令的用法 royal -v or royal version
}

var (
	helpSlice []*Helper
)

func Writer(writer *bufio.Writer) {
	// 注册一个
	helpSlice = append(helpSlice, &Helper{
		name:  "-h (help)",
		tips:  "how to use royal tool?",
		usage: "royal -h or royal help",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "-v (version)",
		tips:  "get version.",
		usage: "royal -v or royal version",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "pwd",
		tips:  "display the name of the current working directory with the pwd command.",
		usage: "royal pwd",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "time",
		tips:  "get the current time.",
		usage: "royal time",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "ls",
		tips:  "list the files in the current directory.",
		usage: "royal ls",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "mkdir",
		tips:  "create one or more new directories under the current directory.",
		usage: "royal mkdir dirname [dirname] [...]",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "touch",
		tips:  "create one or more new files under the current directory.",
		usage: "royal touch filename [filename] [...]",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "cat",
		tips:  "view the contents of a file.",
		usage: "royal cat filename",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "rm",
		tips:  "remove files or directories.",
		usage: "royal rm filename [filename] [...]  or royal rm -r dirname [dirname] [...]",
	})

	helpSlice = append(helpSlice, &Helper{
		name:  "find",
		tips:  "find all files and directories under the current directory.",
		usage: "royal find filename",
	})

	// defer writer.Flush()
	fmt.Fprintln(writer, "Usage of royal tool:")
	for _, v := range helpSlice {
		fmt.Fprintf(writer, "\t%s\n", v.name)
		fmt.Fprintf(writer, "\t\ttips: %s\n", v.tips)
		fmt.Fprintf(writer, "\t\tusage: %s\n", v.usage)
	}
}

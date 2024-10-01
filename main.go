package main

import (
	"CmdTool/cat"
	"CmdTool/find"
	"CmdTool/helper"
	"CmdTool/ls"
	"CmdTool/mkdir"
	"CmdTool/pwd"
	"CmdTool/rm"
	"CmdTool/timer"
	"CmdTool/touch"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 一些基本信息
var (
	writer  *bufio.Writer
	version string   = "royal version royal1.0 windows/amd64"
	oneArgs []string = []string{"-v", "version", "-h", "help", "pwd", "time", "ls"}
	mp      map[string]bool
)

func init() {
	writer = bufio.NewWriter(os.Stdout)
	mp = make(map[string]bool)
	for _, v := range oneArgs {
		mp[v] = true
	}
}

func main() {
	args := os.Args[1:]
	defer writer.Flush()

	length := len(args)
	if length < 1 {
		helper.Writer(writer)
		return
	}

	fisrtArg := args[0]

	switch {
	case length == 1 && mp[fisrtArg]: // 单个参数命令
		switch fisrtArg {
		case "-h", "help":
			helper.Writer(writer)
		case "-v", "version":
			fmt.Fprintln(writer, version)
		case "pwd":
			fmt.Fprintln(writer, pwd.Pwd())
		case "time":
			fmt.Fprintln(writer, timer.GetTime())
		case "ls":
			files := ls.List()
			for _, f := range files {
				fmt.Fprintf(writer, "%s   ", f)
			}
		}
	case length > 1 && !mp[fisrtArg]: // 多参数命令
		switch fisrtArg {
		case "mkdir":
			errInfo := mkdir.Mkdir(args[1:])
			for _, e := range errInfo {
				fmt.Fprintln(writer, e)
			}
		case "touch":
			errInfo := touch.Mkfile(args[1:])
			for _, e := range errInfo {
				fmt.Fprintln(writer, e)
			}
		case "cat":
			switch {
			case length == 2: // cat只能一次查看一个文件
				errInfo := cat.Cat(writer, args[1])
				fmt.Fprintln(writer, errInfo)
			default:
				fmt.Fprintf(writer, "[ERROR] syntax error: royal %s\n", strings.Join(args, " "))
				fmt.Fprintln(writer, "Try 'royal -h' for more information.")
			}
		case "rm":
			switch {
			case args[1] != "-r": // rm filename filename filename ...
				errInfo := rm.Rm(1, args[1:])
				for _, e := range errInfo {
					fmt.Fprintln(writer, e)
				}
			case length > 2 && args[1] == "-r": // rm -r dirname ... or filename ...
				errInfo := rm.Rm(2, args[2:])
				for _, e := range errInfo {
					fmt.Fprintln(writer, e)
				}
			default:
				fmt.Fprintf(writer, "[ERROR] syntax error: royal %s\n", strings.Join(args, " "))
				fmt.Fprintln(writer, "Try 'royal -h' for more information.")
			}
		case "find": // find filename
			switch {
			case length == 2: // find只能一次查找一个文件   打印找到的所有文件
				spendTime, result := find.Find(args[1])
				for _, r := range result {
					fmt.Fprintln(writer, r)
				}
				fmt.Fprintf(writer, "find %d files total\n", len(result))
				fmt.Fprintf(writer, "cost time: %v (so fast, hahaha)\n", spendTime)
			default:
				fmt.Fprintf(writer, "[ERROR] syntax error: royal %s\n", strings.Join(args, " "))
				fmt.Fprintln(writer, "Try 'royal -h' for more information.")
			}
		}
	default:
		fmt.Fprintf(writer, "[ERROR] syntax error: royal %s\n", strings.Join(args, " "))
		fmt.Fprintln(writer, "Try 'royal -h' for more information.")
	}

}

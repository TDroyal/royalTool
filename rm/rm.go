package rm

import (
	"CmdTool/ls"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	mp map[string]int // 1表示是file  2表示是dir
	wg sync.WaitGroup
	m  sync.Mutex
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
}

func Rm(op int, files []string) (err []string) { // 包并发的删
	// 删的过程中出现了错误，也要返回
	if op == 1 { // 删文件file    files中出现的文件夹dir报错
		for _, file := range files {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				v, ok := mp[filename]
				if !ok { // 文件不存在
					m.Lock()
					err = append(err, fmt.Sprintf("[ERROR] rm: '%s': No such file or directory", filename))
					m.Unlock()
					return
				}

				if v == 2 { // 是dir
					m.Lock()
					err = append(err, fmt.Sprintf("[ERROR] rm: cannot remove '%s/': Is a directory", filename))
					m.Unlock()
					return
				}
				// 是存在的文件
				if errinfo := os.Remove(filename); errinfo != nil {
					m.Lock()
					err = append(err, fmt.Sprintf("[ERROR] rm: cannot remove '%s/'", filename)) //
					m.Unlock()
				}

			}(file)
		}
		wg.Wait()
	} else { // 删文件夹dir    无脑删完
		for _, file := range files {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()
				_, ok := mp[filename]
				if !ok { // 文件不存在
					m.Lock()
					err = append(err, fmt.Sprintf("[ERROR] rm: '%s': No such file or directory", filename))
					m.Unlock()
					return
				}

				// 是存在的文件 or dir
				if errinfo := os.RemoveAll(filename); errinfo != nil {
					m.Lock()
					err = append(err, fmt.Sprintf("[ERROR] rm: cannot remove '%s/'", filename)) //
					m.Unlock()
				}

			}(file)
		}
		wg.Wait()
	}

	return
}

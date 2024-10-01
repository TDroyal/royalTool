package find

import (
	"CmdTool/pwd"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	dfs            func(path string, isAWorker bool)
	m              sync.Mutex
	rwm            sync.RWMutex
	maxWorkerCount int = 8 // 最大goroutine数量
	workerCount    int = 1
	Done               = make(chan bool)   // 工人是否搜索结束
	searchRequest      = make(chan string) // 新的搜索任务让新的worker做
)

func consumer() {
	for {
		select {
		case <-Done: // 一次只能接收一个
			wCount := 0
			rwm.Lock()
			workerCount--
			wCount = workerCount
			rwm.Unlock()
			if wCount == 0 {
				return
			}
		case path := <-searchRequest:
			rwm.Lock()
			workerCount++
			rwm.Unlock()
			go dfs(path, true)
		}
	}
}

func Find(findFile string) (spendTime time.Duration, result []string) { // find会报错吗？不会吧  goroutine来做，返回所有的查到的文件
	rootPath := pwd.Pwd()
	dfs = func(path string, isAWorker bool) {
		files, err := os.ReadDir(path)
		if err != nil { // 多半没权限  不可读
			if isAWorker {
				Done <- true
			}
			return
		}

		for _, file := range files {
			fileName := file.Name()
			nextPath := fmt.Sprintf("%s\\%s", path, fileName)
			if fileName == findFile {
				m.Lock()
				result = append(result, nextPath)
				m.Unlock()
			}
			if file.IsDir() { // 还有剩余的worker，就派worker来做
				// dfs(nextPath)
				// 先抢锁来读
				wCount := 0
				rwm.RLock() // 用读锁，工人可能会超过8个，就这样吧，不想超过8个，改成写锁就可以了
				wCount = workerCount
				rwm.RUnlock()
				if wCount < maxWorkerCount {
					searchRequest <- nextPath
				} else {
					dfs(nextPath, false)
				}
			}
		}

		if isAWorker {
			Done <- true
		}
	}

	start := time.Now()
	go dfs(rootPath, true)
	consumer()
	spendTime = time.Since(start)
	return
}

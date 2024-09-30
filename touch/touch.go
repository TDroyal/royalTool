package touch

import (
	"fmt"
	"os"
	"sync"
)

var (
	wg sync.WaitGroup
	m  sync.Mutex
)

func Mkfile(files []string) (errInfo []string) { // 返回错误信息
	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
			if err != nil {
				m.Lock()
				errInfo = append(errInfo, fmt.Sprintf("[ERROR] touch: create %s file error", filename))
				m.Unlock()
			} else {
				f.Close()
			}

		}(file)
	}
	wg.Wait()
	return
}

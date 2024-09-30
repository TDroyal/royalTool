package mkdir

import (
	"fmt"
	"os"
	"sync"
)

var (
	wg sync.WaitGroup
	m  sync.Mutex
)

func Mkdir(dirs []string) (errInfo []string) { // 包开协程来做的
	// fmt.Println(dirs)
	for _, dir := range dirs {
		wg.Add(1)
		go func(dirname string) {
			defer wg.Done()
			if err := os.Mkdir(dirname, 0755); err != nil {
				m.Lock()
				errInfo = append(errInfo, fmt.Sprintf("[ERROR] mkdir: create %s directory error", dirname))
				m.Unlock()
			}

		}(dir)
	}
	wg.Wait()
	return
}

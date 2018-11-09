/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/seeleteam/go-seele/cmd/node/cmd"
	"github.com/seeleteam/go-seele/common"
)

func main() {
	cmd.Execute()
}

func init() {
	go monitor()
}

func monitor() {
	size := uint64(1024 * 1024 * 1024 * 3)
	var info runtime.MemStats
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			runtime.ReadMemStats(&info)
			if info.Alloc > size {
				file := filepath.Join(common.GetTempFolder(), fmt.Sprint("heap-", time.Now().Format("2006-01-02-15-04-05")))
				f, err := os.Create(file)
				if err != nil {
					fmt.Println("monitor create file err:", err)
					return
				}
				pprof.WriteHeapProfile(f)
			}
		}
	}

}

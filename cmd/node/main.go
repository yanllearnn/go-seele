/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package main

import (
	"fmt"
	"math"
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
	go check()
}

func check() {
	tag := uint64(math.Pow(2, 30)) * 3
	for {
		timer := time.NewTimer(10 * time.Second)
		<-timer.C
		var memInfo runtime.MemStats
		runtime.ReadMemStats(&memInfo)
		if memInfo.Alloc > tag {
			name := time.Now().Format("2006-01-02-15.04.05.999999999")
			filename := fmt.Sprint("heap-", name, ".heap")
			f, err := os.Create(filepath.Join(common.GetDefaultDataFolder(), filename))
			if err != nil {
				fmt.Println("Failed to create heap file for test, %s", err)
				return
			}

			pprof.WriteHeapProfile(f)
			time.Sleep(110 * time.Second)
		}
	}

}

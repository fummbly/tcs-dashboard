package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/sys"
)

func main() {

	for {
		getUsage()
		time.Sleep(4 * time.Second)
	}

}

func getUsage() {
	v, _ :=
		fmt.Printf("Total: %v, Free: %v,  UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
}

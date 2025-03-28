package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fummbly/tcs-dashboard/internal/sysmon"
)

func main() {

	info := make(chan sysmon.SysInfo)
	done := make(chan bool)
	ticker := time.NewTicker(30 * time.Second)

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go sysmon.Run(info, done, ticker)

	fmt.Println(<-info)

	sig := <-sigChan
	ticker.Stop()
	done <- true

	fmt.Println("Recieved Signal:", sig)

}

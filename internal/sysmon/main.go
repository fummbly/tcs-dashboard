package sysmon

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/shirou/gopsutil/v4/sensors"
)

type ProcessInfo struct {
	Pid        int32
	Name       string
	Status     string
	CPUPercent float64
	NumThreads int32
}

type TempInfo struct {
	Key  string
	Temp float64
}

type SysInfo struct {
	CPUPercent float64
	Host       string
	Processes  []ProcessInfo
	Temps      []TempInfo
}

func (p ProcessInfo) String() string {

	return fmt.Sprintf("%d: %s --- Status: %s, Number of Threads: %d", p.Pid, p.Name, p.Status, p.NumThreads)

}

func (t TempInfo) String() string {
	return fmt.Sprintf("%s: Tempurture: %f", t.Key, t.Temp)

}

func GetCPUPerc() float64 {

	percentage, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		log.Fatalf("Failed to get cpu usage percentage: %v", err)
		return 0.0
	}

	return percentage[0]

}

func GetHostInfo() string {

	host, err := host.Info()
	if err != nil {
		log.Fatalf("Failed to get host info: %v", err)
		return ""
	}

	return fmt.Sprint(host)

}

func GetProcesses() []ProcessInfo {

	processes, err := process.Processes()
	if err != nil {
		log.Fatalf("Failed to get processes: %v", err)
	}

	ps := []ProcessInfo{}

	for _, process := range processes {

		processName, _ := process.Name()
		processCPU, _ := process.CPUPercent()
		processStatus, _ := process.Status()
		processThreads, _ := process.NumThreads()

		p := ProcessInfo{
			Pid:        process.Pid,
			Name:       processName,
			CPUPercent: processCPU,
			Status:     processStatus[0],
			NumThreads: processThreads,
		}

		ps = append(ps, p)

	}

	return ps
}

func GetTemps() []TempInfo {
	temps, err := sensors.SensorsTemperatures()
	if err != nil {
		fmt.Printf("Failed to get temps: %v", err)
		return nil
	}

	ts := []TempInfo{}

	for _, temp := range temps {
		t := TempInfo{
			Key:  temp.SensorKey,
			Temp: temp.Temperature,
		}

		ts = append(ts, t)

	}

	return ts
}

func GetSysInfo() *SysInfo {
	info := &SysInfo{
		CPUPercent: GetCPUPerc(),
		Host:       GetHostInfo(),
		Processes:  GetProcesses(),
		Temps:      GetTemps(),
	}

	return info
}

func Run(info chan SysInfo, done chan bool, ticker *time.Ticker) {
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			info <- *GetSysInfo()

		}

	}

	/*
		ticker := time.NewTicker(30 * time.Second)

		for _ = range ticker.C {

			fmt.Println("Test")

			fmt.Println(getSysInfo())
		}
	*/

}

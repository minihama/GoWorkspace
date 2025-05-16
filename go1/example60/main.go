package main

import (
	"flag"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	envPtr := flag.String("env", "prod", "환경값(dev or prod)을 입력")
	cpuPtr := flag.Int("cpu", runtime.NumCPU(), "CPU수")
	flag.Parse() // 플래그 값을 읽어 들이기 위해 반드시 필요
	// 환경에 따라 다양한 처리를 하도록 함
	if *envPtr != "prod" {
		// DB 커넥션이나 기타 처리등을 함
	}
	// 입력한 CPU수가 사용 가능한 최대값 넘으면, 최대값으로 변경
	if *cpuPtr > runtime.NumCPU() {
		*cpuPtr = runtime.NumCPU()
	}
	// 0보다 크면, 다음 명령을 통해 입력한 CPU수만 사용하도록 함
	if *cpuPtr > 0 {
		runtime.GOMAXPROCS(*cpuPtr)
	}
	spew.Dump(envPtr)
	spew.Dump(cpuPtr)
}

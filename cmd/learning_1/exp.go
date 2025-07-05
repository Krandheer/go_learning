package cmd

import (
	"fmt"
	"runtime"
)

func BenchMarking(){
	procs := runtime.GOMAXPROCS(0)
	fmt.Println("go maxprocs:", procs)
}
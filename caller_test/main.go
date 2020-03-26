package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.caller() faild \n")
		return
	}

	name := runtime.FuncForPC(pc).Name()
	base := path.Base(file)
	fmt.Println(name)
	fmt.Println(base)
	fmt.Println(line)

}

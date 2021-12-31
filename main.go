package main

import (
	"flag"
	"fmt"
	"os2-2/fileWrite"
)

func sub1() {
	fmt.Println("sub1")
}

func main() {
	u := fileWrite.FileIO{}
	u = getOpts(u)

	u.FileWrite()
	sub1()
	fmt.Println("hello,世界")
}

func getOpts(u fileWrite.FileIO) fileWrite.FileIO {
	fmt.Println("Hello,getOpts")
	flag.BoolVar(&u.Buffering, "b", true, "buffering")
	flag.StringVar(&u.FileName, "filename", "testData.txt", "file name to write")
	flag.IntVar(&u.BufferSize, "buffersize", 32, "buffer size")
	flag.IntVar(&u.WriteSize, "writesize", 1024000, "write size")
	flag.Parse()
	return u
}

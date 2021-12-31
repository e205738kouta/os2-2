package fileWrite

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type FileIO struct {
	FileName   string
	Buffering  bool
	BufferSize int
	WriteSize  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (u *FileIO) FileWrite() {
	fmt.Println("Hello,FileWrite " + u.FileName)
	var f, _ = os.Create(u.FileName)
	// check(error)
	d1 := []byte("h")

	if !u.Buffering {
		u.BufferSize = 1
	}
	w := bufio.NewWriterSize(f, u.BufferSize)
	start := time.Now()
	for i := 0; i < u.WriteSize; i++ {
		w.Write(d1)
	}
	end := time.Now()
	diff := end.Sub(start)
	// result = diff.Milliseconds()
	fmt.Printf("result: %d milliseconds.\n", diff.Milliseconds())
	w.Flush()
	defer f.Close()
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

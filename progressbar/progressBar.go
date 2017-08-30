package main

import (
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

func main() {
	count := 1000
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
}

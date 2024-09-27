package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/mshafiee/mp4"
	"github.com/mshafiee/mp4/filter"
)

var pwd, _ = os.Getwd()

func main() {
	log.SetPrefix(fmt.Sprintf("Mp4 Clip (%d):", os.Getpid()))
	basepath := path.Join(pwd, "1.mp4")

	in, err := os.Open(basepath)
	if err != nil {
		fmt.Println(err)
	}
	defer in.Close()

	v, err := mp4.Decode(in)
	if err != nil {
		fmt.Println(err)
	}

	out, err := os.Create(path.Join(pwd, "clipped.mp4"))
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	if err := filter.EncodeFiltered(out, v, filter.Clip(time.Duration(0)*time.Second, time.Duration(5)*time.Second)); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

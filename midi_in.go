package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

/*
#cgo LDFLAGS: -lasound
#include <alsa/asoundlib.h>
#include <alsa/rawmidi.h>
*/
import "C"

func main() {
	fmt.Printf("hello, world\n")

	var handle_in *C.snd_rawmidi_t
	err := C.snd_rawmidi_open(&handle_in, nil, C.CString("hw:1,0,0"), 0)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "snd_rawmidi_open failed: %d\n", err)
	}

	if handle_in != nil {
		var ch byte
		for {
			C.snd_rawmidi_read(handle_in, unsafe.Pointer(&ch), 1)
			fmt.Printf("read %02x\n", ch)
		}

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			C.snd_rawmidi_drain(handle_in)
			C.snd_rawmidi_close(handle_in)
			os.Exit(0)
		}()
	}
}

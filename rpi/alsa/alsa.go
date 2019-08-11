package alsa

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

func ReadBytes(midiHandle string) <-chan byte {
	out := make(chan byte)
	go func() {
		var handle_in *C.snd_rawmidi_t
		err := C.snd_rawmidi_open(&handle_in, nil, C.CString(midiHandle), 0)
		if err != 0 {
			fmt.Fprintf(os.Stderr, "snd_rawmidi_open failed: %d\n", err)
		}

		if handle_in != nil {
			var ch byte
			for {
				C.snd_rawmidi_read(handle_in, unsafe.Pointer(&ch), 1)
				out <- ch
			}

			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c
				C.snd_rawmidi_drain(handle_in)
				C.snd_rawmidi_close(handle_in)
				close(out)
				os.Exit(0)
			}()
		}
	}()
	return out
}

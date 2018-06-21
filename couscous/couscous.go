package main

import (
	"fmt"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
	"github.com/robstein/couscous/midi"
	"os"
	"os/signal"
	"rand"
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
	// Set up the pipeline.
	c := inputStream("hw:1,0,0")
	midiMessages := midi.Parse(c)

	defer ws2811.Fini()
	err := ws2811.Init(18, 101, 255)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ws2811.Init failed: %d\n", err)
	} else {
		writeToLEDs(midiMessages)
	}
}

func inputStream(midiHandle string) <-chan byte {
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

const MidiKeyboardOffset = 21

func writeToLEDs(in <-chan Message) {
	go func() {
		for m := range in {
			fmt.Printf("Writing %s %d %d\n", message.Status, message.Data1, message.Data2)

			if message.Status == NoteOn {
				ws2811.SetLed(message.Data1-MidiKeyboardOffset, rand.Uint32())
			} else {
				ws2811.SetLed(message.Data1-MidiKeyboardOffset, 0)
			}

			err := ws2811.Render()
			if err != nil {
				ws2811.Clear()
				fmt.Fprintf(os.Stderr, "ws2811.Render failed: %d\n", err)
			}
		}
	}()
}

package main

import (
	"fmt"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
	"github.com/robstein/couscous/midi"
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
	// Set up the pipeline.
	c := inputStream("hw:1,0,0")
	midiMessages := midi.Parse(c)

	defer ws2811.Fini()
	err := ws2811.Init(18, KeyCount, 255)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ws2811.Init failed: %d\n", err)
	} else {
		writeToLEDs(midiMessages)
	}

	for {

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

const KeyCount = 88
const MidiKeyboardOffset = 28

func writeToLEDs(in <-chan midi.Message) {
	go func() {
		pedal_off := true
		for m := range in {
			fmt.Printf("%s %d %d\n", m.Status, m.Data1, m.Data2)

			if m.Status == midi.ControlChange && m.Data1 == 64 && m.Data2 == 127 {
				pedal_off = false
			}
			if m.Status == midi.ControlChange && m.Data1 == 64 && m.Data2 == 0 {
				pedal_off = true
			}

			led := m.Data1 - MidiKeyboardOffset
			if m.Status == midi.NoteOn {
				color := getColor(m.Data1, m.Data2)
				ws2811.SetLed(led, color)
			} else if m.Status == midi.NoteOff && pedal_off {
				ws2811.SetLed(led, 0)
			} else if pedal_off {
				for i := 0; i < KeyCount; i++ {
					ws2811.SetLed(i, 0)
				}
			}

			err := ws2811.Render()
			if err != nil {
				ws2811.Clear()
				fmt.Fprintf(os.Stderr, "ws2811.Render failed: %d\n", err)
			}
		}
	}()
}

func getColor(i int, v int) uint32 {
	out := 0xFFFFF
	switch i % 12 {
	case 0:
		out = 0x9400D3
	case 1:
		out = 0x4B0082
	case 2:
		out = 0x26007D
	case 3:
		out = 0x0000FF
	case 4:
		out = 0x008080
	case 5:
		out = 0x00FF00
	case 6:
		out = 0x80FF00
	case 7:
		out = 0xFFFF00
	case 8:
		out = 0xFFC000
	case 9:
		out = 0xFF7F00
	case 10:
		out = 0xFF4000
	case 11:
		out = 0xFF0000
	}
	return uint32(out)
}

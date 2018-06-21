package main

import (
	"fmt"
	"github.com/robstein/midi"
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
	messages := midi.Parse(c)

	for {
		message := <-messages
		fmt.Printf("%s %d %d\n", message.Status, message.Data1, message.Data2)
	}

	// c_messages := getMidiMessageStreamFromInputByteStream(c_bytes)
	// displayStringsFromMidiMessages(c_messages)

	// c_ledactions := getLEDActionsFromMidiMessages(c_messages)
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

type Message struct {
	Status   byte
	Note     byte
	Velocity byte
}

func getMidiMessageStreamFromInputByteStream(in <-chan byte) <-chan Message {
	out := make(chan Message)
	go func() {
		for ch := range in {
			if ch&0x80 == 0x80 {
				out <- Message{ch, <-in, <-in}
			}
		}
		close(out)
	}()
	return out
}

func displayStringsFromMidiMessages(in <-chan Message) {
	go func() {
		for m := range in {
			str := "?"
			if m.Status&0xF0 == 0x80 || m.Velocity == 0 {
				str = fmt.Sprintf("note off: %X", m.Note)
			} else if m.Status&0xF0 == 0x90 {
				str = fmt.Sprintf("note on: %X", m.Note)
			}
			fmt.Println(str)
		}
	}()
}

// type LEDAction struct {
// }
//
// func getLEDActionsFromMidiMessages(in <-chan Message) <-chan LEDAction {
// 	out := make(chan LEDAction)
// 	go func() {
// 		for m := range in {
// 		}
// 		close(out)
// 	}()
// 	return out
// }

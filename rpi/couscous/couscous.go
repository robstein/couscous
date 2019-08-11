package main

import (
	"fmt"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
	"github.com/robstein/couscous/alsa"
	"github.com/robstein/couscous/midi"
	"github.com/robstein/couscous/ws2812b"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

func main() {
	midiBytes := alsa.ReadBytes("hw:1,0,0")
	midiMessages := midi.Deserialize(midiBytes)
	effects := visualizer.GenerateEffects(midiMessages)

	defer ws2811.Fini()
	err := ws2811.Init(18, midi.KeyCount, 255)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ws2811.Init failed: %d\n", err)
	} else {
		ws2812b.WriteEffects(effects)
	}

	for {

	}
}

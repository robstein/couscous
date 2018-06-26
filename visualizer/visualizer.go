package visualizer

import "github.com/robstein/couscous/midi"

type Effect struct {
	Color uint32
	Led   int
}

const midiKeyboardOffset = 28

func GenerateEffects(in <-chan Message) <-chan Effect {
	out := make(chan Effect)
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

			var effect Effect
			led := m.Data1 - midiKeyboardOffset
			if m.Status == midi.NoteOn {
				color := getColor(m.Data1, m.Data2)
				effect = Effect{led, color}
			} else if m.Status == midi.NoteOff && pedal_off {
				effect = Effect{led, 0}
			} else if pedal_off {
				// for i := 0; i < midi.KeyCount; i++ {
				// 	ws2811.SetLed(i, 0)
				// }
			}

			out <- effect
		}
		close(out)
	}()
	return out
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

// Really we want to go from `midi(t)` to `led(t)`
//  so the input will be a stream: map<timestamp, pair<note, velocity, pedal?>> // re think how the pedal works, and how on and off works.
// the output will be a stream: map<timestamp, pair<triplet<x,y,z>, rgb>

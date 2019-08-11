package ws2812b

import "github.com/jgarff/rpi_ws281x/golang/ws2811"

// func Serialize(in <-chan Effect) <-chan byte {
// }
//
// func WriteBytes(in <-chan byte) {
// }

func WriteEffects(in <-chan Effect) {
	go func() {
		for e := range in {
			ws2811.SetLed(e.Led, e.Color)
			err := ws2811.Render()
			if err != nil {
				ws2811.Clear()
				fmt.Fprintf(os.Stderr, "ws2811.Render failed: %d\n", err)
			}
		}
	}()
}

package midi

type Message struct {
	Status MessageStatus
	Data1  int // note, controller, program
	Data2  int // velocity, pressure, value
}

type MessageStatus string

type messageCategory string

const (
	NoteOff             MessageStatus = "Note Off"
	NoteOn              MessageStatus = "Note On"
	AfterTouch          MessageStatus = "AfterTouch"
	ControlChange       MessageStatus = "Control Change"
	ProgramChange       MessageStatus = "Program (patch) change"
	ChannelPressure     MessageStatus = "Channel Pressure"
	PitchWheel          MessageStatus = "Pitch Wheel"
	SystemExclusive     MessageStatus = "System Exclusive"
	GMSystem            MessageStatus = "GM System Enable/Disable" // Not implemented
	MasterVolume        MessageStatus = "Master Volume"            // Not implemented
	MTCQuarterFrame     MessageStatus = "MTC Quarter Frame Message"
	SongPositionPointer MessageStatus = "Song Position Pointer"
	SongSelect          MessageStatus = "Song Select"
	TuneRequest         MessageStatus = "Tune Request"
	MIDIClock           MessageStatus = "MIDI Clock"
	MIDIStart           MessageStatus = "MIDI Start"
	MIDIContinue        MessageStatus = "MIDI Continue"
	MIDIStop            MessageStatus = "MIDI Stop"
	ActiveSense         MessageStatus = "Active Sense"
	Reset               MessageStatus = "Reset"
	UndefinedStatus     MessageStatus = "Undefined"

	Voice           messageCategory = "Voice"
	System_Common   messageCategory = "System Common"
	System_Realtime messageCategory = "System Realtime"
)

const KeyCount = 88

// This Midi receiver changes NoteOn messages with velocity '0' to NoteOff messages.
func Deserialize(in <-chan byte) <-chan Message {
	out := make(chan Message)
	go func() {
		// running status buffer is cleared at power up.
		runningStatus := UndefinedStatus
		for b := range in {
			status := getStatus(b)
			category := getCategory(status)

			if category == Voice {
				// running status buffer stores the status when a Voice Category Status (ie, 0x80 to 0xEF) is received.
				runningStatus = status
			} else if category == System_Common {
				// running status buffer is cleared when a System Common Category Status (ie, 0xF0 to 0xF7) is received.
				runningStatus = UndefinedStatus
			}
			// Any data bytes are ignored when the running status buffer is cleared.
			if runningStatus == UndefinedStatus && isDataByte(b) {
				continue
			}

			var message Message
			if isDataByte(b) {
				switch runningStatus {
				case NoteOff, NoteOn, AfterTouch, ControlChange, PitchWheel:
					message = Message{runningStatus, int(b), int(<-in)}
				case ProgramChange, ChannelPressure:
					message = Message{runningStatus, int(b), 0}
				}
			} else if isStatusByte(b) {
				switch status {
				case NoteOff, NoteOn, AfterTouch, ControlChange, PitchWheel, SongPositionPointer:
					message = Message{status, int(<-in), int(<-in)}
				case ProgramChange, ChannelPressure, MTCQuarterFrame, SongSelect:
					message = Message{status, int(<-in), 0}
				case TuneRequest, MIDIClock, MIDIStart, MIDIContinue, MIDIStop, ActiveSense, Reset:
					message = Message{status, 0, 0}
				case SystemExclusive, GMSystem, MasterVolume: // Not implemented
				}
			}

			if message.Status == NoteOn && message.Data2 == 0 {
				message.Status = NoteOff
			}

			if getCategory(message.Status) == Voice {
				out <- message
			}
		}
		close(out)
	}()
	return out
}

func isDataByte(b byte) bool {
	return b < 0x80
}

func isStatusByte(b byte) bool {
	return b >= 0x80
}

func getCategory(s MessageStatus) messageCategory {
	var c messageCategory
	switch s {
	case NoteOff, NoteOn, AfterTouch, ControlChange, ProgramChange, ChannelPressure, PitchWheel:
		c = Voice
	case SystemExclusive, GMSystem, MasterVolume, MTCQuarterFrame, SongPositionPointer, SongSelect, TuneRequest:
		c = System_Common
	case MIDIClock, MIDIStart, MIDIContinue, MIDIStop, ActiveSense, Reset:
		c = System_Realtime
	}
	return c
}

func getStatus(b byte) MessageStatus {
	out := UndefinedStatus
	if b >= 0x80 && b <= 0x8F {
		out = NoteOff
	} else if b >= 0x90 && b <= 0x9F {
		out = NoteOn
	} else if b >= 0xA0 && b <= 0xAF {
		out = AfterTouch
	} else if b >= 0xB0 && b <= 0xBF {
		out = ControlChange
	} else if b >= 0xC0 && b <= 0xCF {
		out = ProgramChange
	} else if b >= 0xD0 && b <= 0xDF {
		out = ChannelPressure
	} else if b >= 0xE0 && b <= 0xEF {
		out = PitchWheel
	} else if b == 0xF0 || b == 0xF7 {
		out = SystemExclusive
	} else if b == 0xF1 {
		out = MTCQuarterFrame
	} else if b == 0xF2 {
		out = SongPositionPointer
	} else if b == 0xF3 {
		out = SongSelect
	} else if b == 0xF6 {
		out = TuneRequest
	} else if b == 0xF8 {
		out = MIDIClock
	} else if b == 0xFA {
		out = MIDIStart
	} else if b == 0xFB {
		out = MIDIContinue
	} else if b == 0xFC {
		out = MIDIStop
	} else if b == 0xFE {
		out = ActiveSense
	} else if b == 0xFF {
		out = Reset
	}
	return out
}

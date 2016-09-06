package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string)  {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("这个类型我不会---",mtype)
		return
	}

	p.Play(source)
}

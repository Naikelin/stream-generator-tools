package stream

import (
	"github.com/tjarratt/babble"
)

func GenerateSentence() string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	return babbler.Babble()
}

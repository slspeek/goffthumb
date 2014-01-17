package ffmpeg

import (
	"testing"
)

const (
	input  = "test-data/big_buck_bunny.ogv"
	output = "popeye.jpg"
	size   = 0
	time   = "50%"
)

func TestCreateThumb(t *testing.T) {
	err := CreateThumb(input, output, time, size)
	if err != nil {
		t.Fatal("Create thumb returned: ", err)
	}

}

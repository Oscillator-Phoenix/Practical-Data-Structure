package bitmap

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var x byte = 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%0b\n", ^x)
}

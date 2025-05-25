package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Adder(10, 10)
	want := 20
	if got != want {
		t.Errorf("Want %d , But got %d", want, got)
	}
}

func ExampleAdder() {
	sum := Adder(4, 5)
	fmt.Println(sum)
	// Output: 9
}

package nlp_test

import (
	"fmt"

	nlp "github.com/parikxxit/ardanlab-nlp"
)

func ExampleTokenize() {
	text := "this is input"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)
	// Output:
	// [this is input]
}

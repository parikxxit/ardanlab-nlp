package nlp

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "some new input"
	gotTokens := Tokenize(text)
	expectedTokens := []string{"some", "new", "input"}
	if !reflect.DeepEqual(expectedTokens, gotTokens) {
		t.Fatalf("Expected %#v got %#v", expectedTokens, gotTokens)
	}
}

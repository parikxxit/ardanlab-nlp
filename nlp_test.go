package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	text := "some new input"
	gotTokens := Tokenize(text)
	expectedTokens := []string{"some", "new", "input"}
	require.Equal(t, expectedTokens, gotTokens)
	/* before testify
	if !reflect.DeepEqual(expectedTokens, gotTokens) {
		t.Fatalf("Expected %#v got %#v", expectedTokens, gotTokens)
	}
	*/
}

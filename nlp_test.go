package nlp

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

type tokenizeTestCase struct {
	text   string   `toml:"text"`
	tokens []string `toml:"tokens"`
}
type tokenizeTestCases struct {
	cases []tokenizeTestCase `toml:"cases"`
}

func loadTokenizeTestCase(t *testing.T) []tokenizeTestCase {
	data, err := ioutil.ReadFile("tokenize_cases.toml")
	require.NoError(t, err, "Read File")
	var tc tokenizeTestCases
	err = toml.Unmarshal(data, &tc)
	require.NoError(t, err, "Unmarhaling")
	fmt.Printf("%#v", tc)
	return tc.cases
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range loadTokenizeTestCase(t) {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tokens, tc.tokens)
		})
	}
}

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

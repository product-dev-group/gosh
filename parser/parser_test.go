package parser

import (
	"gosh/lexer"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TokenTypeEof            = iota // ファイルの終端
	TokenTypeEol                   // 改行
	TokenTypePipe                  // |
	TokenTypeRedirect              // >
	TokenTypeRedirectDouble        // >>
	TokenTypeCommand
	TokenTypeOption
	TokenTypeWord
)

func TestTparse(t *testing.T) {
	// Example echo "hello"
	var testLexer = []lexer.Token{
		{TokenType: TokenTypeCommand, Text: "echo"},
		{TokenType: TokenTypeWord, Text: "hello"},
	}

	// var testLexer2 = []lexer.Token{
	// 	{TokenType: TokenTypeCommand, Text: "echo"},
	// 	{TokenType: TokenTypeWord, Text: "world"},
	// }

	var testParser = ParserStruct{
		Command: "echo",
		Arg:     []string{"hello"},
	}

	returnToken := Parse(testLexer)

	require.Equal(t, testParser, returnToken)
}

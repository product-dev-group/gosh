package parser

// "gosh/shellinterface"

import (
	"fmt"
	"gosh/lexer"
	"log"
)

type ParserStruct struct {
	Operand    string
	Command    string
	Option     []string
	Arg        []string
	NextStruct []ParserStruct
}

// 抽象構文木にする
func Execute() {

}

func Parse(lexerTokens []lexer.Token) ParserStruct {
	// Lexer側でエラーハンドリングしてたらいらない
	if len(lexerTokens) == 0 {
		log.Fatalf("No element from Lexer")
	}

	var parserStruct ParserStruct

	// ここの処理が再帰になるからメソッドに切り分けてその中で再帰させる
	for _, token := range lexerTokens {
		switch token.TokenType {
		case 0:
			fmt.Println("EOF")
		case 1:
			fmt.Println("EOL")
		case 2:
			fmt.Println("PIPE")
		case 3:
			fmt.Println("REDIRECT")
		case 4:
			fmt.Println("REDIRECTDOUBLE")
		case 5:
			parserStruct.Command = token.Text
		case 6:
			fmt.Println("OPTION")
		case 7:
			parserStruct.Arg = append(parserStruct.Arg, token.Text)
		}
	}

	// どこかでExecuteを実行
	return parserStruct
}

func TParse(lexerToken []lexer.Token) []lexer.Token {
	// Lexer側でエラーハンドリングしてたらいらない
	if len(lexerToken) == 0 {
		log.Fatalf("No element from Lexer")
	}

	// どこかでExecuteを実行
	fmt.Println(lexerToken)

	return lexerToken
}

// 入力
// type Token struct {
// 	TokenType int,
// 	Text string,
// }

// 例）echo "hello"

// giftForMyo := [Token1, Token2]
/*
Token1 {
	TokenType: TokenType.command,
	  Text "echo",
}
Token2 {
	TokenType: TokenType.word,
	  Text "hello",
}
*/

// 出力
// var giftForTetsu ParserStruct
// 	   giftForTetsu = &model.ParserStruct{
//   		Operand: "",// リダイレクト、パイプ等々はここにまとめる
//      	Command: "echo",
//    	Option: [],
//       	Arg: ["arg1"],
//       	NextStruct nil,
//    }

// 再帰処理

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "github.com/product-dev-group/gosh/blob/main/interface/shellInterface" // できないの要調査
)

func main() {
	const prompt string = "gosh % "

	for {
		// 1-1. プロンプトを表示する
		fmt.Printf("%s", prompt)

		// 1-2. 標準入力から1行読み込む
		// var input1, input2 string
		// fmt.Scan(&input1, &input2)
		// log.Println("input: ", input1)
		scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
		scanner.Scan()                        // １行分の入力を取得する
		// fmt.Println("output: ", scanner.Text())
		// fmt.Printf("type: %s %T\n", scanner.Text(), scanner.Text())

		// 2. 入力を構文解析する
		args := parse(scanner.Text())

		if len(args) > 0 {
			execute(args)
		}
	}
}

// コマンド文字列を受け取り、スペースで分割する
func parse(line string) []string {
	args := strings.Split(line, " ")
	return args
}

// コマンドを実行する
func execute(args []string) {
	fmt.Printf("test ")
	log.Printf("execute: %#v\n", args)
}

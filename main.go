package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/trypophob1a/learn_golang/pkg/top_ten_words_list"
)

func main() {
	args := os.Args
	if len(args) > 1 && args[1] == "-f" {
		top_ten_words_list.PrintTop(top_ten_words_list.GetTextInFile(args[2]))
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		top_ten_words_list.PrintTop(text)
		fmt.Print("please press Enter for exit")
		_, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}
	}
}

package boluobaoLib

import (
	"bufio"
	"fmt"
	"github.com/AlexiaVeronica/boluobaoLib/boluobaomodel"
	"os"
	"strconv"
	"strings"
)

const (
	username      = `androiduser`
	password      = "1a#$51-yt69;*Acv@qxq"
	timeout       = 30
	threadNum     = 32
	baseURL       = "https://api.sfacg.com"
	androidApiKey = "FN_Q29XHVmfV3mYX"
	userAgent     = "boluobao/5.0.62(android;25)/OPPO/%s/OPPO"
)

type continueFunction func(chapter boluobaomodel.ChapterList) bool
type contentFunction func(chapter *boluobaomodel.ContentData)

type bookInfoFunction func(index int, bookInfo boluobaomodel.BookInfo)

func GetUserInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		return number
	}
}

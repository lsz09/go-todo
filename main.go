package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("사용법:")
		fmt.Println("  go run . add \"할 일 내용\"")
		fmt.Println("  go run . list")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("추가할 할 일을 입력하세요.")
			return
		}

		task := os.Args[2]
		fmt.Println("할 일 추가:", task)

	case "list":
		fmt.Println("할 일 목록을 표시합니다.")

	default:
		fmt.Println("알 수 없는 명령어:", command)
	}
}

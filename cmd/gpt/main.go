package main

import (
	"bufio"
	"github.com/daijun4you/gpt/configs"
	"github.com/daijun4you/gpt/internal"
	"os"
)

func main() {
	role, err := configs.Instance.Get("lose-weight", "roles.ini")
	if err != nil {
		println(err.Error())
		return
	}

	gpt := new(internal.GPT)
	gpt.Init(role)

	println("请您提问：\n")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		gpt.Talk(s.Text())
	}
}

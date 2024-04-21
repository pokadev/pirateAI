package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/thread"
	"os"
)

func main() {
	fmt.Println("Welcome to Pirate Bot")
	fmt.Println("---------------------")
	err := getOrSetKey()
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	myThread := thread.New().AddMessages(
		thread.NewSystemMessage().AddContent(
			thread.NewTextContent("All replies must be given in a pirate style of speech"),
		),
		thread.NewUserMessage().AddContent(
			thread.NewTextContent(text),
		),
	)

	err = openai.New().Generate(context.Background(), myThread)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error: Something went wrong. Please check your API KEY & account")
		os.Exit(1)
	}

	fmt.Println("Pirate :" + myThread.LastMessage().Contents[0].AsString())
	os.Exit(0)
}

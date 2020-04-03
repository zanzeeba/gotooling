package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zanzeeba/gotooling/message"
	"os"
	"strings"
)

func main() {
	var name string
	var greeting string
	var prompt bool
	var preview bool

	// Define flags
	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "phrase to use within the greeting")
	flag.BoolVar(&prompt, "prompt", false, "use prompt to input name and message")
	flag.BoolVar(&preview, "preview", false, "use preview to output message without writing to /etc/motd")

	// Parse flags
	flag.Parse()

	// Optionally print flags and exit if DEBUG is set
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name:", name)
		fmt.Println("Greeting:", greeting)
		fmt.Println("Prompt:", prompt)
		fmt.Println("Preview:", preview)

		os.Exit(0)
	}

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Your Greeting: ")
	//phrase, _ := reader.ReadString('\n')
	//phrase = strings.TrimSpace(phrase)
	//
	//fmt.Print("Your Name: ")
	//name, _ = reader.ReadString('\n')
	//name = strings.TrimSpace(name)

	// If no arguments passed, show usage
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	// Conditionally read from stdin
	if prompt {
		name, greeting = renderPrompt()
	}

	// Generate message
	m := message.Greeting(name, greeting)

	f, err := os.OpenFile("motd.txt", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Print("Error: unable to open file 60")
		os.Exit(1)
	}

	defer f.Close()

	m = message.Greeting(name, greeting)
	// Either preview message or write to file
	if preview {
		fmt.Println(m)
	} else {
		// Write content
		f, err := os.OpenFile("motd.txt", os.O_RDWR|os.O_CREATE, 0755)

		if err != nil {
			fmt.Println("Error: Unable to open to motd.txt 85")
			os.Exit(1)
		}

		defer f.Close()

		_, err = f.Write([]byte(m))

		if err != nil {
			fmt.Println("Error: Failed to write to motd 94")
			os.Exit(1)
		}
	}
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/zanzeeba/message"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := message.Greeting(name, phrase)
	_ = ioutil.WriteFile("/etc/motd", []byte(m), 0644)
}

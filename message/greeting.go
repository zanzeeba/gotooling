package message

import "fmt"

func greeting(name, message string) string {
	return fmt.Sprintf("%s, %s", message, name)
}

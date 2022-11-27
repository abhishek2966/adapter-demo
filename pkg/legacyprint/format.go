package legacyprint

import "fmt"

type Object struct{}

func (l *Object) Format(s string) (formattedMsg string) {
	formattedMsg = fmt.Sprintf("Legacy Format: %s\n", s)
	fmt.Print(formattedMsg)
	return
}

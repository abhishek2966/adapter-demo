package newprint

import (
	"fmt"

	"github.com/abhishek2966/adapter-demo/pkg/legacyprint"
)

type PrinterAdapter struct {
	OldPrinterObject *legacyprint.Object
	Msg              string
}

func (p *PrinterAdapter) Format() (formattedMsg string) {
	if p.OldPrinterObject != nil {
		formattedMsg = p.OldPrinterObject.Format(p.Msg)
		return
	}

	formattedMsg = fmt.Sprintf("New Format: %s\n", p.Msg)
	fmt.Print(formattedMsg)
	return
}

// type NewObject struct {
// 	Msg string
// }

// func (l *NewObject) Format() (formattedMsg string) {
// 	formattedMsg = fmt.Sprintf("New Format: %s\n", l.Msg)
// 	fmt.Print(formattedMsg)
// 	return
// }

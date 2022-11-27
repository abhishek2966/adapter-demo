package main

import (
	"github.com/abhishek2966/adapter-demo/pkg/legacyprint"
	"github.com/abhishek2966/adapter-demo/pkg/newprint"
)

func main() {
	f := legacyprint.Object{}
	//f.Format("Hello World") // Legacy Format: Hello World

	// g := newprint.NewObject{Msg: "Hello World!"}
	// g.Format() // New Format: Hello World!

	h := newprint.PrinterAdapter{OldPrinterObject: &f, Msg: "Hello World!"}
	h.Format() // Legacy Format: Hello World!

	h = newprint.PrinterAdapter{Msg: "Hello World!"}
	h.Format() // New Format: Hello World!
}

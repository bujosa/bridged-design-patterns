package structural

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterAPI1 struct{}

func (d *PrinterAPI1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterAPI2 struct {
	Writer io.Writer
}

// PrintMessage prints the given message using the PrinterAPI2.
// It requires an io.Writer to be passed to the PrinterAPI2 instance.
// Returns an error if the Writer is not set.
func (d *PrinterAPI2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterAPI2")
	}

	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}

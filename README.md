# Bridge Pattern in Go
The Bridge pattern is a structural design pattern that decouples an abstraction from its implementation so that the two can vary independently. This pattern involves an interface acting as a bridge which makes the functionality of concrete classes independent from interface implementer classes. Both types of classes can be altered structurally without affecting each other.

In this Go code, the Bridge pattern is implemented using `PrinterAPI` as the bridge interface, `PrinterAPI1` and `PrinterAPI2` as the concrete implementer classes, and `NormalPrinter` and `PacktPrinter` as the abstraction classes.

## Code Explanation

### PrinterAPI interface
This is the bridge implementer interface. It has a method `PrintMessage(string) error`.

### PrinterAPI1 and PrinterAPI2
These are the concrete implementers of the `PrinterAPI` interface. `PrinterAPI1` uses the standard output to print the message while `PrinterAPI2` uses an `io.Writer` to print the message. If the Writer in `PrinterAPI2` is not set, it returns an error.

### PrinterAbstraction Interface
This is the abstraction interface. It has a method `Print() error`.

### NormalPrinter and PacktPrinter
These are the concrete implementers of the `PrinterAbstraction` interface. They use a `PrinterAPI` to print a message. `NormalPrinter` prints the message as is, while `PacktPrinter` adds a prefix to the message.

## Usage

You can use the `NormalPrinter` or `PacktPrinter` to print a message. You need to provide a `PrinterAPI` (either `PrinterAPI1` or `PrinterAPI2`) to these printers. If you use PrinterAPI2, you also need to provide an `io.Writer`.

```go
printer := &NormalPrinter{
    Msg:     "Hello, World!",
    Printer: &PrinterAPI1{},
}
err := printer.Print()
if err != nil {
    log.Fatal(err)
}
```

This pattern allows you to change the way the message is printed (by switching between `PrinterAPI1` and `PrinterAPI2`) without changing the printer classes (`NormalPrinter` and `PacktPrinter`). Similarly, you can change the printer classes without affecting the `PrinterAPI` classes.
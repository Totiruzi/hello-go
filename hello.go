package main
import (
    "fmt"
    "log"
    "greetings/greetings"
)

// import "rsc.io/quote"

func main() {
//     fmt.Println(quote.Go())

    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names
    names := []string {"Chris", "Colorful", "Gabriel", "Daniil"}

//     format := greetings.randomFormat()
    messages, err := greetings.Hellos(names)

    if err != nil {
    log.Fatal(err)
    }
    fmt.Println(messages)
}

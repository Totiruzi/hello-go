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

//     format := greetings.randomFormat()
    message, err := greetings.Hello("Gabriel")

    if err != nil {
    log.Fatal(err)
    }
    fmt.Println(message)
}

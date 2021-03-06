package greetings

import (
    "fmt"
    "math/rand"
    "time"
    "errors"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
//     message := fmt.Sprintf("Welcome %v! to Go World", name)
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    // Loop through the received sliced of name, calling the Hello function to get a message for each name
    for _, name := range names {
        message, err := Hello(name)
        if err !=nil {
        return nil, err
        }
        // In the map, associate the retrieved message with the name
        messages[name] = message
    }
    return messages, nil
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    formats := []string {
        "Hi %v!, nice to have you here",
        "%v!, Welcome to Go land ",
        "It's so beautiful to have you here %v",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

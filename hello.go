package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Create a slice to hold names
    var names []string

    // Create a scanner to read input from the terminal
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Enter names (type 'done' to finish):")

    for {
        fmt.Print("Enter name: ")
        // Scan the input from the terminal
        scanner.Scan()
        input := strings.TrimSpace(scanner.Text())

        // Break the loop if the user types "done"
        if input == "done" {
            break
        }

        // Append the input to the slice
        names = append(names, input)
    }

    // Print the names in the format [ "Alice", "Bob", "Charlie" ]
    fmt.Println("\nList of names entered:")
    fmt.Printf("[ ")
    for i, name := range names {
        if i > 0 {
            fmt.Print(",")
        }
        fmt.Printf("\"%s\"", name)
    }
    fmt.Println(" ]")
}

package main
// Import required packages
import (
    "fmt"
    "os"
)

// Main function - entry point
func main() {

    /*
     * Get list of all environmental variables. Go through each
     * using a for loop and print it.
     */
    for _, e := range os.Environ() {
        fmt.Printf("%s\n",e)
    }
}

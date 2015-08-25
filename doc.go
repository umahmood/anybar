/*
Package anybar provides the ability to send commands to an Anybar.app instance.

The below example shows how to change the color of the dot to blue:

    package main

    import (
        "github.com/umahmood/anybar"
    )

    func main() {
        port := 1738

        err := anybar.Command(port, anybar.Blue)

        if err != nil {
            // handle error
        }
    }
*/
package anybar

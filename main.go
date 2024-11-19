package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stderr, "go-oci-adb")
	fmt.Fprint(os.Stderr, "2024 github.com/icecreammachine/go-oci-adb\n\n")

	OpenDBConnection()

	defer func() {
		CloseDBConnection()
		if err := recover(); err != nil { //catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()
}

package main

import (
	"ddbtest/internal/ddb"
	"fmt"
)

func main() {
	ddb.DB.Connect()
	fmt.Printf("Hello, world.\n")
}

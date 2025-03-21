package main

import (
	"resedist/cmd"
)

// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func main() {
	cmd.Execute()
}

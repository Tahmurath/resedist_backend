package main

import (
	//"github.com/outrigdev/outrig"

	"resedist/cmd"
)

// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func main() {
	//outrig.Init(nil)
	//defer outrig.AppDone()
	cmd.Execute()
}

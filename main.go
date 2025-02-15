package main

import (
	"s3-go-file-handling/cmd"
	"s3-go-file-handling/config"
	"s3-go-file-handling/helpers"
)

func main() {
	// setup logger
	helpers.SetupLogger()

	// setup config
	config.SetupConfig()

	// setup http server
	cmd.SetupHTTP()
}

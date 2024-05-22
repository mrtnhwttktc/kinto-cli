/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mrtnhwttktc/kinto-cli/cmd"
	_ "github.com/mrtnhwttktc/kinto-cli/internal/config"
	_ "github.com/mrtnhwttktc/kinto-cli/internal/log"
)

func main() {
	cmd.Execute()
}

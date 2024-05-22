package main

import (
	"github.com/mrtnhwttktc/kinto-cli/cmd"
	_ "github.com/mrtnhwttktc/kinto-cli/internal/config"
	_ "github.com/mrtnhwttktc/kinto-cli/internal/logger"
)

func main() {
	cmd.Execute()
}

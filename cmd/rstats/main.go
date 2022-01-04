package main

import (
	"dev.maizy.ru/rstats/cmd/rstats/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.MousetrapHelpText = ""
	cmd.Execute("web-ui")
}

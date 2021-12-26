package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dev.maizy.ru/rstats/rstats_app"
	"dev.maizy.ru/rstats/rstats_app/db"
)

var rootCmd = &cobra.Command{
	Use: "rstats",
}

func Execute(defaultCommand string) {
	rootCmd.SetVersionTemplate(`{{printf "%s version: %s\n" .Name .Version}}`)
	rootCmd.Version = rstats_app.GetVersion()

	var commandFound bool
	commands := rootCmd.Commands()
	for _, command := range commands {
		for _, arg := range os.Args[1:] {
			if command.Name() == arg {
				commandFound = true
				break
			}
		}
	}
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "--version" ||
		os.Args[1] == "-v") {
		commandFound = true
	}
	if commandFound == false {
		args := append([]string{defaultCommand}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printErrF(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", a...)
}

func connectToDbOrExit() *db.Connection {
	times, err := db.CheckAndOpenReadonly("dirtrally-laptimes.db", "DB_LAPTIMES")
	if err != nil {
		printErrF("unable to open laptimes DB: %s", err)
		os.Exit(2)
	}

	data, err := db.CheckAndOpenReadonly("dirtrally-lb.db", "DB_DATA")
	if err != nil {
		printErrF("unable to open data DB: %s", err)
		os.Exit(2)
	}
	return &db.Connection{times, data}
}

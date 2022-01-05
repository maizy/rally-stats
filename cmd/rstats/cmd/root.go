package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dev.maizy.ru/rstats/rstats_app"
	"dev.maizy.ru/rstats/rstats_app/db"
	"dev.maizy.ru/rstats/rstats_app/dicts"
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

func anyKeyToExit(exitCode int) {
	println("\nPress return to exit ...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	os.Exit(exitCode)
}

func connectToDbOrExit() *db.DBContext {
	times, err := db.CheckAndOpenReadonly("dirtrally-laptimes.db", "DB_STAGETIMES")
	if err != nil {
		printErrF("unable to open stage times DB: %s", err)
		anyKeyToExit(2)
	}
	return &db.DBContext{times, dicts.LoadDicts()}
}

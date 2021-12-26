package cmd

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	// "dev.maizy.ru/rstats/rstats/db"
	"dev.maizy.ru/rstats/rstats_app/web"
)

// flags
var (
	bindPort  int
	bindHost  string
	debugMode bool
	devMode   bool
)

func init() {
	webUiCmd.PersistentFlags().IntVar(&bindPort, "port", 58367, "bind port")
	webUiCmd.PersistentFlags().StringVar(&bindHost, "host", "127.0.0.1", "bind host")
	webUiCmd.PersistentFlags().BoolVar(&debugMode, "debug", false, "enable debug mode")
	webUiCmd.PersistentFlags().BoolVar(&devMode, "dev-mode", false, "enable dev mode (templates hot reload)")

	rootCmd.AddCommand(webUiCmd)
}

var webUiCmd = &cobra.Command{
	Use:   "web-ui",
	Short: "Launch Web UI (default)",
	Run: func(cmd *cobra.Command, args []string) {
		//conn := connectToDbOrExit()
		//defer db.CloseConnection(conn)

		if !debugMode {
			gin.SetMode(gin.ReleaseMode)
		}

		engine := gin.Default()

		if err := engine.SetTrustedProxies(nil); err != nil {
			panic(err)
		}
		web.SetupMiddlewares(engine)
		web.SetupTemplates(engine, devMode)
		web.AppendRouters(engine, devMode)

		addr := fmt.Sprintf("%s:%d", bindHost, bindPort)
		fmt.Printf("launching rstats at: http://%s\n", addr)
		err := engine.Run(addr)
		if err != nil {
			printErrF("Unable to start rstats: %s", err)
			os.Exit(1)
		}
	},
}

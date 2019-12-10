package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	urlOpenDelayMillis  = 1500
	sentryPresDirectory = "./sentry"
	sentryAssertsPath   = "/assets"
	trayIconName        = "kelp-icon@1-8x.png"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the Sentry GUI",
}

type serveInputs struct {
	port              *uint16
	dev               *bool
	devAPIPort        *uint16
	horizonTestNetURI *string
	horizonPubNetURI  *string
	noHeaders         *bool
}

// 服务初始化
func init() {
	options := serveInputs{}
	options.port = serveCmd.Flags().Uint16P("port", "p", 8000, "port on which to serve")
	options.dev = serveCmd.Flags().Bool("dev", false, "run in dev mode for hot-reloading os JS code")
	options.devAPIPort = serveCmd.Flags().Uint16("dev-api-port", 8001, "port on which to run API server when in dev mode")
	options.horizonTestNetURI = serveCmd.Flags().String("horizon-testNet-uri", "https://horizon-testnet.stellar.org", "URI to use for the horizon instance connected to the Stellar Test Network (must contain the word 'test')")
	options.horizonPubNetURI = serveCmd.Flags().String("horizon-pubNet-uri", "https://horizon.stellar.org", "URI to use for the horizon instance connected to the Stellar Public Network (must not contain the word 'test')")
	options.noHeaders = serveCmd.Flags().Bool("no-headers", false, "do not set X-App-Name and X-App-Version headers on requests to horizon")

	serveCmd.Run = func(cmd *cobra.Command, args []string) {
		log.Printf("启动哨兵GUI服务: %s [%s]\n", version, gitBranch)
		checkInitRootFlags()
		if !strings.Contains(*options.horizonTestNetURI, "test") {
			panic("测试网络请求地址参数必须包含[test]字符")
		}

		if strings.Contains(*options.horizonPubNetURI, "test") {
			panic("生产网络请求地址不能包含[test]字符")
		}

	}
}

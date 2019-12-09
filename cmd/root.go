package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"sentry-bot/support/networking"
	"strings"
)

// 构建参数
var (
	version   string
	gitBranch string
	gitHash   string
	buildDate string
	env       string
)

const (
	envRelease = "release"
	envDev     = "dev"
	rootShort  = "Sentry 是一款针对[xx]市场的通用交易机器人"
	rootLong   = `Sentry 是一款针对[xx]市场的通用交易机器人 (https://github.com)
更多文档请查看: https://github.com
`
	sentryExamples = ""
)

var RootCmd = &cobra.Command{
	Use:     "sentry",
	Short:   rootShort,
	Long:    rootLong,
	Example: sentryExamples,
	Run: func(cmd *cobra.Command, args []string) {
		intro := `
______  __    ___________           ________           _____                
___  / / /_______  /__  /_____      __  ___/_____________  /____________  __
__  /_/ /_  _ \_  /__  /_  __ \     _____ \_  _ \_  __ \  __/_  ___/_  / / /
_  __  / /  __/  / _  / / /_/ /     ____/ //  __/  / / / /_ _  /   _  /_/ / 
/_/ /_/  \___//_/  /_/  \____/      /____/ \___//_/ /_/\__/ /_/    _\__, /  
                                                                   /____/
` + version + `
`
		fmt.Println(intro)

		//TODO server.run()
	},
}

var rootCtxRequestURL *string

// 初始化
func init() {
	validateBuild()
	rootCtxRequestURL = RootCmd.PersistentFlags().String("ctx-rest-url", "", "URL to use for the CCXT-rest API. Takes precendence over the CCXT_REST_URL param set in the botConfg file for the trade command and passed as a parameter into the Kelp subprocesses started by the GUI (default URL is https://localhost:3000)")
	RootCmd.AddCommand(versionCmd)
}

// 检查初始化flag参数
func checkInitRootFlags() {
	if *rootCtxRequestURL != "" {
		*rootCtxRequestURL = strings.TrimSuffix(*rootCtxRequestURL, "/")
		//校验 url是否符合规范
		if !strings.HasPrefix(*rootCtxRequestURL, "http://") && !strings.HasPrefix(*rootCtxRequestURL, "https://") {
			panic("'ctx-request-url' argument must start with either `http://` or `https://`")
		}

		e := testCtxURL(*rootCtxRequestURL)
		if e != nil {
			panic(e)
		}

		//TODO
	}
}

// 验证构建是否正确
func validateBuild() {
	if version == "" || buildDate == "" || gitBranch == "" || gitHash == "" {
		fmt.Println("version information not include,please build using the build script(scripts/build.sh)")
		os.Exit(1)
	}
}

// 测试url请求是否可用
func testCtxURL(ctxUrl string) error {
	e := networking.JSONRequest(http.DefaultClient, "GET", ctxUrl, "", map[string]string{}, nil, "")
	if e != nil {
		return fmt.Errorf("无法解析连接地址: %s", ctxUrl)
	}
	return nil
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

// 检查初始化flag参数
func checkInitRootFlags() {
	if *rootCtxRequestURL != "" {
		*rootCtxRequestURL = strings.TrimSuffix(*rootCtxRequestURL, "/")
		//校验 url是否符合规范
		if !strings.HasPrefix(*rootCtxRequestURL, "http://") && !strings.HasPrefix(*rootCtxRequestURL, "https://") {
			panic("'ctx-request-url' argument must start with either `http://` or `https://`")
		}

		e:=
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
func testCtxURL(ctxUrl string)error  {
	e:=networki
}

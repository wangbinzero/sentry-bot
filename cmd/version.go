package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var versionCmd = &cobra.Command{Use: "version",
	Short: "Version and build information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version:%s\n", version)
		fmt.Printf("  git branch: %s\n", gitBranch)
		fmt.Printf("  git hash: %s\n", gitHash)
		fmt.Printf("  build date: %s\n", buildDate)
		fmt.Printf("  env: %s\n", env)
		fmt.Printf("  GOOS: %s\n", runtime.GOOS)
		fmt.Printf("  GOARCH: %s\n", runtime.GOARCH)
	}}

// main.go
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sdkman-go/internal/registry"
)

func main() {
	var rootCmd = &cobra.Command{Use: "sdk"}
	registry.RegisterCommands(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// main.go
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sdkman-go/internal/registry"
	"strconv"
	"strings"
)

func main() {
	var rootCmd = &cobra.Command{Use: "sdk"}
	registry.RegisterCommands(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		splitErr := strings.Split(err.Error(), ":")
		exitCode, convErr := strconv.Atoi(splitErr[1])

		if convErr != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse exit code: %v", err)
			os.Exit(1)
		}

		fmt.Println(splitErr[0])

		os.Exit(exitCode)
	}
}

package handler

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"sdkman-go/internal/mock"
)

func UseHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) error {
			javaInstallPath := returnJavaPath(args)
			newJavaHome, err := setJavaPath(javaInstallPath)
			if err != nil {
				fmt.Println("Error setting Java path:", err)
				return errors.New("FAIL:1") // return custom exit code and error message
			}
			if shouldRefresh() {
				return errors.New("SUCCESS_REFRESH:100") // return custom exit code and error message
			}
			cmd.Println(newJavaHome)
			return errors.New("SUCCESS:0") // return success exit code and message
		},
	}
}

// Placeholder for your actual refresh logic.
func shouldRefresh() bool {
	// Implement your refresh logic here.
	return true
}

func returnJavaPath(args []string) string {
	javaRegistry := mock.NewJavaRegistry()
	for _, arg := range args {
		if _, ok := javaRegistry.Registry[arg]; ok {
			return javaRegistry.Registry[arg]
		}
	}
	return "Java Version Not Found Or Installed"
}

func setJavaPath(javaInstallPath string) (string, error) {
	psCommand := fmt.Sprintf("[System.Environment]::SetEnvironmentVariable('JAVA_HOME', \"%s\", 'User')", javaInstallPath)
	cmd := exec.Command("powershell", "-Command", psCommand)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute command: %w", err)
	}
	return "JAVA_HOME: " + javaInstallPath, nil
}

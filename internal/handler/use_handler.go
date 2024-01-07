package handler

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"sdkman-go/internal/mock"
)

func UseHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) {
			javaInstallPath := returnJavaPath(args)
			newJavaHome, _ := setJavaPath(javaInstallPath)
			cmd.Println(newJavaHome)
		},
	}
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

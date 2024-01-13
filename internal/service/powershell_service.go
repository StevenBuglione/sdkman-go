package service

import (
	"bytes"
	"fmt"
	"os/exec"
	"sdkman-go/internal/mock"
)

type IPowershellService interface {
	Execute(command string) error
	ExecuteJavaPathUpdate(args []string) error
}

type PowershellService struct {
}

func NewPowershellService() IPowershellService {
	return &PowershellService{}
}

func (p *PowershellService) Execute(command string) error {
	var stderr bytes.Buffer
	cmd := exec.Command("powershell", command)
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", err, stderr.String())
	}

	return nil
}

func (p *PowershellService) ExecuteJavaPathUpdate(args []string) error {
	javaInstallPath := returnJavaPath(args)

	javaInstallVar := "[System.Environment]::SetEnvironmentVariable"
	envVariables := fmt.Sprintf("'JAVA_HOME', \"%s\", 'User'", javaInstallPath)
	psCommand := fmt.Sprintf("%s(%s)", javaInstallVar, envVariables)

	return p.Execute(psCommand)
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

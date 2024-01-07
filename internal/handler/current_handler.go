package handler

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func CurrentHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) {
			javaVersionDetail := GetJavaVersionDetail()
			cmd.Println(javaVersionDetail)
		},
	}
}

func GetJavaVersionDetail() string {
	runtimeLine, err := GetJavaVersionLine()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	startIndex := strings.Index(runtimeLine, "Java(TM) SE Runtime Environment")
	endIndex := strings.Index(runtimeLine, "(build")

	if startIndex != -1 && endIndex != -1 {
		startIndex += len("Java(TM) SE Runtime Environment ")
		return strings.TrimSpace(runtimeLine[startIndex:endIndex])
	}

	return ""
}

func GetJavaVersionLine() (string, error) {
	cmd := exec.Command("java", "-version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "SE Runtime Environment") {
			return line, nil
		}
	}

	return "", fmt.Errorf("unable to find Runtime line")
}

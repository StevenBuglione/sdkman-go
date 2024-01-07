package exitcode

import (
	"fmt"
	"strconv"
)

type ExitCode string

func (e ExitCode) String() string {
	return string(e)
}

const (
	Success        ExitCode = "0"
	SuccessRefresh ExitCode = "100"
	Failure        ExitCode = "1"
)

func ParseExitCode(exitCode string) int {
	num, err := strconv.Atoi(exitCode)
	if err != nil {
		fmt.Printf("Error converting string to integer: %v", err)
		return 1
	}
	return num
}

func IsSuccess() string {
	return Success.String()
}

func IsSuccessRefresh() string {
	return SuccessRefresh.String()
}

func IsFailure() string {
	return Failure.String()
}

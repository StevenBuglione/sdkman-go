package exitcode

type ExitCode int

func (e ExitCode) Value() int {
	return int(e)
}

const (
	Success        ExitCode = 0
	SuccessRefresh ExitCode = 100
	Failure        ExitCode = 1
)

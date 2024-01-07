package exitcode

type ExitCode int

const (
	SUCCESS ExitCode = iota
	FAIL
	SUCCESS_REFRESH
)

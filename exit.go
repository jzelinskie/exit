// Package exit implements an ergonomic mapping from program errors into their
// respective exit codes.
//
// Consider this package a slightly higher-level `os.Exit()`.
package exit

import "os"

// BSD's sysexits.h
const (
	ExUsage       = 64
	ExDataErr     = 65
	ExNoInput     = 66
	ExNoUser      = 67
	ExNoHost      = 68
	ExUnavailable = 69
	ExSoftware    = 70
	ExOSErr       = 71
	ExOSFile      = 72
	ExCantCreat   = 73
	ExIOErr       = 74
	ExTempFail    = 75
	ExProtocol    = 76
	ExNoPerm      = 77
	ExConfig      = 78
)

// Error is representation of any failure that can be represented as an exit
// code.
type Error interface {
	ExitCode() int
}

// UsageError wraps any error that is caused by improper usage by the program
// user.
type UsageError struct{ error }

func (e UsageError) ExitCode() int {
	if e.Error != nil {
		return ExUsage
	}
	return 0
}

// OSError wraps any error that is caused by an OS failure.
type OSError struct{ error }

func (e OSError) ExitCode() int {
	if e.Error != nil {
		return ExOSErr
	}
	return 0
}

// ConfigError wraps any error that is caused by invalid program configuration.
type ConfigError struct{ error }

func (e ConfigError) ExitCode() int {
	if e.Error != nil {
		return ExConfig
	}
	return 0
}

// Now calls os.Exit with the proper exit code for the provided error.
func Now(err Error) {
	os.Exit(err.ExitCode())
}

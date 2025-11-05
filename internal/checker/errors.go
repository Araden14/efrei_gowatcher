package checker

import "fmt"

// UnreachableError is an error that occurs when a URL is unreachable.
type UnreachableError struct {
	URL string
	err error
}

func (e *UnreachableError) Error() string {
	return fmt.Sprintf("URL %s is unreachable: %v", e.URL, e.err)
}

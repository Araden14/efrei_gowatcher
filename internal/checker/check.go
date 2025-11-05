package checker

import (
	"net/http"
	"time"
)

type CheckResult struct {
	InputTarget config.InputTarget
	Status      string
	Err         error
}

type ReportEntry struct {
	Name   string
	URL    string
	Owner  string
	Status string // OK, INNACCESIBLE, ERROR
	ErrMsg string // message d'erreur , optionnel
}

func CheckURL(url config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return CheckResult{
			InputTarget: url,
			Err: &UnreachableError{
				URL: url,
				err: err,
			},
		}
	}
	defer resp.Body.Close()
	return CheckResult{
		InputTarget: url,
		Status:      resp.Status,
		Err:         nil,
	}
}

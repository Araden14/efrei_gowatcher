package checker

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Araden14/efrei_gowatcher/internal/config"
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

func CheckURL(target config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(target.URL)
	if err != nil {
		return CheckResult{
			InputTarget: target,
			Err: &UnreachableError{
				URL: target.URL,
				err: err,
			},
		}
	}
	defer resp.Body.Close()
	return CheckResult{
		InputTarget: target,
		Status:      resp.Status,
		Err:         nil,
	}
}

func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry{
		Name:   res.InputTarget.Name,
		URL:    res.InputTarget.URL,
		Owner:  res.InputTarget.Owner,
		Status: res.Status,
	}
	if res.Err != nil {
		var unreachable *UnreachableError
		if errors.As(res.Err, &unreachable) {
			report.Status = "INNACCESIBLE"
			report.ErrMsg = fmt.Sprintf("URL %s is unreachable: %v", unreachable.URL, unreachable.err)
		} else {
			report.Status = "ERROR"
			report.ErrMsg = fmt.Sprintf("Erreur générique : %v", res.Err)
		}
	}
	return report
}

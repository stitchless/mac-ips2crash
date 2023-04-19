package ips2crash

import (
	"errors"

	"github.com/jpeizer/mac-ips2crash/internal/app/ips2crash"
)

func ProcessCrashReport(inputBytes []byte) (ips2crash.IPSCrash, error) {
	inputFile, err := ips2crash.NewCrashReport(inputBytes)
	if err != nil {
		return ips2crash.IPSCrash{}, err
	}

	if inputFile.Header == "" || inputFile.Payload == "" {
		return ips2crash.IPSCrash{}, errors.New("invalid crash report")
	}

	crashReport, err := ips2crash.ParseCrashReport(inputFile)
	if err != nil {
		return ips2crash.IPSCrash{}, err
	}

	crashReport.FormattedReport = crashReport.Payload.Format()

	crashReport.Output = ips2crash.CleanOutput(crashReport)

	return crashReport, nil
}

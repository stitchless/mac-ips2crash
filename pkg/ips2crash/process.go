package ips2crash

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/jpeizer/mac-ips2crash/internal/app/ips2crash"
)

func ProcessCrashReport(crashPath string) (ips2crash.IPSCrash, error) {
	file, err := os.ReadFile(crashPath)
	if err != nil {
		return ips2crash.IPSCrash{}, err
	}

	inputFile, err := ips2crash.NewCrashReport(file)
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

	fileName := filepath.Base(crashPath)
	crashReport.FileName = fileName
	crashReport.FilePath = crashPath

	crashReport.CrashFormat = crashReport.Payload.Format()

	return crashReport, nil
}

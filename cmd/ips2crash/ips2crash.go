package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jpeizer/mac-ips2crash/pkg/ips2crash"
)

func main() {
	var ipsFilePath, outputPath string
	// var listenToStdin bool

	flag.StringVar(&ipsFilePath, "i", "", "input path to ips file")
	flag.StringVar(&outputPath, "o", "", "output path to save crash report")
	// flag.BoolVar(&listenToStdin, "s", false, "listen to stdin for ips data")

	// print usage if no flags are provided
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Parse the runtime flags
	flag.Parse()

	if len(os.Args) > 1 && os.Args[1] == "help" {
		flag.Usage()
		return
	}

	crashReport, err := ips2crash.ProcessCrashReport(ipsFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := crashReport.FileName
	fileName = strings.Replace(fileName, filepath.Ext(fileName), ".crash", 1)

	if outputPath == "" {
		outputPath = filepath.Join(filepath.Dir(ipsFilePath), fileName)
	} else {
		outputPath = filepath.Join(outputPath, fileName)
	}

	err = os.WriteFile(outputPath, []byte(crashReport.CrashFormat), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Crash report saved to:", outputPath)
}

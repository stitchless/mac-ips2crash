# mac-ips2crash

[![Go IPS2Crash Checks](https://github.com/jpeizer/mac-ips2crash/actions/workflows/status.yml/badge.svg)](https://github.com/jpeizer/mac-ips2crash/actions/workflows/status.yml)

This is a simple tool to convert a macOS IPS file into a .crash report
that can be used with legacy tools and pipelines to symbolicate the crash.  It can be used as a standalone tool or as a library.

This is my first open source project I am attempting to make for public use, so if you run into any issues, don't hesitate to open an issue or PR.

## Tool Usage
```shell
# Tool Release Binaries are available on the releases page
# Usage as a standalone tool
Usage of ./mac-ips2crash:
  -i string
        Path to IPS file
  -o string
        Path to output crash file
        The file will replace any extension with .crash
```

# Library Usage
```shell
go get github.com/jpeizer/mac-ips2crash
```
```go
package main

import (
    "fmt"
    "os"

    "github.com/jpeizer/mac-ips2crash"
)

func main() {
    ipsFilePath := "/path/to/ips/file"
    file, err := os.ReadFile(ipsFilePath)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    // However you have a []byte of the IPS file
    crashReport, err := ips2crash.ProcessCrashReport(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    for _, line := range crashReport.LeadingText {
        fmt.Println(line) // Prints any test that appears before the IPS formatted file
    }
    
    // Prints any text that appears after the IPS formatted file
    for _, line := range crashReport.TrailingText {
        fmt.Println(line)
    }
    
    // Prints the first JSON object in the IPS file (the header)
    fmt.Println(crashReport.Header)
    
    // Prints the second JSON object in the IPS file (the Payload)
    fmt.Println(crashReport.Payload)
    
    // Prints the formatted report minus the leading and trailing text
    fmt.Println(crashReport.FormattedReport) 
    
    // Prints the formatted report with the leading and trailing text
    fmt.Println(crashReport.Output) 
}
```

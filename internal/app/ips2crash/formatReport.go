package ips2crash

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (ips IPSPayload) Format() string {
	caser := cases.Upper(language.English)

	var sb strings.Builder

	if ips.ProcName != "" {
		if ips.Pid != 0 {
			sb.WriteString(fmt.Sprintf("%-23s %s [%d]\n", "Process:", ips.ProcName, ips.Pid))
		} else {
			sb.WriteString(fmt.Sprintf("%-23s %s\n", "Process:", ips.ProcName))
		}
	}

	if ips.ProcPath != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "Path:", ips.ProcPath))
	}

	if ips.BundleInfo.CFBundleIdentifier != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "Identifier:", ips.BundleInfo.CFBundleIdentifier))
	}

	if ips.BundleInfo.CFBundleVersion != "" {
		if ips.BundleInfo.CFBundleShortVersionString != "" {
			sb.WriteString(fmt.Sprintf("%-23s %s (%s)\n", "Version:", ips.BundleInfo.CFBundleShortVersionString, ips.BundleInfo.CFBundleVersion))
		} else {
			sb.WriteString(fmt.Sprintf("%-23s %s\n", "Version:", ips.BundleInfo.CFBundleVersion))
		}
	}

	if ips.CpuType != "" {
		if ips.Translated == false {
			sb.WriteString(fmt.Sprintf("%-23s %s (Native)\n", "Code Type:", ips.CpuType))
		} else {
			sb.WriteString(fmt.Sprintf("%-23s %s\n", "Code Type:", ips.CpuType))
		}
	}

	if ips.ParentProc != "" {
		if ips.ParentPid != 0 {
			sb.WriteString(fmt.Sprintf("%-23s %s [%d]\n", "Parent Process:", ips.ParentProc, ips.ParentPid))
		} else {
			sb.WriteString(fmt.Sprintf("%-23s %s\n", "Parent Process:", ips.ParentProc))
		}
	}

	if ips.UserID != -1 {
		sb.WriteString(fmt.Sprintf("%-23s %d\n", "User ID:", ips.UserID))
	}

	sb.WriteString("\n")

	if ips.ModelCode != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "Hardware Model:", ips.ModelCode))
	}

	if ips.CaptureTime != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "Date/Time:", ips.CaptureTime))
	}

	if ips.OsVersion.Train != "" {
		if ips.OsVersion.Build != "" {
			sb.WriteString(fmt.Sprintf("%-23s %s (%s)\n", "OS Version:", ips.OsVersion.Train, ips.OsVersion.Build))
		} else {
			sb.WriteString(fmt.Sprintf("%-23s %s\n", "OS Version:", ips.OsVersion.Train))
		}
	}

	sb.WriteString(fmt.Sprintf("%-23s %s\n", "Report Version:", "12"))

	if ips.OsVersion.ReleaseType != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "Release Type:", ips.OsVersion.ReleaseType))
	}

	if ips.Incident != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "Anonymous UUID:", ips.CrashReporterKey))
	}

	sb.WriteString("\n")

	if ips.SleepWakeUUID != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "Sleep/Wake UUID:", ips.SleepWakeUUID))
		sb.WriteString("\n")
	}

	if ips.Uptime != 0 {
		sb.WriteString(fmt.Sprintf("%-23s %d seconds\n", "Time Awake Since Boot:", ips.Uptime))
	}

	if ips.WakeTime != 0 {
		sb.WriteString(fmt.Sprintf("%-23s %d seconds\n", "Time Since Wake:", ips.WakeTime))
	}

	// if ips.ProcLaunch != "" {
	// sb.WriteString(fmt.Sprintf("%-23s %s\n", "Launch Time:", ips.ProcLaunch))
	// }

	if ips.Uptime != 0 || ips.WakeTime != 0 || ips.ProcLaunch != "" {
		sb.WriteString("\n")
	}

	if ips.Sip != "" {
		sb.WriteString(fmt.Sprintf("%-23s %s\n", "System Integrity Protection:", ips.Sip))
		sb.WriteString("\n")
	}

	for i, thread := range ips.Threads {
		if thread.Triggered {
			if thread.Name != "" {
				if thread.Queue != "" {
					sb.WriteString(fmt.Sprintf("%-23s %d  %s  Dispatch queue: %s\n", "Crashed Thread:", i, thread.Name, thread.Queue))
				} else {
					sb.WriteString(fmt.Sprintf("%-23s %d  %s\n", "Crashed Thread:", i, thread.Name))
				}
			} else {
				sb.WriteString(fmt.Sprintf("%-23s %d\n", "Crashed Thread:", i))
			}
			sb.WriteString("\n")
			break
		}
	}

	if ips.Exception.Type != "" {
		if ips.Exception.Signal != "" {
			sb.WriteString(fmt.Sprintf("%-19s%s (%s)\n", "Exception Type:", ips.Exception.Type, ips.Exception.Signal))
		} else {
			sb.WriteString(fmt.Sprintf("%-19s%s\n", "Exception Type:", ips.Exception.Type))
		}
	}

	// fist set of codes
	if ips.Exception.Subtype != "" {
		sb.WriteString(fmt.Sprintf("%-19s%s\n", "Exception Codes:", ips.Exception.Subtype))
	}

	// second set of codes
	if ips.Exception.Codes != "" {
		sb.WriteString(fmt.Sprintf("%-19s%s\n", "Exception Codes:", ips.Exception.Codes))
	}

	if ips.Exception.Message != "" {
		sb.WriteString(fmt.Sprintf("%-19s%s\n", "Exception Message:", ips.Exception.Message))
	}

	if ips.IsCorpse == 1 {
		sb.WriteString(fmt.Sprintf("%-19s %s\n", "Exception Note:", "EXC_CORPSE_NOTIFY"))
	}

	sb.WriteString("\n")

	if ips.Termination.Namespace != "" {
		if ips.Termination.Indicator != "" {
			sb.WriteString(fmt.Sprintf("%-23s Namespace Signal, %s %d %s\n", "Termination Reason:", ips.Termination.Namespace, ips.Termination.Code, ips.Termination.Indicator))
		} else {
			sb.WriteString(fmt.Sprintf("%-23s Namespace Signal, %s Code %d\n", "Termination Reason:", ips.Termination.Namespace, ips.Termination.Code))
		}
	}

	if ips.Termination.ByProc != "" {
		if ips.Termination.ByPid != 0 {
			sb.WriteString(fmt.Sprintf("%-23s %s [%d]\n", "Terminating Process:", ips.Termination.ByProc, ips.Termination.ByPid))
		} else {
			sb.WriteString(fmt.Sprintf("%-23s %s\n", "Terminating Process:", ips.Termination.ByProc))
		}
	}

	if ips.Termination.ByProc != "" || ips.Termination.Namespace != "" {
		sb.WriteString("\n")
	}

	if ips.VmRegionInfo != "" {
		sb.WriteString(fmt.Sprintf("VM Region Info: %s\n\n", ips.VmRegionInfo))
	}

	if len(ips.AsiSignatures) > 0 {
		signatures := make([]string, 0)
		for _, signature := range ips.AsiSignatures {
			signatures = append(signatures, signature)
		}

		sb.WriteString("Application Specific Signatures:\n")
		sb.WriteString(fmt.Sprintf("%s\n", strings.Join(signatures, ", ")))
	}

	if ips.Asi != nil {
		asiValues := make([]string, 0)

		// Check if the ASI is a string
		if asiString, ok := ips.Asi.(string); ok {
			asiValues = append(asiValues, asiString)
		}

		// Check if the ASI is a slice
		if asiSlice, ok := ips.Asi.([]interface{}); ok {
			for _, asi := range asiSlice {
				asiValues = append(asiValues, fmt.Sprintf("%s", asi))
			}
		}

		// Check if the ASI is a map
		if asiMap, ok := ips.Asi.(map[string]interface{}); ok {
			for _, asi := range asiMap {
				// check if asi is a slice
				if asiSlice, ok := asi.([]interface{}); ok {
					for _, asi := range asiSlice {
						asiValues = append(asiValues, fmt.Sprintf("%s", asi))
					}
				} else {
					asiValues = append(asiValues, fmt.Sprintf("%s", asi))
				}
			}
		}

		sb.WriteString(fmt.Sprintf("%s\n%s\n", "Application Specific Information:", strings.Join(asiValues, ", ")))
	}

	if len(ips.AsiSignatures) > 0 || ips.Asi != nil {
		sb.WriteString("\n")
	}

	if len(ips.LastExceptionBacktrace) > 0 {
		sb.WriteString(fmt.Sprintf("Application Specific Backtrace 0:\n"))
		for i, frame := range ips.LastExceptionBacktrace {
			baseInt := ips.UsedImages[frame.ImageIndex].Base
			targetInt := baseInt - int64(frame.ImageOffset)
			sb.WriteString(fmt.Sprintf("%-3d %-32s %#016x %s + %d\n", i, ips.UsedImages[frame.ImageIndex].Name, targetInt, frame.Symbol, frame.SymbolLocation))
		}

		sb.WriteString("\n")
	}

	if ips.KernelTriage != "" {
		sb.WriteString(fmt.Sprintf("%s\n%s\n", "Kernel Triage:", ips.KernelTriage))
		sb.WriteString("\n")
	}

	// ------ Threads ------

	for ti, thread := range ips.Threads {
		threadString := fmt.Sprintf("Thread %d", ti)
		if thread.Triggered {
			threadString += " Crashed"
		}
		if thread.Name != "" {
			threadString += ":: " + thread.Name
		} else {
			threadString += ":"
		}

		sb.WriteString(threadString + "\n")

		for i, frame := range thread.Frames {
			baseInt := ips.UsedImages[frame.ImageIndex].Base
			offsetInt := frame.ImageOffset
			targetInt := baseInt + int64(offsetInt)

			if frame.Symbol != "" {
				if frame.SymbolLocation != 0 {
					sb.WriteString(fmt.Sprintf("%-3d %-32s %#016x %s + %d\n", i, ips.UsedImages[frame.ImageIndex].Name, targetInt, frame.Symbol, frame.SymbolLocation))
					continue
				}
				sb.WriteString(fmt.Sprintf("%-3d %-32s %#016x %s\n", i, ips.UsedImages[frame.ImageIndex].Name, targetInt, frame.Symbol))
				continue
			}
			sb.WriteString(fmt.Sprintf("%-3d %-32s %#016x %#x + %d\n", i, ips.UsedImages[frame.ImageIndex].Name, targetInt, baseInt, frame.ImageOffset))
		}
		sb.WriteString("\n")
	}

	// ------ Thread State ------
	for ti, thread := range ips.Threads {
		if thread.Triggered {
			threadState := thread.ThreadState
			flavor := threadState.Flavor
			is64Bit := strings.Contains(ips.CpuType, "64")
			flavor = strings.Split(flavor, "_")[0]
			flavor = caser.String(flavor)

			archLabel := ""

			if is64Bit {
				archLabel = " (64-bit)"
			} else {
				archLabel = "32-bit"
			}

			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("Thread %d crashed with %s Thread State%s:\n ", ti, flavor, archLabel))
			registerPosition := 0

			// I think this is reserved for ARM based devices
			if threadState.X != nil {
				for i, reg := range threadState.X {
					label := fmt.Sprintf("x%d", i)
					sb.WriteString(fmt.Sprintf("%5s: %#016x ", label, reg.Value))
					registerPosition++
					if registerPosition%4 == 0 {
						sb.WriteString("\n ")
						registerPosition = 0
					}
				}

				if registerPosition%3 == 0 {
					sb.WriteString("\n ")
					registerPosition = 0
				}
				sb.WriteString(fmt.Sprintf("%5s: %#016x ", "fp", threadState.Fp.Value))
				sb.WriteString(lineBreakByNumber(&registerPosition, 3))
				sb.WriteString(fmt.Sprintf("%5s: %#016x ", "lr", threadState.Lr.Value))
				sb.WriteString(lineBreakByNumber(&registerPosition, 3))
				sb.WriteString(fmt.Sprintf("%5s: %#016x ", "sp", threadState.Sp.Value))
				sb.WriteString(lineBreakByNumber(&registerPosition, 3))
				sb.WriteString(fmt.Sprintf("%5s: %#016x ", "pc", threadState.Pc.Value))
				sb.WriteString(lineBreakByNumber(&registerPosition, 3))
				sb.WriteString(fmt.Sprintf("%5s: %#016x ", "cpsr", threadState.Cpsr.Value))
				sb.WriteString(lineBreakByNumber(&registerPosition, 3))
				sb.WriteString(fmt.Sprintf("%5s: %#016x ", "far", threadState.Far.Value))
				sb.WriteString(lineBreakByNumber(&registerPosition, 3))
				sb.WriteString(fmt.Sprintf("%5s: %#016x %s", "esr", threadState.Esr.Value, threadState.Esr.Description))
				sb.WriteString(lineBreakByNumber(&registerPosition, 3))
				sb.WriteString("\n")
			} else {
				// TODO: Add support for other architectures
			}
			sb.WriteString("\n")
		}
	}

	// ------ Binary Images ------

	sb.WriteString("Binary Images:\n")

	for _, image := range ips.UsedImages {
		baseInt := image.Base
		sizeInt := image.Size
		endInt := baseInt + int64(sizeInt)
		baseHex := fmt.Sprintf("%#x", baseInt)
		endHex := fmt.Sprintf("%#x", endInt)

		var version string
		if image.CFBundleShortVersionString != "" && image.CFBundleVersion != "" {
			version = image.CFBundleShortVersionString + " - " + image.CFBundleVersion
		} else if image.CFBundleShortVersionString != "" {
			version = image.CFBundleShortVersionString
		} else if image.CFBundleVersion != "" {
			version = image.CFBundleVersion
		} else {
			version = "*"
		}

		var symbolID string
		if image.CFBundleIdentifier != "" {
			symbolID = image.CFBundleIdentifier
		} else {
			symbolID = image.Name
		}

		sb.WriteString(fmt.Sprintf("%18s - %18s %s %s (%s) <%s> %s\n", baseHex, endHex, symbolID, image.Arch, version, strings.ToLower(image.Uuid), image.Path))
	}

	sb.WriteString("\n")

	// TODO: Check if struct is empty
	sb.WriteString("External Modification Summary:\n")
	sb.WriteString(fmt.Sprintf("  Calls made by other processes targeting this process:\n"))
	sb.WriteString(fmt.Sprintf("    task_for_pid: %d\n", ips.ExtMods.Targeted.TaskForPid))
	sb.WriteString(fmt.Sprintf("    thread_create: %d\n", ips.ExtMods.Targeted.ThreadCreate))
	sb.WriteString(fmt.Sprintf("    thread_set_state: %d\n", ips.ExtMods.Targeted.ThreadSetState))
	sb.WriteString(fmt.Sprintf("  Calls made by this process:\n"))
	sb.WriteString(fmt.Sprintf("    task_for_pid: %d\n", ips.ExtMods.Caller.TaskForPid))
	sb.WriteString(fmt.Sprintf("    thread_create: %d\n", ips.ExtMods.Caller.ThreadCreate))
	sb.WriteString(fmt.Sprintf("    thread_set_state: %d\n", ips.ExtMods.Caller.ThreadSetState))
	sb.WriteString(fmt.Sprintf("  Calls made by all processes on this machine:\n"))
	sb.WriteString(fmt.Sprintf("    task_for_pid: %d\n", ips.ExtMods.System.TaskForPid))
	sb.WriteString(fmt.Sprintf("    thread_create: %d\n", ips.ExtMods.System.ThreadCreate))
	sb.WriteString(fmt.Sprintf("    thread_set_state: %d\n", ips.ExtMods.System.ThreadSetState))
	sb.WriteString("\n")

	if ips.VMSummary != "" {
		sb.WriteString(fmt.Sprintf("%s\n%s\n", "VM Region Summary:", ips.VMSummary))
		sb.WriteString("\n")
	}

	return sb.String()
}

func lineBreakByNumber(index *int, nth int) string {
	if (*index+1)%nth == 0 {
		*index = 0
		return "\n "
	}

	*index++
	return ""
}

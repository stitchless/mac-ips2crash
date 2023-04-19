package ips2crash

type IPSCrash struct {
	FileName        string
	FilePath        string
	LeadingText     []string   `json:"leading_text"`
	Header          IPSHeader  `json:"header"`
	Payload         IPSPayload `json:"payload"`
	TrailingText    []string   `json:"trailing_text"`
	FormattedReport string     `json:"crash_format"`
	Output          string     `json:"output"`
}

type IPSHeader struct {
	AppName          string `json:"app_name"`
	Timestamp        string `json:"timestamp"`
	AppVersion       string `json:"app_version"`
	SliceUuid        string `json:"slice_uuid"`
	BuildVersion     string `json:"build_version"`
	Platform         int    `json:"platform"`
	BundleID         string `json:"bundleID"`
	ShareWithAppDevs int    `json:"share_with_app_devs"`
	IsFirstParty     int    `json:"is_first_party"`
	BugType          string `json:"bug_type"`
	OsVersion        string `json:"os_version"`
	RootsInstalled   int    `json:"roots_installed"`
	Name             string `json:"name"`
	IncidentId       string `json:"incident_id"`
}

type IPSPayload struct {
	Uptime        int    `json:"uptime"`
	ProcRole      string `json:"procRole"`
	Version       int    `json:"version"`
	UserID        int    `json:"userID"`
	DeployVersion int    `json:"deployVersion"`
	ModelCode     string `json:"modelCode"`
	CoalitionID   int    `json:"coalitionID"`
	OsVersion     struct {
		Train       string `json:"train"`
		Build       string `json:"build"`
		ReleaseType string `json:"releaseType"`
	} `json:"osVersion"`
	CaptureTime      string `json:"captureTime"`
	Incident         string `json:"incident"`
	Pid              int    `json:"pid"`
	IsCorpse         int    `json:"isCorpse"`
	CpuType          string `json:"cpuType"`
	RootsInstalled   int    `json:"roots_installed"`
	BugType          string `json:"bug_type"`
	IsNonFatal       bool   `json:"isNonFatal"`
	IsSimulated      bool   `json:"isSimulated"`
	ProcLaunch       string `json:"procLaunch"`
	ProcStartAbsTime int64  `json:"procStartAbsTime"`
	ProcExitAbsTime  int64  `json:"procExitAbsTime"`
	ProcName         string `json:"procName"`
	ProcPath         string `json:"procPath"`
	VMSummary        string `json:"vmSummary"`
	Translated       bool   `json:"translated"`
	BundleInfo       struct {
		CFBundleShortVersionString string `json:"CFBundleShortVersionString"`
		CFBundleVersion            string `json:"CFBundleVersion"`
		CFBundleIdentifier         string `json:"CFBundleIdentifier"`
	} `json:"bundleInfo"`
	StoreInfo struct {
		DeviceIdentifierForVendor string `json:"deviceIdentifierForVendor"`
		ThirdParty                bool   `json:"thirdParty"`
	} `json:"storeInfo"`
	ParentProc       string `json:"parentProc"`
	ParentPid        int    `json:"parentPid"`
	CoalitionName    string `json:"coalitionName"`
	CrashReporterKey string `json:"crashReporterKey"`
	WakeTime         int    `json:"wakeTime"`
	SleepWakeUUID    string `json:"sleepWakeUUID"`
	BridgeVersion    struct {
		Build string `json:"build"`
		Train string `json:"train"`
	} `json:"bridgeVersion"`
	Sip          string `json:"sip"`
	VmRegionInfo string `json:"vmRegionInfo"`
	Exception    struct {
		Codes    string `json:"codes"`
		RawCodes []int  `json:"rawCodes"`
		Type     string `json:"type"`
		Signal   string `json:"signal"`
		Subtype  string `json:"subtype"`
		Message  string `json:"message"`
	} `json:"exception"`
	Termination struct {
		Flags     int    `json:"flags"`
		Code      int    `json:"code"`
		Namespace string `json:"namespace"`
		Indicator string `json:"indicator"`
		ByProc    string `json:"byProc"`
		ByPid     int    `json:"byPid"`
	} `json:"termination"`
	Vmregioninfo           string      `json:"vmregioninfo"`
	AsiSignatures          []string    `json:"asiSignatures"`
	Asi                    interface{} `json:"asi,omitempty"`
	LastExceptionBacktrace []struct {
		ImageOffset    int    `json:"imageOffset"`
		Symbol         string `json:"symbol,omitempty"`
		SymbolLocation int    `json:"symbolLocation,omitempty"`
		ImageIndex     int    `json:"imageIndex"`
	} `json:"lastExceptionBacktrace,omitempty"`
	KernelTriage string `json:"ktriageinfo,omitempty"`
	ExtMods      struct {
		Caller struct {
			ThreadCreate   int `json:"thread_create"`
			ThreadSetState int `json:"thread_set_state"`
			TaskForPid     int `json:"task_for_pid"`
		} `json:"caller"`
		System struct {
			ThreadCreate   int `json:"thread_create"`
			ThreadSetState int `json:"thread_set_state"`
			TaskForPid     int `json:"task_for_pid"`
		} `json:"system"`
		Targeted struct {
			ThreadCreate   int `json:"thread_create"`
			ThreadSetState int `json:"thread_set_state"`
			TaskForPid     int `json:"task_for_pid"`
		} `json:"targeted"`
		Warnings int `json:"warnings"`
	} `json:"extMods"`
	FaultingThread int `json:"faultingThread"`
	Threads        []struct {
		Triggered        bool   `json:"triggered,omitempty"`
		Id               int    `json:"id"`
		Queue            string `json:"queue,omitempty"`
		InstructionState struct {
			InstructionStream struct {
				Bytes  []int `json:"bytes"`
				Offset int   `json:"offset"`
			} `json:"instructionStream"`
		} `json:"instructionState,omitempty"`
		Frames []struct {
			ImageOffset    int    `json:"imageOffset"`
			ImageIndex     int    `json:"imageIndex"`
			Symbol         string `json:"symbol,omitempty"`
			SymbolLocation int    `json:"symbolLocation,omitempty"`
		} `json:"frames"`

		ThreadState struct {
			Flavor string `json:"flavor"`
			R13    struct {
				Value uint64 `json:"value"`
			} `json:"r13"`
			Rax struct {
				Value uint64 `json:"value"`
			} `json:"rax"`
			Rflags struct {
				Value uint64 `json:"value"`
			} `json:"rflags"`
			Cpu struct {
				Value uint64 `json:"value"`
			} `json:"cpu"`
			R14 struct {
				Value uint64 `json:"value"`
			} `json:"r14"`
			Rsi struct {
				Value uint64 `json:"value"`
			} `json:"rsi"`
			R8 struct {
				Value uint64 `json:"value"`
			} `json:"r8"`
			Cr2 struct {
				Value uint64 `json:"value"`
			} `json:"cr2"`
			Rdx struct {
				Value uint64 `json:"value"`
			} `json:"rdx"`
			R10 struct {
				Value uint64 `json:"value"`
			} `json:"r10"`
			R9 struct {
				Value uint64 `json:"value"`
			} `json:"r9"`
			R15 struct {
				Value uint64 `json:"value"`
			} `json:"r15"`
			Rbx struct {
				Value uint64 `json:"value"`
			} `json:"rbx"`
			Trap struct {
				Value uint64 `json:"value"`
			} `json:"trap"`
			Err struct {
				Value uint64 `json:"value"`
			} `json:"err"`
			R11 struct {
				Value uint64 `json:"value"`
			} `json:"r11"`
			Rip struct {
				Value uint64 `json:"value"`
			} `json:"rip"`
			Rbp struct {
				Value uint64 `json:"value"`
			} `json:"rbp"`
			Rsp struct {
				Value uint64 `json:"value"`
			} `json:"rsp"`
			R12 struct {
				Value uint64 `json:"value"`
			} `json:"r12"`
			Rcx struct {
				Value uint64 `json:"value"`
			} `json:"rcx"`
			Rdi struct {
				Value uint64 `json:"value"`
			} `json:"rdi"`
			X []struct {
				Value uint64 `json:"value"`
			} `json:"x,omitempty"`
			Fp struct {
				Value uint64 `json:"value"`
			}
			Lr struct {
				Value uint64 `json:"value"`
			}
			Sp struct {
				Value uint64 `json:"value"`
			}
			Pc struct {
				Value uint64 `json:"value"`
			}
			Cpsr struct {
				Value uint64 `json:"value"`
			}
			Far struct {
				Value uint64 `json:"value"`
			}
			Esr struct {
				Value       uint64 `json:"value"`
				Description string `json:"description"`
			}
		} `json:"threadState,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"threads"`
	UsedImages []struct {
		Source                     string `json:"source"`
		Arch                       string `json:"arch"`
		Base                       int64  `json:"base"`
		CFBundleShortVersionString string `json:"CFBundleShortVersionString,omitempty"`
		CFBundleIdentifier         string `json:"CFBundleIdentifier,omitempty"`
		Size                       int    `json:"size"`
		Uuid                       string `json:"uuid"`
		Path                       string `json:"path"`
		Name                       string `json:"name"`
		CFBundleVersion            string `json:"CFBundleVersion,omitempty"`
	} `json:"usedImages"`
	SharedCache struct {
		Base int64  `json:"base"`
		Size int64  `json:"size"`
		Uuid string `json:"uuid"`
	} `json:"sharedCache"`
	LegacyInfo struct {
		ThreadTriggered struct {
			Name  string `json:"name"`
			Queue string `json:"queue"`
		} `json:"threadTriggered"`
	} `json:"legacyInfo"`
	TrialInfo struct {
		Rollouts []struct {
			RolloutId     string `json:"rolloutId"`
			FactorPackIds struct {
				SIRIUNDERSTANDINGTMDC     string `json:"SIRI_UNDERSTANDING_TMDC,omitempty"`
				HEALTHFEATUREAVAILABILITY string `json:"HEALTH_FEATURE_AVAILABILITY,omitempty"`
			} `json:"factorPackIds"`
			DeploymentId int `json:"deploymentId"`
		} `json:"rollouts"`
		Experiments []interface{} `json:"experiments"`
	} `json:"trialInfo"`
}

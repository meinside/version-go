package version

import (
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

const (
	HashLength = 7

	UnknownVersion = "unknown"
)

const (
	None         = 0
	OS           = 1 << iota
	Architecture = 1 << iota
	Revision     = 1 << iota
	Time         = 1 << iota
	Modification = 1 << iota

	All = OS | Architecture | Revision | Time | Modification
)

// Full returns the build version with all flags.
func Full() string {
	return Build(All)
}

// Minimum returns the minimum build version. (debug.BuildInfo.Main.Version only)
func Minimum() string {
	return Build(None)
}

// Build returns the build version of the application with given flag.
func Build(flag int) string {
	if info, success := debug.ReadBuildInfo(); success {
		builds := []string{info.Main.Version}

		if flag&OS > 0 {
			builds = append(builds, runtime.GOOS)
		}
		if flag&Architecture > 0 {
			builds = append(builds, runtime.GOARCH)
		}

		for _, kv := range info.Settings {
			switch kv.Key {
			case "vcs.revision":
				if flag&Revision > 0 {
					revision := kv.Value
					builds = append(builds, revision[:HashLength])
				}
			case "vcs.time":
				if flag&Time > 0 {
					lastCommit, _ := time.Parse(time.RFC3339, kv.Value)
					builds = append(builds, lastCommit.Format("20060102_150405"))
				}
			case "vcs.modified":
				if flag&Modification > 0 {
					if kv.Value == "true" {
						builds = append(builds, "(modified)")
					}
				}
			}
		}

		return strings.Join(builds, "-")
	}

	return UnknownVersion
}

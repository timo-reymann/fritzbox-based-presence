package buildinfo

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"
)

// GitSha of the build
var GitSha = "unknown"

// Version contains the latest version tag or question mark for source builds
var Version = "?"

// BuildTime contains the build time or question mark for source builds
var BuildTime = "?"

func addLine(w *tabwriter.Writer, heading string, val string) {
	_, _ = fmt.Fprintf(w, heading+"\t%s\n", val)
}

// PrintVersionInfo prints a tabular list with build info
func PrintVersionInfo() {
	fmt.Printf("fritzbox-based-presence %s (%s) by Timo Reymann\n", Version, BuildTime)
	println()
	println("Build information")
	w := tabwriter.NewWriter(os.Stderr, 10, 1, 10, byte(' '), tabwriter.TabIndent)
	addLine(w, "GitSha", GitSha)
	addLine(w, "Version", Version)
	addLine(w, "BuildTime", BuildTime)
	addLine(w, "Go-Version", runtime.Version())
	addLine(w, "OS/Arch", runtime.GOOS+"/"+runtime.GOARCH)
	_ = w.Flush()
}

package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func formatBytes(n int64) string {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
	)
	switch {
	case n >= gb:
		return fmt.Sprintf("%.1fGB", float64(n)/float64(gb))
	case n >= mb:
		return fmt.Sprintf("%.1fMB", float64(n)/float64(mb))
	case n >= kb:
		return fmt.Sprintf("%.1fKB", float64(n)/float64(kb))
	default:
		return fmt.Sprintf("%dB", n)
	}
}

func printDownloadProgress(written, total int64, filename string) {
	const width = 20

	var bar string
	var pct int

	if total > 0 {
		pct = int(written * 100 / total)
		if pct > 100 {
			pct = 100
		}
		filled := pct * width / 100
		bar = strings.Repeat("#", filled) + strings.Repeat("-", width-filled)
	} else {
		bar = strings.Repeat("?", width)
		pct = 0
	}

	name := filepath.Base(filename)
	fmt.Fprintf(
		os.Stdout,
		"\r[%s] (%s/%s) (%d%%) %s",
		bar,
		formatBytes(written),
		formatBytes(total),
		pct,
		name,
	)
}

func finishDownloadProgress() {
	fmt.Fprintln(os.Stdout)
}

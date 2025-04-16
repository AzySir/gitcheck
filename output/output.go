package output

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

type Output struct {
	Directory string `json:"directory"`
	Changes bool `json:"changes"`
	Branch *string `json:"branch"`
}

var (
	// Colors and font options via ANSI escape codes
	Reset     = "\033[0m"
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Gray      = "\033[37m"
	White     = "\033[97m"
	Bold      = "\033[1m"
  Italic    = "\033[3m"
	Underline = "\033[4m"
	Invert    = "\033[7m"
)

func getChange (change bool) string {
	if change {
		return Red + Bold + "Yes*" + Reset
	}
		return Green + "No" + Reset
}


func getBranch(branch *string) string {
	if *branch == "" {
		return Yellow + "Not Available" + Reset
	}

	return strings.ReplaceAll(*branch, "ref: refs/heads/", "")
}

func getDirectory(directory string) string {
	if directory == "./" {
		return "./ (current directory)"
	}
	return directory
}

func WriteOutput(data []Output) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Directory", "Active Branch", "Commit Required?"})
	for _, d := range data {
		t.AppendRow(table.Row{getDirectory(d.Directory), getBranch(d.Branch), getChange(d.Changes)})
	}
	fmt.Println(t.Render())
}
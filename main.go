package main

import (
	"fmt"
	"gitcheck/output"
	"os/exec"
	"strconv"
	"strings"
)

const LEVEL = "2"

func getGitChanges(repo string) bool {
	command := fmt.Sprintf("git -C %s status --porcelain | wc -l | xargs", repo)
	cmd := exec.Command("bash", "-c", command)
	result, _ := cmd.Output()
	trimResult := strings.TrimSuffix(string(result), "\n")
	numberOfChanges, err := strconv.Atoi(trimResult)

	if err != nil {
		fmt.Println(err.Error())
	}
	

	if numberOfChanges > 0 {
		return true
	}

	return false
}

func getBranch(dir string) *string {
	cmd := exec.Command("bash", "-c", "cat "+dir+".git/HEAD")
	result, _ := cmd.Output()

	branchName := strings.TrimSuffix(string(result), "\n")
	return &branchName
 }

 func getGitDirectories() []string {
	var gitDirectoriesNoSpaces []string
	cmd := exec.Command("/bin/bash", "-c", "find . -type d -name .git")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println("Error finding git directories:", err)
		return []string{}
	}

	resultRemoveNewline := strings.ReplaceAll(string(result), "\n", "")
	gitDirectories := strings.Split(resultRemoveNewline, ".git")
	
	// Remove empty lines
	for _, curr := range gitDirectories {
		if curr != "" {
			gitDirectoriesNoSpaces = append(gitDirectoriesNoSpaces, curr)
		}
	}

	return gitDirectoriesNoSpaces
 }

func main() {
	var gitDirectories []string
	var _output []output.Output

	gitDirectories = getGitDirectories()

	for _, dir := range gitDirectories {
		if dir != "" {
			_output = append(_output, output.Output{
				Directory: dir,
				Changes:   getGitChanges(dir),
				Branch:    getBranch(dir),
			})
		}
	}

	output.WriteOutput(_output)
}
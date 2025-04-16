package main

import (
	"reflect"
	"testing"
)

func TestGetBranch(t *testing.T) {
    want := getBranch("./test/sub_dir_with_git_commit_required/")

		if *want != "ref: refs/heads/main" {
			t.Errorf("fail test - result: %s", *want)
		}
}

func TestGetGitDirectories(t *testing.T) {
	var want []string
	var expect []string
	want = getGitDirectories()
	expect = []string{
		"./test/",
		"./test/sub_dir_with_git_no_commit/",
		"./test/sub_dir_with_git_commit_required/",
		"./test/sub_dir_with_git_no_branch/", 
		"./",
	}

	if len(want) != len(expect) {
		t.Errorf("Failed! Lenth of want vs expected does not match")
	}

	if !reflect.DeepEqual(want, expect) {
		t.Errorf("Failed!\n------------------------\n\tResult\n------------------------\n want: %v\nexpect: %v", want, expect)
	}
}

func TestGetGitChanges(t *testing.T) {
	successRepo := "./test/sub_dir_with_git_commit_required/"
	wantSuccess := getGitChanges(successRepo)
	if wantSuccess != true {
		t.Errorf("Failed! There ARE changes in %s repository", successRepo)
	}

	failRepo := "./test/sub_dir_with_git_no_commit/"
	wantFail := getGitChanges(failRepo)
	if wantFail != false {
		t.Errorf("Failed! There ARE NO changes in %s repository", failRepo)
	}
}


package git

// Merge ...
type Merge struct {
	InProgress bool
}

// GetMergeStatus ...
func GetMergeStatus() Merge {
	hasError := execGitCheckError("merge", "HEAD", "--quiet")

	return Merge{InProgress: hasError}
}

// AbortMerge ...
func AbortMerge() {
	execGit("merge", "--abort")
}

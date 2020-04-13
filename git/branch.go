package git

import (
	"bufio"
	"fmt"
	"strings"

	find "git-cleanup/utils/find"
)

// Branch ...
type Branch struct {
	Name      string
	Remote    string
	Active    bool
	Protected bool
}

// Available ...
func (b Branch) Available() bool {
	return b.Active == false && b.Protected == false
}

// Remove ...
func (b Branch) Remove() {
	execGit("branch", "-d", b.Name)
}

// Activate ...
func (b Branch) Activate() {
	execGit("checkout", b.Name)
}

// Reset ...
func (b Branch) Reset() {
	b.Activate()

	execGit("reset", "--hard", b.Remote)
}

// Status ...
func (b Branch) Status() string {
	statusString := ""
	if b.Active {
		statusString = "Active"
	}

	if b.Protected {
		statusString = "Protected"
	}

	return statusString
}

// ToString ...
func (b Branch) ToString() string {
	statusString := b.Status()

	if statusString != "" {
		statusString = fmt.Sprintf(" (%s)", statusString)
	}

	return fmt.Sprintf("%s%s", b.Name, statusString)
}

// Branches ...
type Branches []Branch

// Show ...
func (collection Branches) Show() {

	for _, branch := range collection {
		fmt.Printf("%s\n", branch.ToString())

	}
}

func (collection Branches) filter(test func(Branch) bool) (ret Branches) {
	for _, branch := range collection {
		if test(branch) {
			ret = append(ret, branch)
		}
	}
	return
}

// Available ...
func (collection Branches) Available() Branches {
	return collection.filter(func(b Branch) bool { return b.Available() })
}

// NotAvailable ...
func (collection Branches) NotAvailable() Branches {
	return collection.filter(func(b Branch) bool { return b.Available() == false })
}

// RemoveAll ...
func (collection Branches) RemoveAll() {
	for _, b := range collection {
		b.Remove()
	}
}

func parseToBranch(branchstring string) Branch {
	split := strings.Split(branchstring, " ")
	name := ""
	active := false
	protected := false
	remote := ""
	for _, s := range split {
		if s == "*" {
			active = true
		} else if name == "" {
			name = s
		} else if strings.HasPrefix(s, "[") {
			remote = strings.TrimSuffix(strings.TrimPrefix(s, "["), "]")
		}
	}
	return Branch{Name: name, Active: active, Protected: protected, Remote: remote}
}

// FindBranches ...
func FindBranches(search string) Branches {
	config := GetConfig()

	protected := find.GetMatchExpressions(config.Protected...)

	var searchExpr = find.GetMatchExpression(search)
	branches := make([]Branch, 0)
	textoutput := execGit("branch", "-vv")

	scanner := bufio.NewScanner(strings.NewReader(textoutput))
	for scanner.Scan() {
		branch := parseToBranch(scanner.Text())

		branch.Protected = protected.MatchAny(branch.Name)

		if searchExpr.MatchString(branch.Name) {
			branches = append(branches, branch)
		}
	}
	return branches
}

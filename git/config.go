package git

import (
	"bufio"
	dialog "git-cleanup/utils/dialog"
	find "git-cleanup/utils/find"
	"strings"
)

// Config ...
type Config struct {
	DefaultBranch string
	Protected     []string
	IgnoreWarning bool
}

// Protect ...
func (cfg *Config) Protect(list string) {
	split := strings.Split(list, ",")
	for _, s := range split {
		trimmed := strings.Trim(s, " ")

		if trimmed != "" && find.ContainsString(cfg.Protected, trimmed) == false {
			cfg.Protected = append(cfg.Protected, trimmed)
		}

	}
}

// RemoveProtected ...
func (cfg *Config) RemoveProtected(list string) {
	split := strings.Split(list, ",")
	for _, s := range split {
		trimmed := strings.Trim(s, " ")
		if trimmed != "" {
			list := make([]string, 0)
			for _, item := range cfg.Protected {
				if item != trimmed {
					list = append(list, item)
				}
			}

			cfg.Protected = list
		}

	}
}

// ShowProtectedPaths ...
func (cfg *Config) ShowProtectedPaths() {
	if len(cfg.Protected) > 0 {
		for _, path := range cfg.Protected {
			protectedString := ""
			if cfg.DefaultBranch == path {
				protectedString = " [Default]"
			}
			dialog.Help("%s %s", path, protectedString)
		}
	} else {
		dialog.Help("No paths protected")
	}
}

// Save ...
func (cfg *Config) Save() {
	protectText := strings.Join(cfg.Protected, ",")
	ignoreText := boolToString(cfg.IgnoreWarning)

	execGit("config", "--local", "cleanup.protected", protectText)
	execGit("config", "--local", "cleanup.ignoreWarning", ignoreText)
	if cfg.DefaultBranch != "" {
		execGit("config", "--local", "cleanup.defaultbranch", cfg.DefaultBranch)
	}
}

func parseToBool(text string) bool {
	textTrimmedLowerCase := strings.ToLower(strings.Trim(text, " "))
	if textTrimmedLowerCase == "yes" || textTrimmedLowerCase == "true" {
		return true
	}

	return false
}

func boolToString(value bool) string {
	if value {
		return "yes"
	}

	return "no"
}

// GetConfig ...
func GetConfig() Config {
	textoutput := execGit("config", "--list")
	protected := make([]string, 0)
	result := Config{Protected: protected}

	scanner := bufio.NewScanner(strings.NewReader(textoutput))
	for scanner.Scan() {
		ConfigLine := scanner.Text()

		keyValue := strings.Split(ConfigLine, "=")

		if keyValue[0] == "cleanup.protected" {
			result.Protect(keyValue[1])
		}

		if keyValue[0] == "cleanup.defaultbranch" {
			result.DefaultBranch = keyValue[1]
		}
	}

	return result
}

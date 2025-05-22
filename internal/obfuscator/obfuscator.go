package obfuscator

import "regexp"

var patterns = []struct {
	regex       *regexp.Regexp
	replacement string
}{
	{regexp.MustCompile(`(?i)password\s*=\s*[^&\s]+`), "password=***REDACTED***"},
	{regexp.MustCompile(`(?i)apikey\s*=\s*[^&\s]+`), "apiKey=***REDACTED***"},
	{regexp.MustCompile(`(?i)authorization:\s*bearer\s+[^\s]+`), "Authorization: Bearer ***REDACTED***"},
	{regexp.MustCompile(`(?i)secret\s*=\s*[^&\s]+`), "secret=***REDACTED***"},
}

func ObfuscateLine(line string) string {
	for _, p := range patterns {
		line = p.regex.ReplaceAllString(line, p.replacement)
	}
	return line
}

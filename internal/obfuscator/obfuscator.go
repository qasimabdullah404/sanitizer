package obfuscator

import "regexp"

var patterns = []struct {
	regex       *regexp.Regexp
	replacement string
}{
	{regexp.MustCompile(`(?i)password\s*=\s*[^&\s]+`), "password=***REDACTED***"},
	{regexp.MustCompile(`(?i)api[_-]?key\s*=\s*[^&\s]+`), "API_KEY=***REDACTED***"},
	{regexp.MustCompile(`(?i)authorization:\s*bearer\s+[^\s]+`), "Authorization: Bearer ***REDACTED***"},
	{regexp.MustCompile(`(?i)bearertoken:\s*[^\s]+`), "BearerToken: ***REDACTED***"},
	{regexp.MustCompile(`(?i)bearer\s+[^\s]+`), "Bearer ***REDACTED***"}, // <--- added
	{regexp.MustCompile(`(?i)secret\s*=\s*[^&\s]+`), "SECRET=***REDACTED***"},
}

func ObfuscateLine(line string) string {
	for _, p := range patterns {
		line = p.regex.ReplaceAllString(line, p.replacement)
	}
	return line
}

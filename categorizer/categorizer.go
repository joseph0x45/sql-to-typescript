package categorizer

import "strings"

func Categorize(line string) string {
	if strings.HasPrefix(line, "--") {
		return "NEXT"
	}
	if strings.HasPrefix(line, "/*") {
		return "MULTILINE_COMMENT_START"
	}
	if strings.HasSuffix(line, "*/") {
		return "MULTILINE_COMMENT_END"
	}
	if strings.HasPrefix(line, "create table") {
		return "DELIMITER_START"
	}
	if strings.EqualFold(line, ");") {
		return "DELIMITER_END"
	}
	if line == "" {
		return "NEXT"
	}
	return "FIELD"
}

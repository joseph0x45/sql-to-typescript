package purifier

import "strings"

var TOKENS = map[string][]string{
	"DELIMITER_START": {"create", "table", "if", "not", "exists", "("},
	"FIELD":           {"not null", "unique", "default", "primary key", ""},
}

func Purify(line string, category string) string {
	tokens := TOKENS[category]
	for _, token := range tokens {
		line = strings.Replace(line, token, "", -1)
	}
  return line
}

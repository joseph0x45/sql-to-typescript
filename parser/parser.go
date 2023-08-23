package parser

import (
	"errors"
	"fmt"
	"strings"
	"thewisepigeon/sql-to-ts/purifier"
)

var types = map[string]string{
  "uuid":"string",
  "text":"string",
  "char":"string",
  "varchar":"string",
  "integer":"number",
  "numeric":"number",
  "bigint":"number",
  "decimal":"number",
  "float":"number",
  "boolean":"boolean",
  "date":"Date",
  "time":"Date",
  "timestamp":"Date",
  "json":"Record<string, unknown>",
  "jsonb":"Record<string, unknown>",
  "text[]":"string[]",
  "integer[]":"number[]",
  "boolean[]":"boolean[]",
  "date[]":"Date[]",
  "json[]":"Record<string, unknown>[]",
  "jsonb[]":"Record<string, unknown>[]",
}

var registered_types = []string{}

func Parse(line string, context string, category string, parsed_tokens [][]string) (string, error) {
  if category == "NEXT"{
    return "", nil
  }
	if category == "DELIMITER_START" {
    result := strings.Split(strings.TrimSpace(purifier.Purify(line, category)), " ")
    if len(result)>1 {
      return "", errors.New("Parsing error")
    }
    table_name := result[0]
    for _, value := range registered_types{
      if table_name==value{
        return "", errors.New("Duplicate table name")
      }
    }
    registered_types = append(registered_types, table_name)
    return fmt.Sprintf("type %s = { ", table_name), nil
	}
  if category == "DELIMITER_END" {
    return "}", nil
  }
  if category == "FIELD" {
    line = strings.TrimSpace(line)
    tokens := strings.Split(line, " ")
    field_name := tokens[0]
    for _, value := range purifier.TOKENS["FIELD"]{
      if strings.ToLower(field_name) == value{
        return "", errors.New("Invalid column name")
      }
    }
    detected_type := false
    field_type := ""
    for _, value := range tokens{
      if types[value]!=""{
        if detected_type {
          return "", errors.New("Error multiple types found")
        }
        field_type = types[value]
        detected_type = true
      }
    }
    if !detected_type {
      return "", errors.New("No type detected or unsupported type")
    }
    not_null_constraint := ""
    if !strings.Contains(line, "not null") {
      not_null_constraint = "| undefined"
    }
    token := fmt.Sprintf("  %s: %s %s", field_name, field_type, not_null_constraint)
    return token, nil
  }
	return "", nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"thewisepigeon/sql-to-ts/categorizer"
	"thewisepigeon/sql-to-ts/parser"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		println("Missing file path")
		return
	}
	sql_file_path := os.Args[1]
  output_file_name := ""
  if len(os.Args) == 3 {
    output_file_name = os.Args[2]
    if !strings.HasSuffix(output_file_name, ".ts"){
      output_file_name = fmt.Sprintf("%s.ts", output_file_name)
    }
  }else {
    output_file_name = fmt.Sprintf("%v.ts",time.Now().Unix())
  }
  println(output_file_name)
	if !strings.HasSuffix(sql_file_path, ".sql") {
		println("Only SQL files are supported")
		return
	}
	reader, err := os.Open(sql_file_path)
	if err != nil {
		println("Error reading SQL file")
		return
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
  parsed_tokens := [][]string{}
	context := ""
	previous_context := ""
  current_type_index := -1
	line_number := 0
	for scanner.Scan() {
		line_number += 1
		line := strings.ToLower(scanner.Text())
		line_category := categorizer.Categorize(strings.TrimSpace(line))
		if line_category == "MULTILINE_COMMENT_START" {
			previous_context = context
			context = "MULTILINE_COMMENT"
			continue
		}
		if line_category == "MULTILINE_COMMENT_END" {
			if context != "MULTILINE_COMMENT" {
				println("Invalid character error */")
				return
			}
			context = previous_context
		}
		if line_category == "NEXT" {
			context = ""
		}
		if line_category == "DELIMITER_START" {
			if context == "PARSING" || context == "START_PARSING" {
				println("Error parsing")
				return
			}
			context = "START_PARSING"
      current_type_index+=1
      if current_type_index >= len(parsed_tokens){
        parsed_tokens = append(parsed_tokens, []string{})
      }
		}
		if line_category == "FIELD" {
			if context == "MULTILINE_COMMENT" {
				continue
			}
			if context != "START_PARSING" && context != "PARSING" {
				println("Error parsing")
				return
			}
			if context == "START_PARSING" {
				context = "PARSING"
			}
		}
		if line_category == "DELIMITER_END" {
			if context == "START_PARSING" {
				println("Error can not parse empty table")
				return
			}
			if context == "" {
				println("Error parsing")
				return
			}
		}
		token, err := parser.Parse(line, context, line_category, parsed_tokens)
		if err != nil {
			println(
				"Parsing error at line ",
				line_number,
				": ",
				"`",
				line,
				"`",
			)
			return
		}
    if token==""{
      continue
    }
    parsed_tokens[current_type_index] = append(parsed_tokens[current_type_index], token)
	}
  for _, parsed_token := range parsed_tokens{
    for _, token := range parsed_token{
      println(token)
    }
    println("")
  }
}

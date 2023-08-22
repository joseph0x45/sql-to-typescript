package main

import (
	"bufio"
	"os"
	"strings"
	"thewisepigeon/sql-to-ts/categorizer"
	"thewisepigeon/sql-to-ts/parser"
)

func main() {
	if len(os.Args) < 2 {
		println("Missing file path")
		return
	}
	sql_file_path := os.Args[1]
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
	var parsed_tokens [][]string
	var context string
	line_number := 0
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		line_category := categorizer.Categorize(strings.TrimSpace(line))
		err := parser.Parse(line, context, line_category, &parsed_tokens)
		if err != nil {
			println(
				"Parsing error at line ",
				line_number+1,
				": ",
				"`",
				line,
				"`",
			)
		}
	}
}

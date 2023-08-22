package main

import (
	"bufio"
	"os"
)

func main(){
  if len(os.Args) < 2 {
    println("Missing file path")
    return
  }
  sql_file_path := os.Args[1] 
  reader, err := os.Open(sql_file_path)
  if err!= nil {
    println("Error reading SQL file")
    return
  }
  defer reader.Close()
  scanner := bufio.NewScanner(reader)
  for scanner.Scan() {
    line := scanner.Text()
    println(line)
  }
}

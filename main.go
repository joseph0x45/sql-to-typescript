package main;

import (
  "os"
)

func main(){
  if len(os.Args) < 2 {
    println("Missing file path")
    return
  }
  sql_file_path := os.Args[1] 
}

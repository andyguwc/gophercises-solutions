/* 
Creating Command line interfaces 

Task manager 

Use BoltDB to persist the data - key value store and doesn't need MySQL installed. High read low write scenario works well 

$ cobra init --pkg-name cmd
go run main.go add

*/

package main

import (
  "fmt"
  "path/filepath"
  "os"

  "github.com/andyguwc/go-course/gophercises/7-cli-task-manager/cmd"
  "github.com/andyguwc/go-course/gophercises/7-cli-task-manager/db"
  homedir "github.com/mitchellh/go-homedir"

)

func main() {
  
  home, _ := homedir.Dir()
  dbPath := filepath.Join(home, "tasks.db")

  must(db.Init(dbPath))
  must(cmd.RootCmd.Execute())
}

func must(err error) {
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
}
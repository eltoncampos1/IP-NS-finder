package main

import (
 	"command-line/app"
	"os"
	"log"
)
func main(){
  application := app.Generate()
  if err := application.Run(os.Args); err != nil {
	log.Fatal(err)
  }
  
}

package main

import (
  
  "os"


  "github.com/akamensky/argparse"
  "github.com/chrisdaly3/pomodogo/pkg"
)


func main() {
  // Initialize a new parser object
  parser := argparse.NewParser("Start", "Creates a new pomodoro timer")
  // Create int and string flags, respectively
  timeLength := parser.Int("t", "time", &argparse.Options{Required: false, Help: "Set a time-length for the working session", Default: 25})
  goal := parser.String("g", "goal", &argparse.Options{Required: true, Help: "Set a goal for this pomodoro session", Default: ""})
  // Parse input for args
  err := parser.Parse(os.Args)
  if err != nil {
    // In case of error print error + usage
    // this is also returned in the event of the -h / --help flag
    panic(parser.Usage(err))
  }
  // Begin a new timer based on input from flags
  ticker.NewPomodoro(int(*timeLength),string(*goal)) 
}

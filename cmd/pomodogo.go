package main

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/chrisdaly3/pomodogo/conf"
	"github.com/chrisdaly3/pomodogo/pkg"
)


func main() {
  // Initialize a new parser object
  parser := argparse.NewParser("Create Timer", "Creates a new pomodoro timer")

  // Create int and string flags, respectively
  TimeLength := parser.Int("t", "time", &argparse.Options{Required: false, Help: "Set a time-length for the working session", Default: 25})
  BreakLength := parser.Int("b", "break", &argparse.Options{Required: false, Help: "Set a time-length for the break session", Default: 5})
  goal := parser.String("g", "goal", &argparse.Options{Required: true, Help: "Set a goal for this pomodoro session", Default: ""})

  // Parse input for args
  err := parser.Parse(os.Args)
  if err != nil {
    // In case of error print error + usage
    // this is also returned in the event of the -h / --help flag
    panic(parser.Usage(err))
  }
  // Begin a new timer based on input from flags
  if *TimeLength != 25 {
    config.SetPomoTime(*TimeLength)
  } else {
    config.SetPomoTime(25)
  }
  
  if *BreakLength != 5 {
    config.SetBreakTime(*BreakLength)
  } else {
    config.SetBreakTime(5)
  }
  ticker.NewPomodoro(int(*TimeLength),string(*goal)) 
}

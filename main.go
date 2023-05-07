package main

import (
  "time"
  "fmt"
  "os"
  "os/exec"

  fig "github.com/common-nighthawk/go-figure"
)

func main() {
  ticker1 := time.NewTicker(1*time.Second)
  startTime := time.Now()
  done := make(chan bool)

  go func() {
    for {
      select {
        case <-done:
          return
        case <-ticker1.C:
          cmd := exec.Command("clear");cmd.Stdout=os.Stdout; cmd.Run()
          timeElapsed := time.Since(startTime).Truncate(time.Second)
          timeFigure := fig.NewFigure(timeElapsed.String(),"",true)
          timeFigure.Print()
        }
     }
   }()

  time.Sleep(25 * time.Minute)
  ticker1.Stop()
  done <- true
  fmt.Println("Pomodoro is done")
}


package ticker

import (
  "time"
  "fmt"
  "os"
  "os/exec"

  fig "github.com/common-nighthawk/go-figure"
)
// NewPomodoro creates a new Pomodoro timer in the terminal window for time t
// and shows the session goal g.
func NewPomodoro(t int, g string) {
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
          fmt.Println(g)
        }
     }
   }()
  
  time.Sleep(time.Duration(t) * time.Minute)
  ticker1.Stop()
  done <- true
  fmt.Printf("Pomodoro is done, congrats on completing: %v", g)
}


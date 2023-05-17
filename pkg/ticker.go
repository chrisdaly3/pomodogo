package ticker

import (
  "time"
  "fmt"
  "os"
  "os/exec"

  fig "github.com/common-nighthawk/go-figure"
  "github.com/chrisdaly3/pomodogo/conf"
)

// NewPomodoro creates a new Pomodoro timer in the terminal window for time t
// and shows the session goal g.
func NewPomodoro(t int, g string) {
  ticker1 := time.NewTicker(time.Second)
  startTime := time.Now()
  done := make(chan bool)

  go func() {
    for {
      select {
        case <-done:
          return
        case <-ticker1.C:
          cmd := exec.Command("clear");cmd.Stdout=os.Stdout; cmd.Run()
          endTime := startTime.Add(time.Duration(t)*time.Minute)
          timeUntil := time.Until(endTime).Truncate(time.Second)
          timeFigure := fig.NewFigure(timeUntil.String(),"",true)
          timeFigure.Print()
          fmt.Println("\n-----\n"+g+"\n-----")
        }
     }
   }()
  
  time.Sleep(time.Duration(t) * time.Minute)
  ticker1.Stop()
  done <- true
  var response string 
  switch {
  case t==config.UserSettings.Session:
    fmt.Println("Nice work session! Ready for a short break? [y|n]\n")
    _, err := fmt.Scan(&response)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
      }
    if response=="y" {
      // TODO Adjust the way break sessions are handled. If they are parted out in the UserConfig they should be passable to a Pomo function
      NewPomodoro(config.UserSettings.Rest, "--------------\nBreaktime!!\n--------------")
    } else {
      fmt.Println("Goodbye.\n")
      os.Exit(0)
    }
  case t==config.UserSettings.Rest:
    fmt.Println("Welcome back from the break. Shall we continue? [y|n]\n")
    _, err := fmt.Scan(&response)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
      }
    if response=="y" {
      fmt.Println("What's your goal for this work session?\n")
      var nextGoal string
      _, err := fmt.Scan(&nextGoal)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      NewPomodoro(config.UserSettings.Session,nextGoal)
    } else {
      fmt.Println("Goodbye.\n")
      os.Exit(0)
    }
  }
}


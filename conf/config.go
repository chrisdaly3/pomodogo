package config

import (
    "fmt"
)

// Settings is the collection of user entered fields we store from the user's argparse entries.
type Settings struct {
  Session  int
  Rest     int
}

var UserSettings Settings

// SetPomoTime stores the configuration for the user's default working session length
func SetPomoTime(workTime int) {
  UserSettings.Session = workTime
  fmt.Printf("Session length %v stored to user configuration\n", UserSettings.Session)
}

// SetBreakTime stores the configuration for theuses default break session length
func SetBreakTime(breakTime int) {
  UserSettings.Rest = breakTime
  fmt.Printf("Break length %v stored to user configuration\n", UserSettings.Rest)
}

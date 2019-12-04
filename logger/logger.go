package logger

import (
    "fmt"
    "strings"
)

var DebugMode bool

func SetDebugMode(debug bool) {
    DebugMode = debug
}

//Display in debug mode
func D(msg ...interface{}) {
    if DebugMode {
        fmt.Println("[DEBUG]", strings.Trim(fmt.Sprint(msg), "[]"))
    }   
}

//Display info
func I(msg ...interface{}) {
    fmt.Println("[INFO]", strings.Trim(fmt.Sprint(msg), "[]"))
}

//Display error
func E(msg ...interface{}) {
    fmt.Println("[ERROR]", strings.Trim(fmt.Sprint(msg), "[]"))
}

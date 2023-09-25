package command

import "fmt"

const App = "fsync"
const AppVersion = "1.0.0"

var ErrorDone = fmt.Errorf(`ctx is done`)

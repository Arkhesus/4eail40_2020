package player

type Color struct {
	Code string
}

var Reset Color = Color{Code: "\033[0m"}
var Red Color = Color{Code: "\033[31m"}
var Green Color = Color{Code: "\033[32m"}
var Yellow Color = Color{Code: "\033[33m"}
var Blue Color = Color{Code: "\033[34m"}
var Purple Color = Color{Code: "\033[35m"}
var Cyan Color = Color{Code: "\033[36m"}
var Gray Color = Color{Code: "\033[37m"}
var White Color = Color{Code: "\033[38m"}

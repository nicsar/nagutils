/*
Utility for writing go Nagios plugins.

Version: 1.0
Author: Nicola Sarobba <nicola.sarobba@gmail.com>
Date: 2015-02-16

*/
package nagutils

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
)

var Errors = map[string]int{
	"OK":       0,
	"WARNING":  1,
	"CRITICAL": 2,
	"UNKNOWN":  3,
}

var StatusText = map[int]string{
	0: "OK",
	1: "WARNING",
	2: "CRITICAL",
	3: "UNKNOWN",
}

// outputStatus returns the exit message status string to use with NagiosExit.
func outputStatus(status int, msg string) string {
	return (fmt.Sprintf("%s: %s", StatusText[status], msg))
}

// NagiosExit print the message and then it exit with the appropriate code.
func NagiosExit(status int, msg string) {
	fmt.Println(outputStatus(status, msg))
	os.Exit(status)
}

// Basename returns the base name of a command.
func Basename(name string) string {
	var idx int = strings.LastIndex(name, "/")
	if idx == -1 {
		return name
	}
	return (name[idx+1:])
}

type ExitVal struct {
	Status   int
	Problems []string
	PerfData []string // 'label'=value[UOM];[warn];[crit];[min];[max]
}

// Problems2Str returns a string that represents the concatenation of the strings
// composing the slice separated by a ",".
func (e ExitVal) Problems2Str() string {
	out := FormatStrSlice(&e.Problems, ", ")
	return out
}

// PerfData2Str returns a string that represents the concatenation of the strings
// composing the slice separated by a " ".
func (e ExitVal) PerfData2Str() string {
	out := "|" + FormatStrSlice(&e.PerfData, " ")
	return out
}

// FormatStrSlice returns the elements of a string slice as a string.
// Each elements are separated with 'sep'.
// Example: exv.FormatStrSlice(",") -> "a, b, c".
func FormatStrSlice(sl *[]string, sep string) string {
	var buffer bytes.Buffer
	for _, word := range *sl {
		buffer.WriteString(word)
		buffer.WriteString(sep)
	}
	out := buffer.String()
	// Skip last separator.
	out = out[0 : len(out)-len(sep)]
	return out
}

// Round returns 'f' rounded.
// Examples:
// Round(10.5) -> 11
// Rounf(10.3) -> 10
// https://gist.github.com/korya
// https://gist.github.com/DavidVaini/10308388
func Round(f float64) float64 {
	return math.Floor(f + .5)
}

// RoundPlus returns 'f' rounded with precision 'places'.
// Examples:
// RoundPlus(10.619, 2) -> 10.62
// https://gist.github.com/korya
// https://gist.github.com/DavidVaini/10308388
func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

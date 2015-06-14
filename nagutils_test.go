package nagutils

import "testing"

func TestNagiosexit(t *testing.T) {
	type testvalues struct {
		status       int
		msg, correct string
	}
	var cases = []testvalues{
		{0, "Service up", "OK: Service up"},
		{1, "Service problem", "WARNING: Service problem"},
		{2, "Service down", "CRITICAL: Service down"},
		{3, "Service undefined", "UNKNOWN: Service undefined"},
	}
	for _, c := range cases {
		got := outputStatus(c.status, c.msg)
		if got != c.correct {
			t.Errorf("NagiosExit(%q, %q) == %q, correct %q", c.status, c.msg, got, c.correct)
		}
	}
}

func TestBasename(t *testing.T) {
	type testvalues struct {
		commandName, correct string
	}
	var cases = []testvalues{
		{"/usr/lib/nagios/plugins/check_command", "check_command"},
		{"/usr/local/nagios/libexec/check_command", "check_command"},
		{"/lib/plugins/check_command", "check_command"},
		{"./check_command", "check_command"},
		{"./bin/check_command", "check_command"},
		{"/bin/check_command", "check_command"},
		{"check_command", "check_command"},
		{"/usr//bin/check_command", "check_command"},
		{"/usr/bin//check_command", "check_command"},
	}
	for _, c := range cases {
		got := Basename(c.commandName)
		if got != c.correct {
			t.Errorf("Basename(%q) == %q, correct %q", c.commandName, got, c.correct)
		}
	}
}

func TestFormatStrSlice(t *testing.T) {
	type testvalues struct {
		in      []string
		sep     string
		correct string
	}
	var cases = []testvalues{
		{[]string{"One", "Two", "Three", "Four"}, ",", "One,Two,Three,Four"},
		{[]string{"One", "Two", "Three", "Four"}, " ", "One Two Three Four"},
		{[]string{"One", "Two", "Three", "Four"}, ", ", "One, Two, Three, Four"},
	}
	for _, c := range cases {
		got := FormatStrSlice(&c.in, c.sep)
		if got != c.correct {
			t.Errorf("FormatStrSlice(%q) == %q, correct %q", c.in, got, c.correct)
		}
	}
}

func TestProblems2Str(t *testing.T) {
	var e ExitVal
	e.Problems = []string{"[Volume 1] above threshold", "[Volume 2] above threshold"}
	type testvalues struct {
		in      ExitVal
		correct string
	}
	var cases = []testvalues{
		{e, "[Volume 1] above threshold, [Volume 2] above threshold"},
	}
	for _, c := range cases {
		got := c.in.Problems2Str()
		if got != c.correct {
			t.Errorf("Problems2Str() == %q, correct %q", got, c.correct)
		}
	}
}
func TestPerfData2Str(t *testing.T) {
	var e ExitVal
	e.PerfData = []string{"'[Volume Volume-1, Pool 1]'=2.19TB;9;10;0;10.77", "'[Volume Volume-2, Pool 2]'=1.71TB;4;5;0;5.35"}
	type testvalues struct {
		in      ExitVal
		correct string
	}
	var cases = []testvalues{
		{e, "|'[Volume Volume-1, Pool 1]'=2.19TB;9;10;0;10.77 '[Volume Volume-2, Pool 2]'=1.71TB;4;5;0;5.35"},
	}
	for _, c := range cases {
		got := c.in.PerfData2Str()
		if got != c.correct {
			t.Errorf("PerfData2Str() == %q, correct %q", got, c.correct)
		}
	}
}
func TestRound(t *testing.T) {
	type testvalues struct {
		f       float64
		correct float64
	}
	var cases = []testvalues{
		{123.499, 123},
		{123.5, 124},
		{123.999, 124},
	}
	for _, c := range cases {
		got := Round(c.f)
		if got != c.correct {
			t.Errorf("Round(%q) == %q, correct %q", c.f, got, c.correct)
		}
	}
}

func TestRoundPlus(t *testing.T) {
	type testvalues struct {
		f         float64
		precision int
		correct   float64
	}
	var cases = []testvalues{
		{123.554999, 3, 123.555},
		{123.555555, 3, 123.556},
		{123.558, 2, 123.56},
	}
	for _, c := range cases {
		got := RoundPlus(c.f, c.precision)
		if got != c.correct {
			t.Errorf("RoundPlus(%q, %q) == %q, correct %q", c.f, c.precision, got, c.correct)
		}
	}
}

package cronparser

import (
	"reflect"
	"testing"
)

func testTimeAndInterval(t *testing.T, section *CronSection, err error, expectedTime, expectedInterval string) {
	if err != nil {
		t.Fatal(err)
	}

	if section.Time != expectedTime {
		t.Fatalf("Time was not parsed correctly: %q", section.Time)
	}

	if section.Interval != expectedInterval {
		t.Fatalf("Interval was not parsed correctly, should be empty and is %q", section.Interval)
	}
}

var parseableLines = map[string]*CronEntry{
	"17 *    * * *   root    cd / && run-parts --report /etc/cron.hourly": &CronEntry{
		Minute:    &CronSection{Time: "17"},
		Hour:      &CronSection{Time: "*"},
		Day:       &CronSection{Time: "*"},
		Month:     &CronSection{Time: "*"},
		DayOfWeek: &CronSection{Time: "*"},
		User:      "root",
		Command:   "cd / && run-parts --report /etc/cron.hourly",
	},
	"25 6    * * *   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.daily )": &CronEntry{
		Minute:    &CronSection{Time: "25"},
		Hour:      &CronSection{Time: "6"},
		Day:       &CronSection{Time: "*"},
		Month:     &CronSection{Time: "*"},
		DayOfWeek: &CronSection{Time: "*"},
		User:      "root",
		Command:   "test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.daily )",
	},
	"47 6    * * 7   fart test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.weekly )": &CronEntry{
		Minute:    &CronSection{Time: "47"},
		Hour:      &CronSection{Time: "6"},
		Day:       &CronSection{Time: "*"},
		Month:     &CronSection{Time: "*"},
		DayOfWeek: &CronSection{Time: "7"},
		User:      "fart",
		Command:   "test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.weekly )",
	},
}

func TestParseLine(t *testing.T) {
	for line, cronentry := range parseableLines {
		tstCe, err := parseLine(line)
		if err != nil {
			t.Fatal("Could not parse line:", line)
		}

		if !reflect.DeepEqual(cronentry, tstCe) {
			t.Fatal("Cron entries are not equal for line:", line)
		}
	}

	lines := []string{
		"#",
		"2",
		"*/* * * * * root /usr/local/rtm/bin/rtm 9 > /dev/null 2> /dev/null",
		"25/* 6    * * *   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.daily )",
	}

	for _, line := range lines {
		_, err := parseLine(line)
		if err == nil {
			t.Fatal("Parsing succeeded for bad line:", line)
		}
	}
}

func TestParseSection(t *testing.T) {
	section, err := parseSection("1")
	testTimeAndInterval(t, section, err, "1", "")

	section, err = parseSection("*/1")
	testTimeAndInterval(t, section, err, "*", "1")

	section, err = parseSection("*-1")
	if err == nil {
		t.Fatal("did not error on invalid input")
	}

	section, err = parseSection("d")
	if err == nil {
		t.Fatal("did not error on invalid input")
	}

	section, err = parseSection("1/*")
	if err == nil {
		t.Fatal("did not error on invalid input")
	}
}

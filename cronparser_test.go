package cronparser

import "testing"

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

func TestParseLine(t *testing.T) {
	lines := []string{
		"17 *    * * *   root    cd / && run-parts --report /etc/cron.hourly",
		"25 6    * * *   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.daily )",
		"47 6    * * 7   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.weekly )",
		"52 6    1 * *   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.monthly )",
		"*/1 * * * * root /usr/local/rtm/bin/rtm 9 > /dev/null 2> /dev/null",
	}

	for _, line := range lines {
		_, err := parseLine(line)
		if err != nil {
			t.Fatal("Could not parse line:", line)
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

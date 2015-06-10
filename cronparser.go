package cronparser

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	regexSection    = regexp.MustCompile(`^[0-9*]+$`)
	regexRHS        = regexp.MustCompile(`^\d*$`)
	regexWhitespace = regexp.MustCompile(`[ \t]+`)
)

type CronSection struct {
	Time     string
	Interval string
}

type CronEntry struct {
	Minute    *CronSection
	Hour      *CronSection
	Day       *CronSection
	Month     *CronSection
	DayOfWeek *CronSection
	User      string
	Command   string
}

func parseSectionVar(cs **CronSection, str string) error {
	var err error
	*cs, err = parseSection(str)
	if err != nil {
		return fmt.Errorf("Minute is invalid: %q, error: %v", str, err)
	}
	return nil
}

func parseLine(line string) (*CronEntry, error) {
	strs := regexWhitespace.Split(line, 7)
	entry := &CronEntry{User: strs[5], Command: strs[6]}

	if err := parseSectionVar(&entry.Minute, strs[0]); err != nil {
		return nil, err
	}

	if err := parseSectionVar(&entry.Hour, strs[1]); err != nil {
		return nil, err
	}

	if err := parseSectionVar(&entry.Day, strs[2]); err != nil {
		return nil, err
	}

	if err := parseSectionVar(&entry.Month, strs[3]); err != nil {
		return nil, err
	}

	if err := parseSectionVar(&entry.DayOfWeek, strs[4]); err != nil {
		return nil, err
	}

	return entry, nil
}

func parseSection(section string) (*CronSection, error) {
	sections := strings.SplitN(section, "/", 2)

	strs := []string{}

	for _, sect := range sections {
		if !regexSection.MatchString(sect) {
			return nil, fmt.Errorf("Could not parse section part %q", sect)
		}

		strs = append(strs, sect)
	}

	switch len(strs) {
	case 2:
		break
	case 1:
		strs = append(strs, "")
	default:
		return nil, fmt.Errorf("Could not parser cron section %q", section)
	}

	if !regexRHS.MatchString(strs[1]) {
		return nil, fmt.Errorf("Right-hand side may not have a starred interval")
	}

	return &CronSection{Time: strs[0], Interval: strs[1]}, nil
}

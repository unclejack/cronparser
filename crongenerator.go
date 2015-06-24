package cronparser

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// TimeSectionMap is a map of Time value to Cron Section. This is useful for
// associating the two values and is used in (*CronEntry).Times().
type TimeSectionMap struct {
	Time    int
	Section *CronSection
}

// String returns the entry as a string, in crontab format.
func (ce *CronEntry) String() string {
	items := []string{
		ce.Minute.String(),
		ce.Hour.String(),
		ce.Day.String(),
		ce.Month.String(),
		ce.DayOfWeek.String(),
		ce.User,
		ce.Command,
	}
	return strings.Join(items, " ")
}

// Times returns all time-based sections in order.
func (ce *CronEntry) Times(theTime time.Time) []TimeSectionMap {
	return []TimeSectionMap{
		TimeSectionMap{
			Time:    theTime.Minute(),
			Section: ce.Minute,
		},
		TimeSectionMap{
			Time:    theTime.Hour(),
			Section: ce.Hour,
		},
		TimeSectionMap{
			Time:    theTime.Day(),
			Section: ce.Day,
		},
		TimeSectionMap{
			Time:    int(theTime.Month()),
			Section: ce.Month,
		},
		TimeSectionMap{
			Time:    int(theTime.Weekday()),
			Section: ce.DayOfWeek,
		},
	}
}

// String returns the CronSection entry as a string, such as "1" or "*/12"
func (cs *CronSection) String() string {
	if cs.Interval != "" {
		return fmt.Sprintf("%s/%s", cs.Time, cs.Interval)
	}

	return fmt.Sprintf("%s", cs.Time)
}

// String returns a fully compatible crontab generated from the CronParser
// type. Environment variables are stacked at the top of the crontab before any
// entries.
func (cp *CronParser) String() string {
	retval := ""

	keys := []string{}

	for key := range cp.Environment {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		retval += fmt.Sprintf("%s=%q\n", key, cp.Environment[key])
	}

	for _, entry := range cp.CronTab {
		retval += entry.String() + "\n"
	}

	return retval
}

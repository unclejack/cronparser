package cronparser

import (
	"fmt"
	"sort"
	"strings"
)

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

func (cs *CronSection) String() string {
	if cs.Interval != "" {
		return fmt.Sprintf("%s/%s", cs.Time, cs.Interval)
	}

	return fmt.Sprintf("%s", cs.Time)
}

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

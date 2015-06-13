## cronparser - parse crontabs into data structures

 Invoke cronparser like so:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	cp := NewCronParser()
	b, err := ioutil.ReadFile("/etc/crontab")
	if err != nil {
		os.Exit(1)
	}

	if err := cp.ParseCronTab(string(b)); err != nil {
		os.Exit(1)
	}

  // Dump the crontab

	for key, value := range cp.Environment {
		fmt.Println(key, value)
	}

	for _, entry := range cp.CronTab {
		fmt.Println("Minute:", entry.Minute.Time, "/", entry.Minute.Interval)
		// do the rest at your leisure.
	}

	// or generate a real one
	
	fmt.Println(cp.String())
}
```


Most of what you need to know is in the
[godoc](http://godoc.org/github.com/erikh/cronparser). Please refer there for a
full description of how to use this library.

### Author

Erik Hollensbe <erik+github@hollensbe.org>

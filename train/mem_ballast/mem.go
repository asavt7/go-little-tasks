package main

import (
	"fmt"
	"math"
	"time"
)

// https://blog.twitch.tv/en/2019/04/10/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap/
// run `go run ./mem.go`  and see htop
// or ```ps -eo pmem,comm,pid,maj_flt,min_flt,rss,vsz --sort -rss | numfmt --header --to=iec --field 4-5 | numfmt --header --from-unit=1024 --to=iec --field 6-7 | column -t | egrep "mem"```
func main() {
	_ = make([]byte, 10<<30) //10 Gi
	fmt.Println("started")
	<-time.After(time.Duration(math.MaxInt64))
}

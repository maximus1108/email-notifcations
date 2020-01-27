package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Format("02-01-2006"))
}

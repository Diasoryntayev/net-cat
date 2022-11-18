package pkg

import (
	"fmt"
	"time"
)

func TimeValue() string {
	time := time.Now()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
}

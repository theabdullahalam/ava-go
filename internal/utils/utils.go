package utils

import (
	"fmt"
	"time"
)

func GetTimeStampString() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
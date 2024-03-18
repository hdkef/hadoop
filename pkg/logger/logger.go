package logger

import (
	"fmt"
	"runtime"
)

func LogError(err error) {

	_, file, line, _ := runtime.Caller(1)

	fmt.Printf("level=err, file=%s, line=%d, %s\n", file, line, err.Error())

}

package util

import (
	"fmt"
	"os"

	"github.com/wudaoluo/golog"
)

func Exit(code int, msg string) {
	golog.Error(msg, "code", code)
	fmt.Println(msg)
	os.Exit(code)
}

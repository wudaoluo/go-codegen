package util

import (
	"fmt"
	"github.com/wudaoluo/golog"
	"os"
)

func Exit(code int,msg string) {
	golog.Error(msg,"code",code)
	fmt.Println(msg)
	os.Exit(code)
}

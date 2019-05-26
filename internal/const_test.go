package internal

import (
	"fmt"
	"testing"
)

func Test_Gen(t *testing.T) {
	var a = GEN_MYSQL_DOC
	fmt.Println(a)
	if a.String() != "mysql add doc" {
		t.Failed()
	}

}

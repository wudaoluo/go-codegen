package generate

import (
	"strings"
)

func (g *Generate)WithNotFirstTitle(args string) string {
	if !strings.Contains(args,"_") {
		return args
	}

	list := strings.Split(args,"_")
	var a string
	for i,word := range list {
		if i == 0 {
			a = a+ word
			continue
		}
		a = a+ strings.Title(word)
	}

	return a
}

func (g *Generate)WithTitle(args string) string {
	if !strings.Contains(args,"_") {
		return strings.Title(args)
	}

	list := strings.Split(args,"_")
	var a string
	for _,word := range list {
		a = a+ strings.Title(word)
	}

	return a
}

func (g *Generate)WithComment(args string) string {
	if args == "" {
		return args
	}

	return "// " + args

}





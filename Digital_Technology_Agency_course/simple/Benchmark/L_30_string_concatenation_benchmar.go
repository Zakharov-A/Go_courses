package main

import (
	"bytes"
	"strings"
)

// ---- String concatenation Union----

// var params = []string{"str-1", "str-2", "str-3", "str-4"}

// func main() {
// 	println(conPlus(params))
// }

// func conPlus(args []string) string {
// 	result := ""
// 	for _, arg := range args {
// 		result += arg
// 	}
// 	return result
// }

// ----

// ---- Bytes ----

// var params = []string{"str-1", "str-2", "str-3", "str-4"}

// func main() {
// 	println(conBytes(params))
// }

// func conPlus(args []string) string {
// 	result := ""
// 	for _, arg := range args {
// 		result += arg
// 	}
// 	return result
// }

// func conBytes(args []string) string {
// 	buffer := bytes.Buffer{}
// 	for _, arg := range args {
// 		buffer.WriteString(arg)
// 	}
// 	return buffer.String()

// }

// ----

// ---- Package String ----

var params = []string{"str-1", "str-2", "str-3", "str-4"}

func main() {
	println(conCopy(params))
}

func conCopy(args []string) string {
	lenb := make([]byte, len(args)*len(args[0]))
	offset := 0
	for _, arg := range args {
		offset += copy(lenb[offset:], arg)
	}
	return string(lenb[:])

}

func conPlus(args []string) string {
	result := ""
	for _, arg := range args {
		result += arg
	}
	return result
}

func conBytes(args []string) string {
	buffer := bytes.Buffer{}
	for _, arg := range args {
		buffer.WriteString(arg)
	}
	return buffer.String()

}

func conStrings(args []string) string {
	builder := strings.Builder{}
	for _, arg := range args {
		builder.WriteString(arg)
	}
	return builder.String()
}

// ----

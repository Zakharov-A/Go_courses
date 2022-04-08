package main

// ---- File creation ----

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Printf("Error creating file [%s]", err.Error())
// 		return
// 	}
// 	println(file.Name())
// }

// ----

// ---- Writing information to a file ----

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	text := "Test words"
// 	file, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Printf("Error creating file [%s]", err.Error())
// 		return
// 	}
// 	writeString, err := file.WriteString(text)
// 	if err != nil {
// 		fmt.Printf("File write error [%s]", err.Error())
// 		return
// 	}
// 	println(writeString)
// }

// ----

// ---- Reading information from a file ----

import (
	"fmt"
	"io"
	"os"
)

func main() {
	text := "Test code in file"
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("Error creating file [%s]", err.Error())
		return
	}
	stringLen, err := file.WriteString(text)
	if err != nil {
		fmt.Printf("File write error [%s]", err.Error())
		return
	}
	openFile, err := os.Open(file.Name())
	if err != nil {
		fmt.Printf("File read error [%s]", err.Error())
		return
	}
	data := make([]byte, stringLen)
	for {
		_, err := openFile.Read(data)
		if err != io.EOF {
			break

		}
		if err != nil {
			fmt.Printf("File read error [%s]", err.Error())
			return
			break
		}
	}
	fmt.Printf(string(data))

}

// ----

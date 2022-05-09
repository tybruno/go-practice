package assignment_3

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("myfile.txt")
	bytes, err := io.Copy(os.Stdout, file)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bytes)
}

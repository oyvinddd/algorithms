package graphs

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ReadFromFile ...
func ReadFromFile(filename string) {

	path := strings.Join([]string{"/datastructures/graphs/data", filename}, "/")
	pwd, _ := os.Getwd()

	file, err := os.Open(pwd + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		// TODO: add s,t to graph
		if err != nil {
			break
		}
	}
}

// PrintGraph ...
func PrintGraph() {}

package graphs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
)

// ReadFromFile ...
func ReadFromFile(graph *graphs.Graph, filename string) {

	path := strings.Join([]string{"/datastructures/graphs/data", filename}, "/")
	pwd, _ := os.Getwd()

	file, err := os.Open(pwd + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string

	// 1. Read number of edges (this is the number on line one)
	line, err = reader.ReadString('\n')

	fmt.Printf("E: %v\n", line)

	// 2. Read number of vertices (this is the number on line two)
	line, err = reader.ReadString('\n')

	fmt.Printf("V: %v\n", line)

	for {
		// Read edges from file one by one
		line, err = reader.ReadString('\n')
		fmt.Printf("EDGE: %v\n", line)

		if err != nil {
			// EOF. Gtfo!
			break
		}
	}
}

// PrintGraph prints graph properties to the console
// TODO: make this accept both graphs and digraphs
func PrintGraph(g *graphs.Graph) {
	fmt.Println("#### PRINTING GRAPH... ####")
	fmt.Printf("No. of e: %v\nNo. of v: %v\n", g.NoOfE(), g.NoOfV())
	for v := 0; v < g.NoOfV(); v++ {
		for w := range g.GetAdj(v) {
			fmt.Printf("(%v, %v)\n", v, w.(int))
		}
	}
	fmt.Println("###########################")
}

// Package galeshapeley implements the Gale-Shapeley algorithm
package galeshapeley

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/queue"
)

// GaleShapeley runs the Gale-Shapeley algorithm
func GaleShapeley(men []Person, women []Person) {

	mPrefs := make(map[string]*queue.Queue)
	wPrefs := make(map[string]map[string]int)
	wEngag := make(map[string]*Person)
	freeMen := queue.NewQueue()

	for _, man := range men {

		// Create a queue of all possible women for a man
		mPrefs[man.Name] = queue.NewQueue()
		for _, wmn := range man.Pref {
			q := mPrefs[man.Name]
			q.Enqueue(wmn)
		}

		// Add man to list of free men
		freeMen.Enqueue(man)
	}

	// Create preference dictionary on how women rank men (O(n^2))
	for _, wmn := range women {
		wPrefs[wmn.Name] = make(map[string]int)
		for idx, man := range wmn.Pref {
			wPrefs[wmn.Name][man.Name] = idx
		}
	}

	// Run algorithm until there are no free men
	for !freeMen.IsEmpty() {

		m := (*freeMen.Dequeue()).(Person)
		w := (*mPrefs[m.Name].Dequeue()).(Person)

		// If woman is free (never been proposed to)
		if wEngag[w.Name] == nil {
			wEngag[w.Name] = &m
		} else {
			m2 := wEngag[w.Name] // m'
			// if w prefers m to m'
			if wPrefs[w.Name][m.Name] < wPrefs[w.Name][m2.Name] {
				freeMen.Enqueue(*m2)
				wEngag[w.Name] = &m
			} else {
				freeMen.Enqueue(m)
				wEngag[w.Name] = m2
			}
		}
	}

	// Print matching to the console
	// TODO: print if there is an instability, if the matching is stable and/or perfect
	printMatching(women, wEngag)
}

func printMatching(women []Person, e map[string]*Person) {
	for _, woman := range women {
		man := e[woman.Name]
		fmt.Printf("WOMAN %v is married to MAN %v\n", woman.Name, man.Name)
	}
}

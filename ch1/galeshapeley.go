package ch1

import "fmt"

// Person ...
type Person struct {
	Name      string
	EngagedTo *Person
	Priority  []*Person
}

// GS - Gale-Shapeley algorithm
// TOOD: need to test if the algorithm is actually correct
// It's wrong. We cannot assume the first element in Priority list of w is part of 'prefers'
func GS(men []Person, women []Person) {
	engagedCount, counter := 0, 0
	for engagedCount < len(men) {
		m := &men[counter]
		if m.isFree() && len(m.Priority) > 0 {
			w := m.pop()
			fmt.Printf("Popping %s for %s!\n", w.Name, m.Name)
			// w is free, get engaged
			if w.isFree() {
				engagement(m, w)
				engagedCount++
				fmt.Printf("#1 %s is engaged to %s (%d)\n", m.Name, w.Name, len(m.Priority))
			} else {
				if w.prefers(m) {
					engagement(m, w)
					fmt.Printf("#2 %s is engaged to (%s)\n", m.Name, w.Name)
				}
			}
		}
		counter++
		// reset counter if it exceeds number of people
		if counter > len(men)-1 {
			counter = 0
		}
	}
	printPairs(men)
}

func (p *Person) pop() *Person {
	if len(p.Priority) == 0 {
		return nil
	}
	person := p.Priority[0]
	p.Priority = p.Priority[1:]
	return person
}

func (p *Person) isFree() bool {
	return p.EngagedTo == nil
}

func (p *Person) prefers(p2 *Person) bool {
	if len(p.Priority) == 0 {
		return false
	}
	return p.Priority[0] == p2
}

func engagement(m *Person, w *Person) {
	oldBf := w.EngagedTo
	if oldBf != nil {
		oldBf.EngagedTo = nil
	}
	m.EngagedTo = w
	w.EngagedTo = m
}

func printPairs(people []Person) {
	for _, person := range people {
		fmt.Printf("%s is married to %s\n", person.Name, person.EngagedTo.Name)
	}
}

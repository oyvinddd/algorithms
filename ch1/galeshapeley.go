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
func GS(men []Person, women []Person) {
	engagedCount, counter := 0, 0
	for engagedCount < len(men) {
		m := &men[counter]
		if m.isFree() && len(m.Priority) > 0 {
			w := m.pop()
			// w is free, get engaged
			if w.isFree() {
				engagement(m, w)
				engagedCount++
				fmt.Printf("%s is engaged to %s\n", m.Name, w.Name)
			} else {
				if w.prefers(m) {
					engagement(m, w)
					fmt.Printf("%s is engaged to %s\n", m.Name, w.Name)
				}
			}
		}
		counter++
	}
	printPairs(men)
}

func (p Person) pop() *Person {
	if len(p.Priority) == 0 {
		return nil
	}
	person := p.Priority[0]
	p.Priority = p.Priority[1:]
	return person
}

func (p Person) isFree() bool {
	return p.EngagedTo == nil
}

func (p Person) prefers(p2 *Person) bool {
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

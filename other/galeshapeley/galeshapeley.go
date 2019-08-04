package galeshapeley

import "fmt"

// Person ...
type Person struct {
	Name      string
	EngagedTo *Person
	Priority  []*Person
}

// GS - Gale-Shapeley algorithm
func GS(men []Person, women []Person) {
	engagedCount, counter := 0, 0
	for engagedCount < len(men) {
		m := &men[counter]
		if m.isFree() && len(m.Priority) > 0 {
			w := m.pop()
			// w is free, get engaged directly
			if w.isFree() {
				engagement(m, w)
				engagedCount++
			} else {
				if w.prefers(m) {
					engagement(m, w)
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
	ci, ni := 0, 0
	for i, person := range p.Priority {
		// we're comparing strings instead of actual memory addresses since priority
		// list of men on the women has different addresses than the actual men (they were copied at some point, I guess)
		// ALSO, this could be optimised by having an NxN matrix with the womens preferences instead of a list on each woman
		if p.EngagedTo.Name == person.Name {
			ci = i
		}
		if p2.Name == person.Name {
			ni = i
		}
	}
	return ni < ci
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
	fmt.Println("########## RESULTS ##########")
	for _, person := range people {
		fmt.Printf("%s is married to %s\n", person.Name, person.EngagedTo.Name)
	}
	fmt.Println("#############################")
}

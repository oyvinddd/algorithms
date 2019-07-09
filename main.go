package main

import "bitbucket.com/oyvind_hauge/inf234/ch1"

func main() {

	m1 := &ch1.Person{
		Name:      "m1",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	m2 := &ch1.Person{
		Name:      "m2",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	m3 := &ch1.Person{
		Name:      "m3",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	w1 := &ch1.Person{
		Name:      "w1",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	w2 := &ch1.Person{
		Name:      "w2",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	w3 := &ch1.Person{
		Name:      "w3",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}

	// populate priority lists for all people
	m1.Priority = []*ch1.Person{w1, w2, w3}
	m2.Priority = []*ch1.Person{w3, w2, w1}
	m3.Priority = []*ch1.Person{w3, w1, w2}
	w1.Priority = []*ch1.Person{m1, m2, m3}
	w2.Priority = []*ch1.Person{m3, m2, m1}
	w3.Priority = []*ch1.Person{m2, m3, m1}

	men := []ch1.Person{*m1, *m2, *m3}
	women := []ch1.Person{*w1, *w2, *w3}

	ch1.GS(men, women)
}

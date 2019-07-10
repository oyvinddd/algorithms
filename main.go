package main

import "bitbucket.com/oyvind_hauge/inf234/ch1"

func main() {

	ryan := &ch1.Person{
		Name:      "Ryan",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	josh := &ch1.Person{
		Name:      "Josh",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	blake := &ch1.Person{
		Name:      "Blake",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	connor := &ch1.Person{
		Name:      "Connor",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	lizzy := &ch1.Person{
		Name:      "Lizzy",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	sarah := &ch1.Person{
		Name:      "Sarah",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	zoey := &ch1.Person{
		Name:      "Zoey",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}
	daniela := &ch1.Person{
		Name:      "Daniela",
		EngagedTo: nil,
		Priority:  []*ch1.Person{},
	}

	// populate priority lists for all people
	ryan.Priority = []*ch1.Person{lizzy, sarah, zoey, daniela}
	josh.Priority = []*ch1.Person{sarah, lizzy, daniela, zoey}
	blake.Priority = []*ch1.Person{sarah, daniela, zoey, lizzy}
	connor.Priority = []*ch1.Person{lizzy, sarah, zoey, daniela}

	lizzy.Priority = []*ch1.Person{ryan, blake, josh, connor}
	sarah.Priority = []*ch1.Person{ryan, blake, connor, josh}
	zoey.Priority = []*ch1.Person{connor, josh, ryan, blake}
	daniela.Priority = []*ch1.Person{ryan, josh, connor, blake}

	men := []ch1.Person{*ryan, *josh, *blake, *connor}
	women := []ch1.Person{*lizzy, *sarah, *zoey, *daniela}

	ch1.GS(men, women)
}

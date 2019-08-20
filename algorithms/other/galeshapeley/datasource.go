package galeshapeley

// Datasource for the GS algorithm

// Person struct
type Person struct {
	Name string
	Pref []Person
}

var m1 = Person{Name: "A"}
var m2 = Person{Name: "B"}
var m3 = Person{Name: "C"}
var w1 = Person{Name: "X"}
var w2 = Person{Name: "Y"}
var w3 = Person{Name: "Z"}

// GetMen get all men
func GetMen() []Person {
	m1.Pref = []Person{w1, w2, w3}
	m2.Pref = []Person{w1, w3, w2}
	m3.Pref = []Person{w3, w2, w1}
	return []Person{m1, m2, m3}
}

// GetWomen get all women
func GetWomen() []Person {
	w1.Pref = []Person{m3, m2, m1}
	w2.Pref = []Person{m2, m1, m3}
	w3.Pref = []Person{m3, m2, m1}
	return []Person{w1, w2, w3}
}

// GetPeople get two equal lists of men and women
func GetPeople() (m []Person, w []Person) {
	return GetMen(), GetWomen()
}

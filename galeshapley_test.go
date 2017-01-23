package match

import "testing"

func TestPartners(t *testing.T) {
	abe := Person{name: "Abe"}
	bob := Person{name: "Bob"}
	col := Person{name: "Col"}
	dan := Person{name: "Dan"}
	ed := Person{name: "Ed"}
	fred := Person{name: "Fred"}
	gav := Person{name: "Gav"}
	hal := Person{name: "Hal"}
	ian := Person{name: "Ian"}
	jon := Person{name: "Jon"}
	abi := Person{name: "Abi"}
	bea := Person{name: "Bea"}
	cath := Person{name: "Cath"}
	dee := Person{name: "Dee"}
	eve := Person{name: "Eve"}
	fay := Person{name: "Fay"}
	gay := Person{name: "Gay"}
	hope := Person{name: "Hope"}
	ivy := Person{name: "Ivy"}
	jan := Person{name: "Jan"}

	abe.candidates = []string{"Abi", "Eve", "Cath", "Ivy", "Jan", "Dee", "Fay", "Bea", "Hope", "Gay"}
	bob.candidates = []string{"Cath", "Hope", "Abi", "Dee", "Eve", "Fay", "Bea", "Jan", "Ivy", "Gay"}
	col.candidates = []string{"Hope", "Eve", "Abi", "Dee", "Bea", "Fay", "Ivy", "Gay", "Cath", "Jan"}
	dan.candidates = []string{"Ivy", "Fay", "Dee", "Gay", "Hope", "Eve", "Jan", "Bea", "Cath", "Abi"}
	ed.candidates = []string{"Jan", "Dee", "Bea", "Cath", "Fay", "Eve", "Abi", "Ivy", "Hope", "Gay"}
	fred.candidates = []string{"Bea", "Abi", "Dee", "Gay", "Eve", "Ivy", "Cath", "Jan", "Hope", "Fay"}
	gav.candidates = []string{"Gay", "Eve", "Ivy", "Bea", "Cath", "Abi", "Dee", "Hope", "Jan", "Fay"}
	hal.candidates = []string{"Abi", "Eve", "Hope", "Fay", "Ivy", "Cath", "Jan", "Bea", "Gay", "Dee"}
	ian.candidates = []string{"Hope", "Cath", "Dee", "Gay", "Bea", "Abi", "Fay", "Ivy", "Jan", "Eve"}
	jon.candidates = []string{"Abi", "Fay", "Jan", "Gay", "Eve", "Bea", "Dee", "Cath", "Ivy", "Hope"}
	abi.candidates = []string{"Bob", "Fred", "Jon", "Gav", "Ian", "Abe", "Dan", "Ed", "Col", "Hal"}
	bea.candidates = []string{"Bob", "Abe", "Col", "Fred", "Gav", "Dan", "Ian", "Ed", "Jon", "Hal"}
	cath.candidates = []string{"Fred", "Bob", "Ed", "Gav", "Hal", "Col", "Ian", "Abe", "Dan", "Jon"}
	dee.candidates = []string{"Fred", "Jon", "Col", "Abe", "Ian", "Hal", "Gav", "Dan", "Bob", "Ed"}
	eve.candidates = []string{"Jon", "Hal", "Fred", "Dan", "Abe", "Gav", "Col", "Ed", "Ian", "Bob"}
	fay.candidates = []string{"Bob", "Abe", "Ed", "Ian", "Jon", "Dan", "Fred", "Gav", "Col", "Hal"}
	gay.candidates = []string{"Jon", "Gav", "Hal", "Fred", "Bob", "Abe", "Col", "Ed", "Dan", "Ian"}
	hope.candidates = []string{"Gav", "Jon", "Bob", "Abe", "Ian", "Dan", "Hal", "Ed", "Col", "Fred"}
	ivy.candidates = []string{"Ian", "Col", "Hal", "Gav", "Fred", "Bob", "Abe", "Ed", "Jon", "Dan"}
	jan.candidates = []string{"Ed", "Hal", "Gav", "Abe", "Bob", "Jon", "Col", "Ian", "Fred", "Dan"}

	proposers := []Person{abe, bob, col, dan, ed, fred, gav, hal, ian, jon}
	candidates := []Person{abi, bea, cath, dee, eve, fay, gay, hope, ivy, jan}

	correct := map[string]string{}
	correct["Abe"] = "Ivy"
	correct["Bob"] = "Cath"
	correct["Col"] = "Dee"
	correct["Dan"] = "Fay"
	correct["Ed"] = "Jan"
	correct["Fred"] = "Bea"
	correct["Gav"] = "Gay"
	correct["Hal"] = "Eve"
	correct["Ian"] = "Hope"
	correct["Jon"] = "Abi"

	matches := GaleShapley(proposers, candidates)
	for _, pair := range matches {
		if correct[pair.name] != pair.partner.name {
			t.Fatalf("Incorrect pairing, expected %s to pair with %s. received: %s", pair.name, correct[pair.name], pair.partner.name)
		}
	}
}

func TestSimple(t *testing.T) {
	alex := Person{name: "Alex"}
	bradley := Person{name: "Bradley"}
	christina := Person{name: "Christina"}
	dolores := Person{name: "Dolores"}

	alex.candidates = []string{"Dolores", "Bradley"}
	bradley.candidates = []string{"Christina", "Dorlores"}
	christina.candidates = []string{"Alex", "Bradley"}
	dolores.candidates = []string{"Bradley", "Alex"}

	a := []Person{alex, bradley}
	b := []Person{christina, dolores}

	correct := make(map[string]string)
	correct["Alex"] = "Dolores"
	correct["Bradley"] = "Christina"

	matches := GaleShapley(a, b)
	for _, pair := range matches {
		if correct[pair.name] != pair.partner.name {
			t.Fatalf("Incorrect pairing, expected %s to pair with %s. received: %s", pair.name, correct[pair.name], pair.partner.name)
		}
	}
}

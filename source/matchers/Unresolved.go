package matchers

import "strings"

type Unresolved struct {
	Candidates []Package
}

func NewUnresolved() Unresolved {

	var unresolved Unresolved

	unresolved.Candidates = make([]Package, 0)

	return unresolved

}

func ToUnresolved(value string) Unresolved {

	var unresolved Unresolved

	unresolved.Candidates = make([]Package, 0)
	unresolved.Parse(value)

	return unresolved

}

func (unresolved *Unresolved) Parse(value string) {

	if strings.Contains(value, " || ") {

		values := strings.Split(value, " || ")

		for v := 0; v < len(values); v++ {

			candidate := ToPackage(strings.TrimSpace(values[v]))

			if candidate.IsValid() {
				unresolved.Candidates = append(unresolved.Candidates, candidate)
			}

		}

	} else if strings.Contains(value, " | ") {

		values := strings.Split(value, " | ")

		for v := 0; v < len(values); v++ {

			candidate := ToPackage(strings.TrimSpace(values[v]))

			if candidate.IsValid() {
				unresolved.Candidates = append(unresolved.Candidates, candidate)
			}

		}

	} else {

		candidate := ToPackage(strings.TrimSpace(value))

		if candidate.IsValid() {
			unresolved.Candidates = append(unresolved.Candidates, candidate)
		}

	}

}

func (unresolved *Unresolved) Matches(name string, version string, manager string, vendor string) bool {

	var matches bool = false

	for c := 0; c < len(unresolved.Candidates); c++ {

		candidate := unresolved.Candidates[c]

		if candidate.Matches(name, version, manager, vendor) {
			matches = true
			break
		}

	}

	return matches

}

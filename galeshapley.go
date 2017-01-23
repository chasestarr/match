package match

// Person struct represents an element of a set
type Person struct {
	name           string
	candidates     []string
	candidateIndex int
	partner        *Person
}

func (p *Person) rank(c *Person) int {
	for i := range p.candidates {
		if p.candidates[i] == c.name {
			return i
		}
	}
	return len(p.candidates) + 1
}

func (p *Person) prefers(c Person) bool {
	if p.partner == nil {
		return true
	}

	if p.rank(&c) < p.rank(p.partner) {
		return true
	}
	return false
}

func (p *Person) nextCandidate(candidates []Person) *Person {
	if p.candidateIndex >= len(p.candidates) {
		return nil
	}

	name := p.candidates[p.candidateIndex]

	var candidate *Person
	for i := range candidates {
		if candidates[i].name == name {
			candidate = &candidates[i]
		}
	}

	p.candidateIndex++
	return candidate
}

func (p *Person) partnerTo(c *Person) {
	if c.partner != nil {
		c.partner.partner = nil
	}
	c.partner = p
	if p.partner != nil {
		p.partner.partner = nil
	}
	p.partner = c
}

// GaleShapley returns a stable pairing of two sets
// based on Gale-Shapley algorithm:
// http://cramton.umd.edu/market-design/gale-shapley-college-admissions.pdf
func GaleShapley(a, b []Person) []Person {
	var done bool
	for !done {
		done = true
		for i := range a {
			current := &a[i]
			if current.partner == nil {
				done = false
				candidate := current.nextCandidate(b)
				if candidate.prefers(a[i]) {
					current.partnerTo(candidate)
				}
			}
		}
	}
	return a
}

## [Match](http://cramton.umd.edu/market-design/gale-shapley-college-admissions.pdf)

[![GoDoc](https://godoc.org/github.com/chasestarr/match?status.svg)](https://godoc.org/github.com/chasestarr/match)

```go
import (
  "fmt"

  "github.com/chasestarr/match"
)

func main() {
	alex := Person{name: "Bradley"}
	bradley := Person{name: "Bradley"}
	christina := Person{name: "Christina"}
	dolores := Person{name: "Dolores"}

	alex.candidates = []string{"Dolores", "Bradley"}
	bradley.candidates = []string{"Christina", "Dorlores"}
	christina.candidates = []string{"Alex", "Bradley"}
	dolores.candidates = []string{"Bradley", "Alex"}

	a := []Person{alex, bradley}
	b := []Person{christina, dolores}

	matches := GaleShapley(a, b)
	fmt.Println(matches) // "Alex" --> "Dolores", "Bradley" --> "Christina"
}
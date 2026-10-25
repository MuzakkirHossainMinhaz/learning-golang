package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Welcome to Maps in Go")

	mapBasics()
	lookupAndDelete()
	iteration()
	keyTypes()
	nestedMaps()
	mapFunctionArguments()
	mapsPackageFunctions()
}

// A map associates keys with values. Keys must be comparable, and maps are
// reference-like: copying a map value does not copy its underlying map data.
func mapBasics() {
	fmt.Println("\n=== 1. Creating, Adding and Updating Entries to Maps ===")

	// A map literal creates and initializes a map.
	languages := map[string]string{
		"JS": "JavaScript",
		"RB": "Ruby",
		"PY": "Python",
	}
	fmt.Printf("Languages: %v (type %T)\n", languages, languages)

	// make creates an empty map that can accept entries.
	countries := make(map[string]string)
	countries["BD"] = "Bangladesh"
	countries["JP"] = "Japan"
	countries["BD"] = "People's Republic of Bangladesh" // Existing key: update.

	fmt.Println("Countries:", countries)
	fmt.Println("Number of countries:", len(countries))

	// The zero value of a map is nil. You may read from it, but writing to a
	// nil map panics; use make or a map literal before adding entries.
	var nilMap map[string]int
	fmt.Printf("nil map: %v, is nil: %t, len: %d\n",
		nilMap, nilMap == nil, len(nilMap))
	fmt.Println("Reading missing key from nil map:", nilMap["missing"])
}

func lookupAndDelete() {
	fmt.Println("\n=== 2. Looking-up, Checking, and Deleting Entries ===")

	ages := map[string]int{"Ava": 28, "Noah": 0}

	// A single-value lookup returns the value's zero value for a missing key.
	fmt.Println("Ava's age:", ages["Ava"])
	fmt.Println("Missing person's age:", ages["Mina"])

	// The comma-ok form distinguishes a missing key from a key whose value is zero.
	age, found := ages["Noah"]
	fmt.Printf("Noah: age %d, key exists: %t\n", age, found)
	age, found = ages["Mina"]
	fmt.Printf("Mina: age %d, key exists: %t\n", age, found)

	delete(ages, "Ava") // Deleting a missing key is safe.
	delete(ages, "unknown")
	fmt.Println("After delete:", ages)

	// clear removes every entry but leaves the map non-nil and reusable.
	clear(ages)
	fmt.Printf("After clear: %v, len: %d, is nil: %t\n",
		ages, len(ages), ages == nil)
	ages["Kai"] = 31
	fmt.Println("Reusable map:", ages)
}

func iteration() {
	fmt.Println("\n=== 3. Iterating Over Maps ===")

	scores := map[string]int{
		"Charlie": 91,
		"Alice":   98,
		"Bob":     85,
	}

	// Map iteration order is unspecified and can differ between iterations.
	for name, score := range scores {
		fmt.Printf("%s scored %d\n", name, score)
	}

	// Sort keys when output needs a predictable order.
	names := make([]string, 0, len(scores))
	for name := range scores {
		names = append(names, name)
	}
	slices.Sort(names)
	fmt.Println("Names in sorted order:")
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, scores[name])
	}

	// Iterating over a nil map is safe and executes zero times.
	var nilMap map[string]int
	for key, value := range nilMap {
		fmt.Println("This does not run:", key, value) // never executes/reached
	}
}

func keyTypes() {
	fmt.Println("\n=== 4. Comparable Map Keys ===")

	// Strings, numbers, pointers, and structs whose fields are all comparable
	// can be map keys. Slices, maps, and functions cannot be map key types.
	type coordinate struct {
		X, Y int
	}
	locations := map[coordinate]string{
		{X: 0, Y: 0}: "origin",
		{X: 3, Y: 4}: "lookout",
	}
	fmt.Println("Location at (3, 4):", locations[coordinate{X: 3, Y: 4}])

	// This does not compile because a slice is not comparable:
	// invalid := map[[]int]string{}
}

func nestedMaps() {
	fmt.Println("\n=== 5. Maps Containing Slices and other Maps ===")

	// A map value can be any type, including a slice or another map.
	courses := map[string][]string{
		"Go":         {"arrays", "slices"},
		"JavaScript": {"objects", "functions"},
	}
	courses["Go"] = append(courses["Go"], "maps")
	fmt.Println("Courses:", courses)

	// Inner maps must be initialized before assigning their entries.
	teams := map[string]map[string]int{
		"Falcons": {"wins": 8, "losses": 2},
	}
	teams["Falcons"]["wins"]++
	teams["Sharks"] = make(map[string]int)
	teams["Sharks"]["wins"] = 6
	fmt.Println("Teams:", teams)
}

// Maps passed to functions refer to the same underlying map, so updates to
// entries are visible to the caller. A nil map still cannot be assigned into.
func addScore(scores map[string]int, name string, score int) {
	scores[name] = score
}

func mapFunctionArguments() {
	fmt.Println("\n=== 6. Maps as Function Arguments ===")

	scores := map[string]int{"Ava": 90}
	addScore(scores, "Mina", 96)
	fmt.Println("Map after function call:", scores)
}

func mapsPackageFunctions() {
	fmt.Println("\n=== 7. The standard-library Maps Package ===")

	original := map[string]int{"Go": 10, "Rust": 8, "C": 7}

	// Clone makes a shallow copy. Keys and values are copied, but reference
	// values such as slices or pointers inside the map are not deep-copied.
	clone := maps.Clone(original)
	clone["Go"] = 100
	fmt.Println("Original after changing clone:", original)
	fmt.Println("Changed clone:", clone)
	fmt.Println("Equal to original:", maps.Equal(original, clone))

	// Copy inserts source entries into destination and overwrites matching keys.
	destination := map[string]int{"C": 1, "Python": 5}
	maps.Copy(destination, original)
	fmt.Println("After Copy:", destination)

	// DeleteFunc removes each key/value pair for which the predicate returns true.
	filtered := maps.Clone(original)
	maps.DeleteFunc(filtered, func(_ string, score int) bool {
		return score < 8
	})
	fmt.Println("After DeleteFunc (scores below 8):", filtered)

	// EqualFunc compares values using a caller-provided equality function.
	// This is useful when map values are not comparable, such as slices.
	tagsA := map[string][]string{"language": {"GO", "TOOLS"}}
	tagsB := map[string][]string{"language": {"go", "tools"}}
	equalIgnoringCase := maps.EqualFunc(tagsA, tagsB, func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if !strings.EqualFold(a[i], b[i]) {
				return false
			}
		}
		return true
	})
	fmt.Println("Equal ignoring case:", equalIgnoringCase)

	// All, Keys, and Values return sequences that can be ranged over.
	fmt.Println("Entries from maps.All:")
	for key, value := range maps.All(original) {
		fmt.Printf("%s => %d\n", key, value)
	}
	fmt.Print("Keys:")
	for key := range maps.Keys(original) {
		fmt.Print(" ", key)
	}
	fmt.Println()
	fmt.Print("Values:")
	for value := range maps.Values(original) {
		fmt.Print(" ", value)
	}
	fmt.Println()

	// Collect builds a map from a sequence; Insert adds a sequence's entries
	// into a map, overwriting entries with matching keys.
	collected := maps.Collect(maps.All(original))
	fmt.Println("Collected map:", collected)

	maps.Insert(destination, maps.All(map[string]int{"Go": 11, "Python": 12}))
	fmt.Println("After Insert:", destination)
}

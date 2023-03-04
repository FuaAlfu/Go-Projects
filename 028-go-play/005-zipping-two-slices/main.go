package main

type A struct {
	Name  string
	Age   string
	Roles []B
}

type B struct {
	PersonName string
	RoleName   string
	Date       string
}

type NormalizedName = string

// TODO: Left as an exercise to the reader
func NormalizeName(name string) NormalizedName { return name }

func Zip(sliceA []A, sliceB []B) []A {
	bCache := make(map[string][]B)
	for _, b := range sliceB {
		// I assume there can be more than one Role for each Person.
		// I would not recommend using PersonName as a key for roles because names are supplied by a user and can have varying case.
		// I would recommend that you give 'B' and 'A' both a concept of an ID, and use that to identify your users.
		// Otherwise, you need to find a way to normalize b.PersonName and a.Name.
		normalized := NormalizeName(b.PersonName)
		bCache[normalized] = append(bCache[normalized], b)
	}

	combined := make([]A, len(sliceA))
	// This creates a copy of sliceA, presumably because you don't want to modify the original slice.
	// slice sliceA is []A and not []*A, this will copy all of the values in sliceA too, so modifying the values in combined wont affect the values in sliceA.
	copy(combined, sliceA)
	for i := 0; i < len(combined); i++ {
		// No need to create a new instance of B here - you can just use the version that's in bCache.
		// because bCache is map[string]B (and not map[string]*B), changes to B in the returned slice will not be reflected
		// in map[string]*B.
		roles, ok := bCache[NormalizeName(combined[i].Name)]
		if ok {
			combined[i].Roles = roles
		}
	}

	return combined
}

func main() {
	var sliceA []A
	var sliceB []B

	Zip(sliceA, sliceB)
}
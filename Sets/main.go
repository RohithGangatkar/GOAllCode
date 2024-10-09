package main

import "fmt"

type SetContainer struct {
	sets map[string]struct{}
}

func newSet() *SetContainer {
	return &SetContainer{
		sets: make(map[string]struct{}),
	}
}

func (s *SetContainer) Add(value string) {
	s.sets[value] = struct{}{}
}

func (s *SetContainer) Remove(value string) {
	exists := s.Exists(value)
	if exists {
		delete(s.sets, value)
	} else {
		fmt.Println("element not exist")
	}
}

func (s *SetContainer) Exists(value string) bool {

	_, exists := s.sets[value]
	return exists

}

func (s *SetContainer) Union(other *SetContainer) *SetContainer {
	unionSet := newSet()
	for eachSet := range s.sets {
		unionSet.Add(eachSet)
	}

	for eachOtherSet := range other.sets {
		unionSet.Add(eachOtherSet)
	}
	return unionSet
}

func (s *SetContainer) Intersection(other *SetContainer) *SetContainer {
	intersectionSets := newSet()
	for eachSet := range s.sets {
		if other.Exists(eachSet) {
			intersectionSets.Add(eachSet)
		}
	}
	return intersectionSets
}

func (s *SetContainer) Difference(other *SetContainer) *SetContainer {
	differenceSets := newSet()
	for eachSet := range s.sets {
		if !other.Exists(eachSet) {
			differenceSets.Add(eachSet)
		}
	}
	return differenceSets
}
func (s *SetContainer) Elements() []interface{} {
	keys := make([]interface{}, 0, len(s.sets))
	for k := range s.sets {
		keys = append(keys, k)
	}
	return keys
}
func main() {

	setList := SetContainer{
		sets: make(map[string]struct{}),
	}
	// newSet()

	setList.Add("a")
	setList.Add("a") // not added
	setList.Add("b")
	setList.Remove("b")
	fmt.Println(setList)

	set1 := newSet()
	set1.Add("a")
	set1.Add("b")
	set1.Add("c")

	set2 := newSet()
	set2.Add("d")
	set2.Add("e")
	set2.Add("f")
	set2.Add("a")

	fmt.Printf("Union set %v: \n", set1.Union(set2))
	fmt.Printf("Intersection set %v: \n", set1.Intersection(set2))
	fmt.Printf("Intersection set %v: \n", set1.Difference(set2))

}

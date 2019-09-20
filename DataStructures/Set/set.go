//Package set https://flaviocopes.com/golang-data-structure-set/
package set

//Item of Set can be any type
type Item interface{}

//Set it's easy to model sets with a map in golang
type Set struct {
	item map[Item]bool
}

//Add add a Item to the Set, return a pointer to the Set.
func (s *Set) Add(t Item) *Set {
	if s.item == nil {
		s.item = make(map[Item]bool)
	}
	if _, ok := s.item[t]; !ok {
		s.item[t] = true
	}
	return s
}

//Clear remove all elements from the Set
func (s *Set) Clear() {
	s.item = make(map[Item]bool)
}

//Delete remove the item from the Set if has
func (s *Set) Delete(t Item) {
	if s.Has(t) {
		delete(s.item, t)
	}
}

//Has return true if the Set contains the item
func (s *Set) Has(t Item) bool {
	_, ok := s.item[t]
	return ok
}

//Traversal return all item stored
func (s *Set) Traversal() []Item {
	item := []Item{}
	for i := range s.item {
		item = append(item, i)
	}
	return item
}

//Size of set
func (s *Set) Size() int {
	return len(s.item)
}

// Union returns a new set with elements from both
// the given sets
func (s *Set) Union(s2 *Set) *Set {
	s3 := Set{}
	s3.item = make(map[Item]bool)
	for i := range s.item {
		s3.item[i] = true
	}
	for i := range s2.item {
		_, ok := s3.item[i]
		if !ok {
			s3.item[i] = true
		}
	}
	return &s3
}

// Intersection returns a new set with elements that exist in
// both sets
func (s *Set) Intersection(s2 *Set) *Set {
	s3 := Set{}
	s3.item = make(map[Item]bool)
	for i := range s2.item {
		_, ok := s.item[i]
		if ok {
			s3.item[i] = true
		}
	}
	return &s3
}

// Difference returns a new set with all the elements that
// exist in the first set and don't exist in the second set
func (s *Set) Difference(s2 *Set) *Set {
	s3 := Set{}
	s3.item = make(map[Item]bool)
	for i := range s.item {
		_, ok := s2.item[i]
		if !ok {
			s3.item[i] = true
		}
	}
	return &s3
}

// Subset returns true if s is a subset of s2
func (s *Set) Subset(s2 *Set) bool {
	for i := range s.item {
		_, ok := s2.item[i]
		if !ok {
			return false
		}
	}
	return true
}

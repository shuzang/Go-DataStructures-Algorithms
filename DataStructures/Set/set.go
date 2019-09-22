//Package set https://flaviocopes.com/golang-data-structure-set/
package set

//Item of Set can be any type
type Item interface{}

//Set it's easy to model sets with a map in golang
type Set struct {
	item map[Item]bool
}

//IsMember 判断元素是否是集合的成员
func (s *Set) IsMember(t Item) bool {
	_, ok := s.item[t]
	return ok
}

//Insert 插入成功返回0，如果插入的成员在集合中已存在返回-1
func (s *Set) Insert(t Item) interface{} {
	if s.item == nil {
		s.item = make(map[Item]bool)
	}
	if !s.IsMember(t) {
		s.item[t] = true
		return 0
	}
	return -1
}

//Remove 移除元素
func (s *Set) Remove(t Item) {
	if s.IsMember(t) {
		delete(s.item, t)
	}
}

//Clear 移除集合中的所有元素
func (s *Set) Clear() {
	s.item = make(map[Item]bool)
}

//Size 返回集合中元素的个数
func (s *Set) Size() int {
	return len(s.item)
}

// Union 求两个集合的并集
func (s *Set) Union(s2 *Set) *Set {
	tmps := Set{}
	tmps.item = make(map[Item]bool)
	for key := range s.item {
		tmps.item[key] = true
	}
	for key := range s2.item {
		if _, ok := tmps.item[key]; !ok {
			tmps.item[key] = true
		}
	}
	return &tmps
}

// Intersection 求两个集合的交集
func (s *Set) Intersection(s2 *Set) *Set {
	tmps := Set{}
	tmps.item = make(map[Item]bool)
	for key := range s2.item {
		if s.IsMember(key) {
			tmps.item[key] = true
		}
	}
	return &tmps
}

// Difference 求s和s2两个集合的差集
func (s *Set) Difference(s2 *Set) *Set {
	tmps := Set{}
	tmps.item = make(map[Item]bool)
	for key := range s.item {
		if !s2.IsMember(key) {
			tmps.item[key] = true
		}
	}
	return &tmps
}

// Subset 判断s2是否是s的子集
func (s *Set) Subset(s2 *Set) bool {
	for key := range s2.item {
		if !s.IsMember(key) {
			return false
		}
	}
	return true
}

//Equal 判断两个集合是否相等
func (s *Set) Equal(s2 *Set) bool {
	if len(s.item) == len(s2.item) {
		for key := range s2.item {
			if !s.IsMember(key) {
				return false
			}
		}
		return true
	}
	return false
}

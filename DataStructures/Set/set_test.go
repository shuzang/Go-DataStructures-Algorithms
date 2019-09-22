package set

import (
	"fmt"
	"testing"
)

var set Set

func init() {
	set = Set{}
	return &set
}

func BenchmarkInsert(b *testing.B) {
	var flag int
	for i := 0; i < b.N; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)
	if result := set.Insert("item1"); result == -1 {
		fmt.Println("Item already in Set!")
	}
}

func TestClear(t *testing.T) {
	var flag int
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)
	set.Clear()
	if size := set.Size(); size != 0 {
		fmt.Println("Clera Set failed!")
	}
}

func BenchmarkRemove(b *testing.B) {
	var flag int
	for i := 0; i < b.N; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)
	for i := 0; i < b.N; i++ {
		set.Remove(fmt.Sprintf("item%d", i))
	}
}

func TestIsMember(t *testing.T) {
	var flag int
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)
	if set.IsMember("item2") {
		fmt.Println("item2 is the number of the Set!")
	}
	set.Remove("item2")
	if set.IsMember("item2") {
		fmt.Println("item2 is not the number of the Set now!")
	}
}

func TestSize(t *testing.T) {
	var flag int
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag, set.Size())
}

func TestUnion(t *testing.T) {
	var flag int
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	set2 := Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set2.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	set3 := set.Union(set2)
	fmt.Println(len(set3))
}

func TestIntersection(t *testing.T) {
	var flag int
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	set2 := Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set2.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	set3 := set.Intersection(set2)
	fmt.Println(len(set3))
}

func TestDifference(t *testing.T) {
	var flag int
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	set2 := Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set2.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	set3 := set.Difference(set2)
	fmt.Println(len(set3))
}

func TestSubset(t *testing.T) {
	var flag int
	for i := 0; i < 5; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	set2 := Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set2.Insert(fmt.Sprintf("item%d", i)); ok == 0 {
			flag = ok
		}
	}
	fmt.Println(flag)

	result := set.Subset(set2)
	fmt.Println(result)
}

package set

import (
	"fmt"
	"testing"
)

var set *Set

func init() {
	set = &Set{}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}
	if result := set.Insert("item1"); result != -1 {
		fmt.Println("Insert Succeed")
	}
}

func TestClear(t *testing.T) {
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}
	set.Clear()
	if size := set.Size(); size != 0 {
		fmt.Println("Clera Set failed!")
	}
}

func BenchmarkRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}
	for i := 0; i < b.N; i++ {
		set.Remove(fmt.Sprintf("item%d", i))
	}
}

func TestIsMember(t *testing.T) {
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}
	if set.IsMember("item2") {
		fmt.Println("item2 is the number of the Set!")
	}
	set.Remove("item2")
	if set.IsMember("item2") {
		fmt.Println("item2 is not the number of the Set now!")
	}
}

func TestSize(t *testing.T) {
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}
	fmt.Println("The size of Set:", set.Size())
}

func TestUnion(t *testing.T) {
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	set2 := &Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	set3 := set.Union(set2)
	fmt.Println("the size of union set:", len(set3.item))
}

func TestIntersection(t *testing.T) {
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	set2 := &Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	set3 := set.Intersection(set2)
	fmt.Println("the size of Intersection set:", len(set3.item))
}

func TestDifference(t *testing.T) {
	for i := 0; i < 3; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	set2 := &Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	set3 := set.Difference(set2)
	fmt.Println("the size of difference set", len(set3.item))
}

func TestSubset(t *testing.T) {
	for i := 0; i < 5; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	set2 := &Set{}
	set2.item = make(map[Item]bool)
	for i := 2; i < 5; i++ {
		if ok := set.Insert(fmt.Sprintf("item%d", i)); ok != 0 {
			fmt.Println("Item already in Set!")
		}
	}

	result := set.Subset(set2)
	fmt.Println("Is set2 is the subset of set:", result)
}

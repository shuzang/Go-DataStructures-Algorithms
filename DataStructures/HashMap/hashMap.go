package hashmap

type hashMap interface {
	Init(interface{})
	Lookup(interface{}) bool
	Insert(interface{})
	Remove(interface{})
}

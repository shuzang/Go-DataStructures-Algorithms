package hashmap

type item interface{}

type openHashMap struct {
	hashMap
	table map[item]bool
}

func (h *openHashMap) Init(cap int) {

}

func (h *openHashMap) Destory() {

}

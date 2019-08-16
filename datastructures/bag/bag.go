package bag

type Bag struct {
	first *node
	count int
}

type node struct {
	item *string
	next *node
}

func (b *Bag) Add(item *string) {
	oldFirst := b.first
	b.first = &node{
		item: item,
		next: oldFirst,
	}
	b.count++
}

// TODO: implement some kind of iterator

func (b *Bag) Size() int {
	return b.count
}

func (b *Bag) IsEmpty() bool {
	return b.count == 0
}

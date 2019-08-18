package bag

// Bag struct
type Bag struct {
	first *node
	count int
}

type node struct {
	item *string
	next *node
}

// NewBag is a constructor
func NewBag() *Bag {
	return &Bag{first: nil, count: 0}
}

// Add an item to the bag
func (b *Bag) Add(item string) {
	oldFirst := b.first
	b.first = &node{
		item: &item,
		next: oldFirst,
	}
	b.count++
}

// BagIterator returns a channel (to be used as an iterator on the bag content)
func (b *Bag) BagIterator() <-chan *string {
	ch := make(chan *string)
	go func() {
		current := b.first
		for current != nil {
			item := current.item
			current = current.next
			ch <- item
		}
		close(ch)
	}()
	return ch
}

// Size returns the current size of the bag
func (b *Bag) Size() int {
	return b.count
}

// IsEmpty checks if the bag is empty
func (b *Bag) IsEmpty() bool {
	return b.count == 0
}

package main

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
}

type PeekingIterator struct {
	iter            *Iterator
	cache           int
	isNextAvaliable bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	if !iter.hasNext() {
		return &PeekingIterator{
			isNextAvaliable: false,
		}
	}
	return &PeekingIterator{
		cache:           iter.next(),
		iter:            iter,
		isNextAvaliable: true,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.isNextAvaliable
}

func (this *PeekingIterator) next() int {
	if !this.isNextAvaliable {
		panic("I think we should be panic here")
	}
	cache := this.cache
	this.isNextAvaliable = this.iter.hasNext()
	if this.isNextAvaliable {
		this.cache = this.iter.next()
	}
	return cache
}

func (this *PeekingIterator) peek() int {
	return this.cache
}

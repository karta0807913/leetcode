package main

type Spanner struct {
	Price int
	Span  int
}

type StockSpanner struct {
	data []Spanner
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

// this.data is desc
func (this *StockSpanner) Search(price int) int {
	left, right := 0, len(this.data)
	for left < right {
		mid := (left + right) / 2
		midPrice := this.data[mid].Price
		if midPrice > price {
			left = mid + 1
		} else if midPrice < price {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right
}

func (this *StockSpanner) Next(price int) int {
	n := this.Search(price)
	if n == len(this.data) {
		this.data = append(this.data, Spanner{
			Price: price,
			Span:  1,
		})
		return 1
	}
	this.data[n].Price = price
	this.data[n].Span = 0
	for _, data := range this.data[n+1:] {
		this.data[n].Span += data.Span
	}
	this.data = this.data[:n+1]
	this.data[n].Span++
	return this.data[n].Span
}

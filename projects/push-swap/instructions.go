package main

func pa(a, b *Stack) {
	top, ok := b.pop()
	if ok {
		a.push(top)
	}
}
func pb(a, b *Stack) {
	top, ok := a.pop()
	if ok {
		b.push(top)
	}
}
func sa(a *Stack) {
	if len(a.data) >= 2 {
		n := len(a.data)
		a.data[n-1], a.data[n-2] = a.data[n-2], a.data[n-1]
	}
}
func sb(b *Stack) {
	if len(b.data) >= 2 {
		n := len(b.data)
		b.data[n-1], b.data[n-2] = b.data[n-2], b.data[n-1]
	}
}
func ss(a, b *Stack) {
	sa(a)
	sb(b)
}

func ra(a *Stack) {
	if len(a.data) > 1 {
		top := a.data[len(a.data)-1]
		a.data = a.data[:len(a.data)-1]
		a.data = append([]int{top}, a.data...)
	}
}
func rb(b *Stack) {
	if len(b.data) > 1 {
		top := b.data[len(b.data)-1]
		b.data = b.data[:len(b.data)-1]
		b.data = append([]int{top}, b.data...)
	}
}
func rr(a, b *Stack) {
	ra(a)
	rb(b)
}
func rra(a *Stack) {
	if len(a.data) > 0 {
		elem := a.data[0]
		a.data = a.data[1:]
		a.data = append(a.data, elem)
	}
}
func rrb(b *Stack) {
	if len(b.data) > 0 {
		elem := b.data[0]
		b.data = b.data[1:]
		b.data = append(b.data, elem)
	}
}
func rrr(a, b *Stack) {
	rra(a)
	rrb(b)
}

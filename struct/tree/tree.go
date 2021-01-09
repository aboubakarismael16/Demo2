package main

import "fmt"

type tree struct {
	value int
	left,right *tree
}

func Sort(values []int)  {
	var root *tree
	for _,v := range values {
		root = add(root,v)
	}
	appendValues(values[:0],root)
}

//appendValues appends the element of t in values in order
//and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values,t.left)
		values = append(values,t.value)
		values = appendValues(values,t.right)
	}
	return values

}

func add(t *tree,value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left,value)
	} else {
		t.right = add(t.right,value)
	}
	return t

}

func main() {
	s := []int{20,30,10,5,3,1,0}
	Sort(s)
	fmt.Println(s)

}

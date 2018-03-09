package main

import "fmt"

//input structure
//if lenght obj.next > 0 it's an array of obj otherwise the struct is an integer
type obj struct {
	value int
	next  []obj
}

func flat(r *[]int, o *obj) {
	n := len(o.next)
	if n == 0 {
		*r = append(*r, o.value)
		return
	}
	for i := 0; i < n; i++ {
		flat(r, &o.next[i])
	}
}

func main() {
	var result []int
	//root sample of composite array [1,[2,3],4]
	var root obj
	root.next = make([]obj, 3)
	root.next[0].value = 1
	root.next[1].next = make([]obj, 2)
	root.next[2].value = 4
	root.next[1].next[0].value = 2
	root.next[1].next[1].value = 3
	flat(&result, &root)
	fmt.Println(result)
}

package main

import "testing"

func TestFlat(t *testing.T) {
	output := []int{1, 2, 3, 4} //output expected
	var result []int            //result of flat function
	var root obj                //input nested array
	root.next = make([]obj, 3)
	root.next[0].value = 1
	root.next[1].next = make([]obj, 2)
	root.next[2].value = 4
	root.next[1].next[0].value = 2
	root.next[1].next[1].value = 3
	flat(&result, &root)
	if len(result) != len(output) {
		t.Error("Array lenght", len(result),
			"expected", len(output))
		return
	}
	for i := 0; i < len(output); i++ {
		if output[i] != result[i] {
			t.Error("For", i,
				"expected", output[i],
				"got", result[i])
			return
		}
	}
}

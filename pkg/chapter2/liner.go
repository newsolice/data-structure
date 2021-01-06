package main

import "fmt"

const maxSize = 10

type Element int
type SeqList struct {
	elements [maxSize]Element
	last     int
}

func (s *SeqList) Locate(ele Element) (index int) {
	index = -1
	i := 0
	for i < s.last && s.elements[i] != ele {
		i++
	}
	if i < s.last {
		index = i
	}
	return
}

func main() {
	var s SeqList
	locate := s.Locate(3)
	fmt.Println(locate)
}

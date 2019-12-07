package file

import "sort"

// Sort by name first then with special order.
// Special order should be generated by the order package.
//
// This is based on how gitignore.io sorts the result:
// https://github.com/toptal/gitignore.io/blob/83dbcff2f7e3e084807ab64f7b545d256df7ca64/Sources/App/RouteHandlers/APIRouteHandlers.swift#L131-L132
func Sort(items []string, specialOrder map[string]int) []string {
	s := sorter{
		items: items,
		less:  []lessFn{bySpecial(specialOrder), byCanon},
	}
	sort.Sort(&s)
	return items
}

type lessFn func(a, b string) bool

type sorter struct {
	items []string
	less  []lessFn
}

func (s *sorter) Len() int {
	return len(s.items)
}

func (s *sorter) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s *sorter) Less(i, j int) bool {
	a, b := s.items[i], s.items[j]

	var k int
	for k = 0; k < len(s.less)-1; k++ {
		less := s.less[k]

		switch {
		case less(a, b):
			return true
		case less(b, a):
			return false
		}
	}

	return s.less[k](a, b)
}

func byCanon(a, b string) bool {
	return Canon(a) < Canon(b)
}

func bySpecial(specialOrder map[string]int) func(a, b string) bool {
	return func(a, b string) bool {
		return specialOrder[Canon(a)] < specialOrder[Canon(b)]
	}
}

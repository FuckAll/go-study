package iterator

import "testing"

func TestNewListIterator(t *testing.T) {
	testCast := []struct {
		request  []int
		response []int
		err      error
	}{
		{
			request:  []int{1, 2, 4},
			response: []int{1, 2, 4},
			err:      nil,
		},
	}
	for _, test := range testCast {
		concreteContainer := NewConcreteContainer(test.request)
		iterator := concreteContainer.getIterator()
		var cursor int
		for iterator.hasNext() {
			item, err := iterator.currentItem()
			if err != test.err {
				t.Error(err)
			}
			if item != test.response[cursor] {
				t.Errorf("cursor %d, want %d, but get %d", cursor, test.response[cursor], item)
			}
			cursor++
			iterator.next()
		}
	}
}

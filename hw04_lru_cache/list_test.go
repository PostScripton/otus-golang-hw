package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		require.Equal(t, 80, l.Front().Value)
		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]
		require.Equal(t, 70, l.Front().Value)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("only_one_element_in_the_list", func(t *testing.T) {
		l := NewList()

		element := l.PushBack(1234)
		require.Equal(t, 1, l.Len())
		require.Nil(t, element.Prev)
		require.Nil(t, element.Next)

		l.MoveToFront(element)
		require.Nil(t, element.Prev)
		require.Nil(t, element.Next)

		l.Remove(element)
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("links_are_lost_for_detached_element", func(t *testing.T) {
		l := NewList()

		first := l.PushBack(1)
		element := l.PushBack(2)
		last := l.PushBack(3)

		require.Equal(t, 3, l.Len())
		require.Equal(t, first, element.Prev)
		require.Equal(t, last, element.Next)

		l.Remove(element)
		require.Equal(t, last, first.Next)
		require.Equal(t, first, last.Prev)
		require.Nil(t, element.Prev)
		require.Nil(t, element.Next)
	})

	t.Run("remove_edge_element", func(t *testing.T) {
		l := NewList()

		first := l.PushBack(1)
		element := l.PushBack(2)
		last := l.PushBack(3)

		l.Remove(first)
		l.Remove(last)

		require.Equal(t, element, l.Front())
		require.Equal(t, element, l.Back())
	})
}

func TestMyList(t *testing.T) {
	t.Run("big check", func(t *testing.T) {
		l := NewList()

		l.PushFront(20) // [20]
		l.PushBack(30)  // [20, 30]
		l.PushFront(10) // [10, 20, 30]
		l.PushBack(40)  // [10, 20, 30, 40]
		l.PushBack(50)  // [10, 20, 30, 40, 50]
		l.PushBack(60)  // [10, 20, 30, 40, 50, 60]
		require.Equal(t, 6, l.Len())

		l.Remove(l.Front().Next) // [10, 30, 40, 50, 60]
		l.Remove(l.Front().Next) // [10, 40, 50, 60]

		require.Equal(t, 40, l.Front().Next.Value)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 4, l.Len())

		l.MoveToFront(l.Front().Next.Next) // [50, 10, 40, 60]
		require.Equal(t, 50, l.Front().Value)

		l.MoveToFront(l.Back().Prev) // [40, 50, 10, 60]
		l.MoveToFront(l.Back().Prev) // [10, 40, 50, 60]
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 60, l.Back().Value)

		l.PushFront(20)                    // [20, 10, 40, 50, 60]
		l.PushBack(30)                     // [20, 10, 40, 50, 60, 30]
		l.MoveToFront(l.Back())            // [30, 20, 10, 40, 50, 60]
		l.MoveToFront(l.Front().Next)      // [20, 30, 10, 40, 50, 60]
		l.MoveToFront(l.Front().Next.Next) // [10, 20, 30, 40, 50, 60]
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 20, l.Front().Next.Value)
		require.Equal(t, 60, l.Back().Value)
	})
}

package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(is); i++ {
		stack.Push(is[i])
	}
	result := make([]int, 9)
	count := 0
	for {
		v := stack.Pop()
		vi, ok := v.(int)
		if !ok {
			break
		}
		result[count] = vi
		count++
	}
	fmt.Println(result)
}
func TestSafeStack(t *testing.T) {
	stack := NewSafeStack()
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(is); i++ {
		stack.Push(is[i])
	}
	result := make([]int, 9)
	count := 0
	for {
		v := stack.Pop()
		vi, ok := v.(int)
		if !ok {
			break
		}
		result[count] = vi
		count++
	}
	fmt.Println(result)
}
func BenchmarkSafeStack(b *testing.B) {
	stack := NewSafeStack()
	for i := 0; i < b.N; i++ {
		is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < len(is); i++ {
			stack.Push(is[i])
		}
		result := make([]int, 9)
		count := 0
		for {
			v := stack.Pop()
			vi, ok := v.(int)
			if !ok {
				break
			}
			result[count] = vi
			count++
		}
	}
}

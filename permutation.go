package main

import (
	"fmt"
	"strings"
)

// PermutationType interface for permutation
type PermutationType interface {
	Count() int
	List() []string
}

// Permutable implements PermutationType
type Permutable struct {
	Elements []interface{}
}

// Count for PermutationType
func (p *Permutable) Count() int {
	n := len(p.Elements)
	r := 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}

// List all permutation
func (p *Permutable) List() []string {

	n := len(p.Elements)
	tmp := make([]string, n)
	c := make([]int, n)
	result := make([]string, 0)
	for k, item := range p.Elements {
		c[k] = 0
		tmp[k] = fmt.Sprintf("%v", item)
	}

	i := 0
	result = append(result, strings.Join(tmp[:], ","))
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				t := tmp[0]
				tmp[0] = tmp[i]
				tmp[i] = t
			} else {
				t := tmp[c[i]]
				tmp[c[i]] = tmp[i]
				tmp[i] = t
			}
			result = append(result, strings.Join(tmp[:], ","))
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}

	return result
}

// PermutateNum create new permutation type
func PermutateNum(n int) PermutationType {
	t := make([]interface{}, n)
	for i := 0; i < n; i++ {
		t[i] = i
	}
	return &Permutable{Elements: t}
}

// PermutateSlice create new permutation type
func PermutateSlice(n []interface{}) PermutationType {
	return &Permutable{Elements: n}
}

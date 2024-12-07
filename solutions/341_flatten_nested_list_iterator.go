package solutions

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	isInteger bool
	integer   int
	list      []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (ni NestedInteger) IsInteger() bool {
	return ni.isInteger
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (ni NestedInteger) GetInteger() int {
	return ni.integer
}

// Set this NestedInteger to hold a single integer.
func (ni *NestedInteger) SetInteger(value int) {
	ni.isInteger = true
	ni.integer = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (ni *NestedInteger) Add(elem NestedInteger) {
	ni.isInteger = false
	ni.list = append(ni.list, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (ni NestedInteger) GetList() []*NestedInteger {
	return ni.list
}

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	// Flatten the nested list.
	stack := []*NestedInteger{}
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	return &NestedIterator{stack}
}

func (ni *NestedIterator) Next() int {
	// Assume HasNext() was called before Next() and the top element is an integer.
	top := ni.stack[len(ni.stack)-1]
	ni.stack = ni.stack[:len(ni.stack)-1]
	return top.GetInteger()
}

func (ni *NestedIterator) HasNext() bool {
	for len(ni.stack) > 0 {
		top := ni.stack[len(ni.stack)-1]
		if top.IsInteger() {
			return true
		}
		// If the top element is a list, pop it from the stack and push its elements.
		ni.stack = ni.stack[:len(ni.stack)-1]
		for i := len(top.GetList()) - 1; i >= 0; i-- {
			ni.stack = append(ni.stack, top.GetList()[i])
		}
	}
	return false
}

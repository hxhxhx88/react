// Code generated by immutableGen. DO NOT EDIT.

package model

//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

//
// People is an immutable type and has the following template:
//
// 	[]*Person
//
type People struct {
	theSlice []*Person
	mutable  bool
	__tmpl   *_Imm_People
}

var _ immutable.Immutable = new(People)
var _ = new(People).__tmpl

func NewPeople(s ...*Person) *People {
	c := make([]*Person, len(s))
	copy(c, s)

	return &People{
		theSlice: c,
	}
}

func NewPeopleLen(l int) *People {
	c := make([]*Person, l)

	return &People{
		theSlice: c,
	}
}

func (m *People) Mutable() bool {
	return m.mutable
}

func (m *People) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *People) Get(i int) *Person {
	return m.theSlice[i]
}

func (m *People) AsMutable() *People {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *People) dup() *People {
	resSlice := make([]*Person, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &People{
		theSlice: resSlice,
	}

	return res
}

func (m *People) AsImmutable(v *People) *People {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *People) Range() []*Person {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *People) WithMutable(f func(mi *People)) *People {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *People) WithImmutable(f func(mi *People)) *People {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *People) Set(i int, v *Person) *People {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *People) Append(v ...*Person) *People {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
func (s *People) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	if s.Len() == 0 {
		return true
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true

	for _, v := range s.theSlice {
		if v != nil && !v.IsDeeplyNonMutable(seen) {
			return false
		}
	}
	return true
}

//
// Person is an immutable type and has the following template:
//
// 	struct {
// 		Name	string
// 		Age	int
// 	}
//
type Person struct {
	field_Name string
	field_Age  int

	mutable bool
	__tmpl  *_Imm_Person
}

var _ immutable.Immutable = new(Person)
var _ = new(Person).__tmpl

func (s *Person) AsMutable() *Person {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *Person) AsImmutable(v *Person) *Person {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *Person) Mutable() bool {
	return s.mutable
}

func (s *Person) WithMutable(f func(si *Person)) *Person {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *Person) WithImmutable(f func(si *Person)) *Person {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}

func (s *Person) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	return true
}
func (s *Person) Age() int {
	return s.field_Age
}

// SetAge is the setter for Age()
func (s *Person) SetAge(n int) *Person {
	if s.mutable {
		s.field_Age = n
		return s
	}

	res := *s
	res.field_Age = n
	return &res
}
func (s *Person) Name() string {
	return s.field_Name
}

// SetName is the setter for Name()
func (s *Person) SetName(n string) *Person {
	if s.mutable {
		s.field_Name = n
		return s
	}

	res := *s
	res.field_Name = n
	return &res
}
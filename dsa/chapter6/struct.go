package main

import "sort"

type Commit struct {
	username string
	lang     string
	numlines int
}

type lessFunc func(p1, p2 *Commit) bool

type MutilSorter struct {
	Commits   []Commit
	lessFuncs []lessFunc
}

// Less implements sort.Interface.
func (m *MutilSorter) Less(i int, j int) bool {
	p := m.Commits[i]
	q := m.Commits[j]
	var k int
	for k = 0; k < len(m.lessFuncs)-1; k++ {
		less := m.lessFuncs[k]
		switch {
		case less(&p, &q):
			return true
		case less(&q, &p):
			return false

		}
	}
	return m.lessFuncs[k](&p, &q)
}

// Swap implements sort.Interface.
func (m *MutilSorter) Swap(i int, j int) {
	m.Commits[i] = m.Commits[j]
	m.Commits[j] = m.Commits[i]
}

func (m *MutilSorter) Sort(commits []Commit) {
	m.Commits = commits
	sort.Sort(m)
}

func OrderedBy(lessFunction ...lessFunc) *MutilSorter {
	return &MutilSorter{
		lessFuncs: lessFunction,
	}
}

func (m *MutilSorter) Len() int {
	return len(m.Commits)
}

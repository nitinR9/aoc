package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
	z int
}

type DistPairs struct {
	distance int64
	i        int
	j        int
}

type DSU struct {
	parent   []int
	size     []int
	setCount int
}

func distance(p1, p2 Position) int64 {
	d := int64((p2.x-p1.x)*(p2.x-p1.x)) + int64((p2.y-p1.y)*(p2.y-p1.y)) + int64((p2.z-p1.z)*(p2.z-p1.z))
	return d
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &DSU{
		parent:   p,
		size:     s,
		setCount: n,
	}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // path compression
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	// union by size (attach smaller under larger)
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	d.setCount--
	return true
}

func (d *DSU) Count() int { return d.setCount }

func GroupsAfterFirstNConnections(nPts int, distPairs []DistPairs, nConn int) (int, [][]int) {
	if nConn < 0 {
		nConn = 0
	}
	if nConn > len(distPairs) {
		nConn = len(distPairs)
	}

	dsu := NewDSU(nPts)

	// Apply first nConn edges
	for i := 0; i < nConn; i++ {
		e := distPairs[i]
		dsu.Union(e.i, e.j)
	}

	// Collect components
	compMap := make(map[int][]int)
	for v := 0; v < nPts; v++ {
		r := dsu.Find(v)
		compMap[r] = append(compMap[r], v)
	}

	// Flatten to slice, sort members for readability
	groups := make([][]int, 0, len(compMap))
	for _, members := range compMap {
		sort.Ints(members)
		groups = append(groups, members)
	}

	return dsu.Count(), groups
}

func TopKGroupsAfterFirstNConnections(nPts int, distPairs []DistPairs, nConn, k int) ([][]int, int) {
	if k < 1 {
		k = 1
	}
	totalGroups, groups := GroupsAfterFirstNConnections(nPts, distPairs, nConn)

	type comp struct {
		members []int
	}
	comps := make([]comp, 0, len(groups))
	for _, g := range groups {
		comps = append(comps, comp{members: g})
	}

	sort.Slice(comps, func(i, j int) bool {
		if len(comps[i].members) != len(comps[j].members) {
			return len(comps[i].members) > len(comps[j].members)
		}
		return comps[i].members[0] < comps[j].members[0]
	})

	if k > len(comps) {
		k = len(comps)
	}

	top := make([][]int, 0, k)
	for idx := 0; idx < k; idx++ {
		top = append(top, comps[idx].members)
	}

	return top, totalGroups
}

func KruskalMST(nPts int, distPairs []DistPairs) ([]DistPairs, int64) {
	dsu := NewDSU(nPts)
	mst := make([]DistPairs, 0, nPts-1)
	var total int64

	for _, e := range distPairs {
		if dsu.Union(e.i, e.j) {
			mst = append(mst, e)
			total += e.distance
			if len(mst) == nPts-1 {
				break
			}
		}
	}
	return mst, total
}

func part2(input []DistPairs) (int, int) {
	mst, _ := KruskalMST(len(input), input)
	pair := mst[len(mst)-1]
	return pair.i, pair.j
}

func part1(input []DistPairs, connections int) int64 {
	p1 := int64(1)

	top3, _ := TopKGroupsAfterFirstNConnections(len(input), input, connections, 3)
	for _, g := range top3 {
		p1 *= int64(len(g))
	}
	return p1
}

func day8() {
	data := readFile("inputs/day8.txt", "\r\n")
	parsedInput := []Position{}

	for _, line := range data {
		stringVals := strings.Split(string(line), ",")

		xpos, _ := strconv.ParseInt(stringVals[0], 10, 64)
		ypos, _ := strconv.ParseInt(stringVals[1], 10, 64)
		zpos, _ := strconv.ParseInt(stringVals[2], 10, 64)

		parsedInput = append(parsedInput, Position{x: int(xpos), y: int(ypos), z: int(zpos)})
	}

	distPairs := []DistPairs{}
	for i := 0; i < len(parsedInput)-1; i++ {
		for j := i + 1; j < len(parsedInput); j++ {
			d := distance(parsedInput[i], parsedInput[j])
			p := DistPairs{distance: d, i: i, j: j}
			distPairs = append(distPairs, p)
		}
	}

	sort.Slice(distPairs, func(i, j int) bool {
		return distPairs[i].distance < distPairs[j].distance
	})

	fmt.Println("Day 8 part 1:", part1(distPairs, 1000))

	idx1, idx2 := part2(distPairs)

	fmt.Println("Day 8 part 2:", parsedInput[idx1].x*parsedInput[idx2].x)
	fmt.Println("--------------------------")
}
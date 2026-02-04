package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/aclements/go-z3/z3"
)

type Machine struct {
	len int
	target int16
	buttons []int16
	jolts []float64
}

func combine(arr []int16, k int, start int, curr []int16, result *[][]int16) {
    if len(curr) == k {
        combo := make([]int16, k)
        copy(combo, curr)
        *result = append(*result, combo)
        return
    }
    for i := start; i < len(arr); i++ {
        curr = append(curr, arr[i])
        combine(arr, k, i+1, curr, result)
        curr = curr[:len(curr)-1]
    }
}

func allCombinations(arr []int16) [][][]int16 {
    var all [][][]int16
    for size := 1; size <= len(arr); size++ {
        var result [][]int16
        combine(arr, size, 0, []int16{}, &result)
        all = append(all, result)
    }
    return all
}

func (m *Machine) setDiagram(str string) {
	var t int16 = 0
	for _, ch := range str {
		switch ch {
			case '#': t = (t << 1) | 1
			case '.': t = t << 1
		}
	}
	m.target = t
	m.len = len(strings.Trim(str, "[]"))
}

func (m *Machine) setButtons(strArr []string) {
	m.buttons = make([]int16, 0, len(strArr))
	for _, str := range strArr {
		result := int16(0)
		str = strings.Trim(str, "()")
		values := strings.Split(str, ",")
		for _, num := range values {
			idx, _ := strconv.Atoi(strings.TrimSpace(num))
			pos := m.len - 1 - idx
			result |= 1 << pos
		}
		m.buttons = append(m.buttons, result)
	}
}

func (m *Machine) setJolts(str string) {
	joltStr := strings.Split(strings.Trim(str, "{}"), ",")

	for _, jStr := range joltStr {
		jolt, _ := strconv.ParseFloat(strings.TrimSpace(jStr), 64)
		m.jolts = append(m.jolts, jolt)
	}
}

func (m *Machine) getButtonPress() int {
	combos := allCombinations(m.buttons)

	for idx, sets := range combos {
		for _, set := range sets {
			var combined int16 = 0
			for _, val := range set {
				combined ^= val
			}
			if combined == m.target {
				return idx + 1
			}
		}
	}
		
	return -1
}

func (m *Machine) getJolts() int {
	bitsArr := [][]int{}
	for _, button := range m.buttons {
		bits := int16ToBits(button, len(m.jolts))
		bitsArr = append(bitsArr, bits)
	}

	rowMax, colMax := len(m.jolts), len(m.buttons)
	A := make([][]int, rowMax)
	for i := 0; i < rowMax; i++ {
		A[i] = make([]int, colMax)
	}

	for i := 0; i < rowMax; i++ {
		for j := 0; j < colMax; j++ {
			A[i][j] = bitsArr[j][i]
		}
	}

	ctx := z3.NewContext(nil)
	solver := z3.NewSolver(ctx)

	vars := make([]z3.Int, colMax)
	for i := range vars {
		vars[i] = ctx.IntConst(fmt.Sprintf("x%d", i))
		solver.Assert(vars[i].GE(ctx.FromInt(0, ctx.IntSort()).(z3.Int)))
	}

	for i:=0; i < rowMax; i++ {
		var rowExpr z3.Int
		first := true
		for j:=0; j < colMax; j++ {
			if A[i][j] != 0 {
				coeff := ctx.FromInt(int64(A[i][j]), ctx.IntSort()).(z3.Int)
				term := coeff.Mul(vars[j])
				if first {
					rowExpr = term
					first = false
				} else {
					rowExpr = rowExpr.Add(term)
				}
			}
		}

		target := ctx.FromInt(int64(m.jolts[i]), ctx.IntSort()).(z3.Int)
		solver.Assert(rowExpr.Eq(target))
	}

	var lastModel z3.Model
	foundAny := false

	for {
		if ok, err := solver.Check(); err != nil || !ok {
			break
		}
		foundAny = true
		lastModel = *solver.Model()
		var totalSum z3.Int
		currentSumVal := int64(0)
		for i, v := range vars {
			valStr := lastModel.Eval(v, true).String()
			valInt, _ := strconv.ParseInt(valStr, 10, 64)
			currentSumVal += valInt
			if i == 0{
				totalSum = v
			} else {
				totalSum = totalSum.Add(v)
			}
		}

		nextBound := ctx.FromInt(currentSumVal, ctx.IntSort()).(z3.Int)
		solver.Assert(totalSum.LT(nextBound))
	}

	if foundAny {
		total := 0
		for _, v := range vars {
			val, _ := strconv.ParseInt(lastModel.Eval(v, true).String(), 10, 64)
			total += int(val)
		}
		return total
	} else {
		fmt.Println("No solution found")
	}

	return 0
}

func int16ToBits(n int16, size int) []int {
    bits := make([]int, size)
    for i := 0; i < size; i++ {
        // check bit from MSB to LSB within given size
        if (n & (1 << (size-1-i))) != 0 {
            bits[i] = 1
        } else {
            bits[i] = 0
        }
    }
    return bits
}

func day10part1(machines []Machine) {
	total := 0
	for _, machine := range machines {
		presses := machine.getButtonPress()
		total += presses
	}
	fmt.Println("Day 10 Part 1:", total)
}

func day10part2(machines []Machine) {
	totalJolts := 0
	for _, machine := range machines {
		x := machine.getJolts()
		totalJolts += x
	}
	
	fmt.Println("Day 10 Part 2:", totalJolts)
}


func parseDay10Input(data []string) []Machine {
	var machines []Machine
	for _, line := range data {
		values := strings.Split(line, " ")
		diagramStr := values[0]
		buttonsStr := values[1:len(values)-1]
		joltsStr := values[len(values)-1]

		machine := Machine{}
		machine.setDiagram(diagramStr)
		machine.setButtons(buttonsStr)
		machine.setJolts(joltsStr)
		machines = append(machines, machine)
	}
	
	return machines
}

func day10() {
	data := readFile("inputs/day10.txt", "\n")

	machines := parseDay10Input(data)

	day10part1(machines)
	day10part2(machines)
}
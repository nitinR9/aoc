package main

import (
	"2015/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type num uint16

type Operator func(num, num) num

type Wire struct {
	op   string
	args []string
}

var Wires = map[string]Wire{}
var Gates = map[string]num{}

var insRegexp = regexp.MustCompile(`[A-Z]+`)
var varRegexp = regexp.MustCompile(`[a-z|0-9]+`)

func Convert(s string) num {
	val, _ := strconv.ParseInt(s, 10, 64)
	return num(val)
}

func ParseIns(s string) (Wire, string) {
	statement := strings.Split(s, " -> ")
	lhs, rhs := varRegexp.FindAllString(statement[0], -1), statement[1]
	op := ""

	if val := insRegexp.FindString(s); val != "" {
		op = val
	} else {
		op = "ASSIGN"
	}
	return Wire{
		args: lhs,
		op:   op,
	}, rhs
}

func isNumber(s string) (num, bool) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return num(val), true
}

func CalcWire(name string) num {
	wire := Wires[name]

	if val, ok := Gates[name]; ok {
		return val
	}

	if val, ok := isNumber(name); ok {
		return val
	}

	switch wire.op {
	case "AND":
		Gates[name] = CalcWire(wire.args[0]) & CalcWire(wire.args[1])
	case "OR":
		Gates[name] = CalcWire(wire.args[0]) | CalcWire(wire.args[1])
	case "LSHIFT":
		Gates[name] = CalcWire(wire.args[0]) << CalcWire(wire.args[1])
	case "RSHIFT":
		Gates[name] = CalcWire(wire.args[0]) >> CalcWire(wire.args[1])
	case "NOT":
		Gates[name] = ^CalcWire(wire.args[0])
	case "ASSIGN":
		Gates[name] = CalcWire(wire.args[0])
	default:
		fmt.Println("wrong operation")
	}
	return Gates[name]
}

func Part1(input []string) {
	for _, val := range input {
		wire, dest := ParseIns(val)
		Wires[dest] = wire
	}

	fmt.Println("Part1:", CalcWire("a"))
}

func Part2() {
	temp := Gates["a"]
	Gates = map[string]num{}
	Gates["b"] = temp
	fmt.Println("Part2", CalcWire("a"))
}

func main() {
	instructions := common.Process("input.txt")
	Part1(instructions)
	Part2()
}

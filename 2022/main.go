package main

type runner interface {
	Run()
}

type solutions struct {
	day01 runner
	day02 runner
	day03 runner
	day04 runner
	day05 runner
	day06 runner
	day07 runner
	day08 runner
	day09 runner
	day10 runner
	day11 runner
	day12 runner
	day13 runner
	day14 runner
	day15 runner
	day16 runner
	day17 runner
	day18 runner
	day19 runner
	day20 runner
	day21 runner
	day22 runner
	day23 runner
	day24 runner
	day25 runner
	day26 runner
	day27 runner
	day28 runner
	day29 runner
	day30 runner
}

func NewSolutions() solutions {
	return solutions{
		day01: NewDay01(),
		day02: nil,
		day03: nil,
		day04: nil,
		day05: nil,
		day06: nil,
		day07: nil,
		day08: nil,
		day09: nil,
		day10: nil,
		day11: nil,
		day12: nil,
		day13: nil,
		day14: nil,
		day15: nil,
		day16: nil,
		day17: nil,
		day18: nil,
		day19: nil,
		day20: nil,
		day21: nil,
		day22: nil,
		day23: nil,
		day24: nil,
		day25: nil,
		day26: nil,
		day27: nil,
		day28: nil,
		day29: nil,
		day30: nil,
	}

}

func main() {
	solution := NewSolutions()
}

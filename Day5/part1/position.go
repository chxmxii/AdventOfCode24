package main

type P map[int]Set

func p() P {
	return map[int]Set{}
}

func (p P) has(x int) bool {
	_, ok := p[x]
	return ok
}

func (p P) add(r1, r2 int) {
	if _, ok := p[r1]; !ok {
		p[r1] = set()
	}
	if _, ok := p[r2]; !ok {
		p[r2] = set()
	}
	p[r1].add(r2)
}

func (p P) before(r1, r2 int) bool {
	if !p.has(r1) || !p.has(r2) {
		return false
	}
	return p[r1].has(r2)
}

func (p P) rulesMapping(rule []Rule) {
	for _, r := range rule {
		p.add(r.r1, r.r2)
	}
}

func (p P) checkOrder(pages []int) bool {
	for i := 1; i < len(pages)-1; i++ {
		if !p.before(pages[i], pages[i+1]) {
			return false
		}
	}
	return true
}

.PHONY: day1
day1:
	@./Day1/day1 -f Day1/input.txt

.PHONY: day2
day2:
	@./Day2/day2

.PHONY: day3
day3:
	@./Day3/day3

.PHONY: day4
day4:
	@echo "Part1;"
	@./Day4/part1/main -f Day4/part1/input.txt
	@echo "Part2;"
	@./Day4/part2/main -f Day4/part2/input.txt

.PHONY: day5
day5:
	@echo "Part1;"
	@./Day4/part1/main -f Day4/part1/puzzel.input
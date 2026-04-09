DAY   ?= $(day)
INPUT ?= $(if $(filter f,$(MAKECMDGOALS)),full,demo)

.DEFAULT_GOAL := run

run:
	@go build -o ./dist/day$(DAY) ./src/day$(DAY)/main.go
	@echo ""
	@echo "=== Day $(DAY) - $(INPUT) ==="
	@./dist/day$(DAY) < ./data/day$(DAY)/$(INPUT).txt
	@echo ""

new:
	@mkdir -p ./src/day$(DAY) ./data/day$(DAY)
	@touch ./data/day$(DAY)/demo.txt ./data/day$(DAY)/full.txt
	@cp template.go ./src/day$(DAY)/main.go
	@echo "Scaffolded day$(DAY)"

clean:
	rm -rf ./dist

f: run

.PHONY: run new clean f

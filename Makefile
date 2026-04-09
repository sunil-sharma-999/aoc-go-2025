DAY   ?= $(day)
INPUT ?= $(if $(filter f,$(MAKECMDGOALS)),full,demo)

.DEFAULT_GOAL := run

run:
	@go build -o ./dist/day$(DAY) ./src/day$(DAY)/main.go
	@echo "=== Day $(DAY) - $(INPUT) ==="
	@start=$$(python3 -c 'import time;print(int(time.time()*1e9))'); \
	./dist/day$(DAY) < ./data/day$(DAY)/$(INPUT).txt; \
	end=$$(python3 -c 'import time;print(int(time.time()*1e9))'); \
	echo "--- $$(( (end - start) / 1000000 ))ms ---"

new:
	@mkdir -p ./src/day$(DAY) ./data/day$(DAY)
	@touch ./data/day$(DAY)/demo.txt ./data/day$(DAY)/full.txt
	@cp template.go ./src/day$(DAY)/main.go
	@echo "Scaffolded day$(DAY)"

clean:
	rm -rf ./dist

f: run

.PHONY: run new clean f

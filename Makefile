export AOC_COOKIE=

all: run

run:
	DAY=3 go run cmd/aocgo2024/main.go

.SILENT: run
.PHONY: run

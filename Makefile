export AOC_COOKIE=

all: run

run:
	DAY=4 go run cmd/aocgo2024/main.go

.SILENT: run
.PHONY: run

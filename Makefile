export AOC_COOKIE=

all: run

run:
	DAY=1 go run cmd/aocgo2024/main.go

.SILENT: run
.PHONY: run

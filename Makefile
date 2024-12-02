export AOC_COOKIE=53616c7465645f5f6bb908bb8c69212ec34ccb460aa6e4f10ad796e79e6f1189884cd638df538ba9670b626345aa58b95402efacc7bb122756bab5f6a70fcaba

all: run

run:
	DAY=2 go run cmd/aocgo2024/main.go

.SILENT: run
.PHONY: run

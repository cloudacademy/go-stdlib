# ======================
# Strings

builder:
	go run builder.go

formatter:
	go run formatter.go

strings:
	go run strings.go

parse:
	go run ./strings/parse.go

# ======================
# OS

env:
	CLOUDACADEMY_PORT=80 CLOUDACADEMY_HOST=192.168.10.12 go run ./os/env.go

file:
	go run ./os/file.go

exec:
	go run ./os/exec.go

args:
	go run ./os/args.go	red green blue yellow

# ======================
# OS

all: builder formatter strings parse env file exec args math filereader filewriter httpget httppost jsonmarshal jsonunmarshal xmlmarshal xmlunmarshal

# ======================
# strings

.PHONY: builder formatter strings parse

builder:
	go run ./strings/builder.go

formatter:
	go run ./strings/formatter.go

strings:
	go run ./strings/strings.go

parse:
	go run ./strings/parse.go

# ======================
# os

.PHONY: env file exec args

env:
	CLOUDACADEMY_PORT=80 CLOUDACADEMY_HOST=192.168.10.12 go run ./os/env.go

file:
	go run ./os/file.go

exec:
	go run ./os/exec.go

args:
	go run ./os/args.go	red green blue yellow

# ======================
# math

.PHONY: math

math:
	go run ./math/math.go

# ======================
# io

.PHONY: io

filereader:
	go run ./io/filereader.go ./io/sample.txt

filewriter:
	go run ./io/filewriter.go ./io/sampleout.txt

# ======================
# net

.PHONY: httpget httppost

httpget:
	go run ./net/httpget.go

httppost:
	go run ./net/httppost.go

# ======================
# encoding

.PHONY: jsonmarshal jsonunmarshal xmlmarshal xmlunmarshal

jsonmarshal:
	go run ./encoding/jsonmarshal.go

jsonunmarshal:
	go run ./encoding/jsonunmarshal.go

xmlmarshal:
	go run ./encoding/xmlmarshal.go

xmlunmarshal:
	go run ./encoding/xmlunmarshal.go

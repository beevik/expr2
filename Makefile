expr: parse.go tokenize.go main.go
	go build -o $@

parse.go: parse.y
	go tool yacc -p expr -v parse.output -o $@ parse.y
	gofmt -w $@
	rm parse.output

tokenize.go: tokenize.l
	golex -o $@ tokenize.l
	gofmt -w $@

clean:
	rm -f parse.go tokenize.go parse.output

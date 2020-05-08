MEGAAPP := ./cmd/megaapp
BINARY  := megaapp

.PHONY: megaapp
megaapp: test vet
	go build -v -o $(BINARY) $(MEGAAPP)

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm $(BINARY)

.PHONY: run
run:
	go run $(MEGAAPP)

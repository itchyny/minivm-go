DIR=minivm
CMD=cmd
BUILD=build

all: $(DIR)/parser.go $(BUILD)/minivm

$(DIR)/parser.go: $(DIR)/parser.go.y
	go tool yacc -o $@ -v $(DIR)/parser.output $<

build: all

$(BUILD)/minivm: $(DIR)/lex.go $(DIR)/parser.go $(DIR)/opcode.go $(DIR)/value.go $(DIR)/stack.go $(DIR)/variable.go $(DIR)/codegen.go $(DIR)/vm.go $(CMD)/main.go
	go build -o $@ $(CMD)/main.go

test: testdeps build
	go test -v ./$(CMD)...

testdeps:
	go get -d -v -t ./...

LINT_RET = .golint.txt
lint: lintdeps build
	rm -f $(LINT_RET)
	golint ./... | tee $(LINT_RET)
	test ! -s $(LINT_RET)

lintdeps:
	go get -d -v -t ./...
	go get -u github.com/golang/lint/golint

clean:
	rm -rf ./$(BUILD)
	go clean

.PHONY: build test testdeps lint lintdeps clean

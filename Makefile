DIR=minivm
CMD=cmd
BUILD=build

all: $(DIR)/parser.go $(BUILD)/minivm

$(DIR)/parser.go: $(DIR)/parser.go.y
	go tool yacc -o $@ -v $(DIR)/parser.output $<

build: all

$(BUILD)/minivm: $(DIR)/lex.go $(DIR)/parser.go $(DIR)/opcode.go $(DIR)/value.go $(DIR)/codegen.go $(DIR)/vm.go $(CMD)/main.go
	go build -o $@ $(CMD)/main.go

test: testdeps build
	go test -v ./$(CMD)...

testdeps:
	go get -d -v -t ./...

clean:
	rm -rf ./$(BUILD)
	go clean

.PHONY: build test testdeps clean

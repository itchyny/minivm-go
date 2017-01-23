DIR=minivm

all: $(DIR)/parser.go

$(DIR)/parser.go: $(DIR)/parser.go.y
	go tool yacc -o $@ -v $(DIR)/parser.output $<

PB_PATH = $(shell pwd)/proto
DIR = $(shell pwd)/app

SERVICES := gateway task operator
service = $(word 1, $@)
BIN = $(shell pwd)/bin
.PHONY: proto
proto:
	@for file in $(PB_PATH)/*.proto; do \
		protoc -I $(PB_PATH) $$file --micro_out=$(PB_PATH)/pb --go_out=$(PB_PATH)/pb; \
	done
	@for file in $(shell find $(PB_PATH)/pb/* -type f); do \
		protoc-go-inject-tag -input=$$file; \
	done
.PHONY: $(SERVICES)
$(SERVICES):
	go build -o $(BIN)/$(service) $(DIR)/$(service)/cmd

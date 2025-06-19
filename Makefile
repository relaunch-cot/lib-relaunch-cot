PROTO_DIR = proto
PACKAGE = lib-relaunch-cot.com/m
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

pkg :=user

all: $(pkg)
user: $@

$(pkg):
	@protoc -I$(PROTO_DIR) \
    		--go_out=. --go_opt=module=$(PACKAGE) \
    		--go-grpc_out=. --go-grpc_opt=module=$(PACKAGE) \
    		$(PROTO_DIR)/$@/*.proto

help:
	@${HELP_CMD}
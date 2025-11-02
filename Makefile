PROTO_DIR = proto
PACKAGE = github.com/relaunch-cot/lib-relaunch-cot
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

pkg :=user base_models chat project notification post

all: $(pkg)
user: $@
base_models: $@
chat: $@
project: $@
notification: $@
post: $@

$(pkg):
	@protoc -I$(PROTO_DIR) \
    		--go_out=. --go_opt=module=$(PACKAGE) \
    		--go-grpc_out=. --go-grpc_opt=module=$(PACKAGE) \
    		$(PROTO_DIR)/$@/*.proto

help:
	@${HELP_CMD}
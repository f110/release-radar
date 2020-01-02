.PHONY: update-deps
update-deps:
	bazel run //:gazelle -- update-repos -from_file=go.mod
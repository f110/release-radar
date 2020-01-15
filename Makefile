.PHONY: update
update:
	bazel run //:gazelle

.PHONY: update-deps
update-deps:
	go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%external_deps

build_rest:
	rm -f exercise_tracker_rest
	go build -o exercise_tracker_rest cmd/rest/main.go
build_cli:
	rm -f exercise_tracker_cli
	go build -o exercise_tracker_cli cmd/cli/main.go
run_rest:
	go run cmd/rest/main.go -env=$(env)
run_cli:
	go run cmd/cli/main.go -env=$(env)

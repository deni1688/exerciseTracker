build_rest:
	rm -f exercise_tracker_rest
	go build -o exercise_tracker_rest cmd/rest/main.go
build_cli:
	rm -f exercise_tracker_cli
	go build -o exercise_tracker_cli cmd/cli/main.go
run_rest:
	go run cmd/rest/main.go
run_consumer:
	go run cmd/rabbitmq/main.go
run_cli:
	go run cmd/cli/main.go

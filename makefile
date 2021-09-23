build_http:
	rm -f exercise_tracker_http
	go build -o exercise_tracker_http cmd/http/main.go
build_cli:
	rm -f exercise_tracker_cli
	go build -o exercise_tracker_cli cmd/cli/main.go
run_http:
	go run cmd/http/main.go
run_cli:
	go run cmd/cli/main.go

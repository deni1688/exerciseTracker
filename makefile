http:
	rm -f exercise_tracker_http
	go build -o exercise_tracker_http cmd/http/main.go
cli:
	rm -f exercise_tracker_cli
	go build -o exercise_tracker_cli cmd/cli/main.go

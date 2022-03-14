hello:
	echo "Hello"

run_server_a:
	cd ./server-a && go run main.go

run_server_b:
	cd ./server-b && go run main.go

run_debouncer:
	cd ./debouncer && go run main.go
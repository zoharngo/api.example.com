GPORT ?=3000

run:
	go run main.go -port=${GPORT}
build:
	make buid_db
buid_db:
	go run pkg/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db
	




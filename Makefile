# ==============================================================================
# Main
run: 
	go run cmd/server/main.go
migrate:
	migrate -database "mysql://root:tin14091998@tcp(0.0.0.0:3306)/hsr" -path migrations up
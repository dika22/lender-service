APP=lender-service
APP_EXECUTABLE=${APP}

serve-http:
	go run main.go serve-http

start-worker:
	go run main.go start-worker
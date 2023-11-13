install:
	cd server & go mod tidy

run_webapp:
	cd webapp && npm start webapp 

run_server:
	cd server && go run ./...  

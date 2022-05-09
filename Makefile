DB_HOST:="localhost"
DB_USER:="postgres"
DB_NAME:="db_url_shortener"

# DB ######
db-create:  ##@database create db
	createdb -h $(DB_HOST) -U $(DB_USER) -O$(DB_USER) -Eutf8 $(DB_NAME)

db-drop:  ##@database drop db
	dropdb -h $(DB_HOST) -U $(DB_USER) --if-exists $(DB_NAME)

# RUN ###
run:
	mkdir -p bin
	go build -o ./bin/app ./cmd
	bin/app

setup:
	go mod tidy
	go mod vendor
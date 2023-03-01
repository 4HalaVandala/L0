# L0
wildberries test task level 0

# ports
 ### Nats-streaming: <br>
 * localhost:4222 <br>
 * localhost:8222 <br>
 ### Postgresql: <br>
 * localhost:5433 <br>
 ### Http server <br>
 * localhost:8000 <br>
 
 # Quickstart
    docker-compose up -d &&
    migrate -path ./schema -database 'postgres://postgres:1122@localhost:5433/postgres?sslmode=disable' up &&
    go run cmd/main.go &&
    go run publisher/main.go

Script to publish json data (publisher/main.go) looking for "model.json" file in root directory then pubishing it to nats-channel.
Http server (cmd/main.go) listening messages on :4222 port and saving it in database. At the start server initializing cache and saving all rows from db in-memory.



 
 

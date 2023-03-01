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
    go run producer/main.go


 
 

# kalahgo
## running a mindless random agent
- in a terminal window, run `nc -l localhost 12345`
- run the mindless ranoom agent (for development purpose), `go run hello.go`
- run the game engine using `java -jar ManKalah.jar "nc localhost 12345" "java -jar MKRefAgent.jar`

## building the agent
- run `go build main.go`

## running the agent
- `java -jar ManKalah.jar "nc ./main" "java -jar MKRefAgent.jar`

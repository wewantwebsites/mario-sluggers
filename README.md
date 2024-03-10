# Mario Sluggers!
This is intended to an api and a frontend that provides player details and stats, chemistry, and a draft board. Mathematically there can only be a maximum of 4 teams per draft. 

## Running the app
```bash
go run server.go
```
that cmd will run the server. but in dev mode you probably want more...


## Hot Reload the Server
install `air` and it's CLI globally with 
```bash
go get -u github.com/cosmtrek/air
```
Now you can run the server with HMR with 
```bash
air server.go
```



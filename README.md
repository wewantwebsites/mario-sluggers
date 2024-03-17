# Mario Sluggers

This is intended to an api and a frontend that provides player details and stats, chemistry, and a draft board. Mathematically there can only be a maximum of 4 teams per draft.

## Running the app

```zsh
go run server.go
```

that cmd will run the server. but in dev mode you probably want more...

## Hot Reload the Server

install `air` and it's CLI globally with

```zsh
go get -u github.com/cosmtrek/air
```

Now you can run the server with HMR with

```zsh
air server.go
```

## Swagger

Generate swagger API documentation from the cmd line after udpating an endpoint

```zsh
swag init
```

# Frontier

This is an experimental game written in [Go](https://golang.org) with
[Ebiten](https://ebiten.org/). Fair warning: I have never made a game and I
don't know Go well.

## To build for macOS

```
go build
./frontier
```

## To build for Web

```
GOOS=js GOARCH=wasm go build -o frontier.wasm github.com/exastencil/frontier
python server.py
open http://localhost:8000/frontier.html
```

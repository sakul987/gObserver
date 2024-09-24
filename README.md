# gObserver
Server usage monitor using go and vue

## Build Dependencies
node.js 18+
npm
go
(lm-sensors)
(df)

## Run Dependencies
(lm-sensors)
(df)

## Install
```git clone https://github.com/sakul987/gObserver.git```
```cd gObserver/gObserver-ui```
```npm i```
Add certificate vite.crt and vite.key
```npm run build```
```cd ..```
```go build -o gObserver cmd/main.go```
```./gObserver```

# gopherjs + electron テスト

## prepare

```
npm -g install electron-prebuilt
npm init -y
go get -u github.com/gopherjs/gopherjs
```

## build, run

```
gopherjs build index.go
electron .
```

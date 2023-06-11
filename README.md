
### stoplight stdio

サンプル (みれますか？？)

https://keshiwo.stoplight.io/studio/my-api?source=0eimrf7v&symbol=%252Fp%252Freference%252Fjinshiho.yaml%252Fpaths%252F%7E1users%252Fpost

### goのライブラリ

https://github.com/deepmap/oapi-codegen/

```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

### APIの定義生成

```
oapi-codegen -generate gin,spec  -package generated jinshiho.yaml >  backend/server/generated/jinshiho.gen.go
```

### 型定義生成

```
oapi-codegen -generate types  -package generated jinshiho.yaml >  backend/server/generated/jinshiho.type.go
```
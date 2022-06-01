SET CGO_LDFLAGS=-Wl,--kill-at
SET CGO_LDFLAGS_ALLOW=-Wl,--kill-at
SET CGO_ENABLED=1

go build -ldflags="-s -w  -extldflags '-static'" -buildmode=c-shared -o huiguangbo.cat.dll
pause

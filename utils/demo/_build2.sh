CGO_LDFLAGS=-Wl,--kill-at CGO_LDFLAGS_ALLOW=-Wl,--kill-at CGO_ENABLED=1 CC=/usr/local/bin/i686-w64-mingw32-gcc GOOS=windows GOARCH=386 go build -ldflags="-s -w  -extldflags '-static'" -buildmode=exe -o demo.exe
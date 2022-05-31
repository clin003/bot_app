xx

	CGO_LDFLAGS=-Wl,--kill-at CGO_LDFLAGS_ALLOW=-Wl,--kill-at CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -ldflags="-s -w  -extldflags '-static'" -buildmode=c-shared -o demo.cat.dll

win7

	CGO_LDFLAGS=-Wl,--kill-at CGO_LDFLAGS_ALLOW=-Wl,--kill-at CGO_ENABLED=1 go build -ldflags="-s -w  -extldflags '-static'" -buildmode=c-shared -o demo.cat.dll
	
	
mac


	CGO_LDFLAGS=-Wl,--kill-at CGO_LDFLAGS_ALLOW=-Wl,--kill-at CGO_ENABLED=1 CC=/usr/local/bin/i686-w64-mingw32-gcc GOOS=windows GOARCH=386 go build -ldflags="-s -w  -extldflags '-static'" -buildmode=c-shared -o demo.cat.dll
Requires BearLibTerminal be installed 

Mac OS X: download BearLibTerminal here: 
- http://foo.wyrd.name/en:bearlibterminal/#download
- Create link in /usr/local/lib
- `ln libBearLibTerminal.dylib /usr/local/lib/libBearLibTerminal.dylib`

Windows DLL also in download link

Also requires `dep`:
https://github.com/golang/dep

Mac OS X:
```
$ brew install dep
$ brew upgrade dep
```

On other platforms you can use the install.sh script:

```
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
````

To set up: `make`

To clean up: `make clean`

Sources: 
- https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2
- https://gitlab.com/pantomath-io/demo-grpc
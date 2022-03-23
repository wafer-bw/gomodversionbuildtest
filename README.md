# gomodversionbuildtest
An example of how to run a go application such that the build info includes the module version via git tag

```sh
go install github.com/wafer-bw/gomodversionbuildtest@latest
# go: downloading github.com/wafer-bw/gomodversionbuildtest v0.4.0
go run github.com/wafer-bw/gomodversionbuildtest@latest
# 2022/03/23 12:02:45 :8998
# 2022/03/23 12:02:45 go  go1.18
# path    github.com/wafer-bw/gomodversionbuildtest
# mod     github.com/wafer-bw/gomodversionbuildtest       v0.4.0  h1:BFRPzyQfU6PJtNIlrpfNyYrNr2pmmDqyEzxol1s4Msw=
# dep     github.com/gorilla/mux  v1.8.0  h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=
# build   -compiler=gc
# build   CGO_ENABLED=1
# build   CGO_CFLAGS=
# build   CGO_CPPFLAGS=
# build   CGO_CXXFLAGS=
# build   CGO_LDFLAGS=
# build   GOARCH=amd64
# build   GOOS=darwin
# build   GOAMD64=v1
```

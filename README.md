# code coverage

go test -cover poms/...
go test -coverprofile=cover.out poms/...
go tool cover -func=cover.out
go tool cover -html cover.out

# hitmap

go test -covermode=count -coverprofile=cover.out poms/...
go tool cover -func=cover.out
go tool cover -html cover.out

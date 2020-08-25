$ mkdir hello # Alternatively, clone it if it already exists in version control.
$ cd hello
$ go mod init example.com/user/hello
go: creating new go.mod: module example.com/user/hello
$ cat go.mod
module example.com/user/hello

go 1.14
$



`
go run -gcflags="-B" outofbounds_main.go

go tool compile -S outofbounds_main.go
`

### Print out GOSSAFUNC for a method
`
set GOSSAFUNC=printList 
go run outofbounds_main.go
`

####This will generate a `ssa.html` file
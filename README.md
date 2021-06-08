```
.
|-- greetings
|   |-- go.mod
|   |-- greetings.go
|   `-- greetings_test.go
`-- hello
    |-- go.mod
    |-- go.sum
    `-- hello.go
```

# main branch
**Use "example.com/*" as the module path**
greetings/go.mod
```
module example.com/greetings

go 1.16
```
hello/go.mod
```
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require (
        example.com/greetings v0.0.0-00010101000000-000000000000
        rsc.io/quote v1.5.2
)
```


# project1 branch
**Use "github.com/XihuanYang/*" as the module path**
greetings/go.mod
```
module github.com/XihuanYang/greetings

go 1.16
```
hello/go.mod
```
module github.com/XihuanYang/hello

go 1.16

replace github.com/XihuanYang/greetings => ../greetings

require (
        github.com/XihuanYang/greetings v0.0.0-00010101000000-000000000000
        rsc.io/quote v1.5.2
)
```

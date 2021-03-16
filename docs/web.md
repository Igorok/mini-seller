## Web application with golang

### GVM
gvm - golang version manager, it's util to install and use different versions of golang.
```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

gvm install go1.15.6 -B
gvm use go1.15.6 [--default]
gvm list
```

### Go modules
A module is a collection of Go packages stored in a file tree with a go.mod file at its root.
```
go mod init mini-seller
```

### Viper
Viper is package for manage configuration
```
go get github.com/spf13/viper
```

### Validation
validator.v9 - package very helpful for validation of structures
```
go get gopkg.in/go-playground/validator.v9
```

### Mongo
Driver for mongodb
```
go get go.mongodb.org/mongo-driver
```

### Testing
```
go test -v ./...
```

### Gin
Gin is a web framework written in Go
```
go get -u github.com/gin-gonic/gin
```

### GraphQL
GraphQL is kind of webapi, it provide a query language, understandable description of the data in your API and has very powerful playground.
```
go get github.com/graphql-go/graphql
```

Query
```
query getCatalog($skip: Int, $limit: Int) {
  getCatalog(skip: $skip, limit: $limit) {
    products {
      id
      name
      price
      count
      Organization {
        id
        name
      }
      Category {
        id
        name
      }
    }
    count
  }
}
```

Variables
```
{
    "skip": 2,
    "limit": 2
}
```






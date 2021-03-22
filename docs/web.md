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
GraphQL is kind of web api, it provide a query language, understandable documentation for your requests and has pretty good playground.
```
go get github.com/graphql-go/graphql
```

With GraphQL you could select fields that you need from backend.

Query
```
query getProductList($skip: Int, $limit: Int) {
  getProductList(skip: $skip, limit: $limit) {
    products {
      id
      name
      price
      count
      Category {
        id
        name
      }
      Organization {
        id
        email
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

You could use Aliases and get different data in one request. You could use Fragment for common fields.

Query
```
query getProductDetail($id_cola: String, $id_salad: String) {
    cola: getProductDetail(id: $id_cola) {
        ...detailFields
    },
    salad: getProductDetail(id: $id_salad) {
        ...detailFields
    }
}

fragment detailFields on ProductForList {
    id
    name
    price
    count
    Category {
        id
        name
    }
    Organization {
        id
        email
    }
} }
}
```

Variables
```
{
    "id_cola": "604497558ffcad558eb8e1f5",
    "id_salad": "604497558ffcad558eb8e1f4"
}
```
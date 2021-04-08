## Big text about small golang application for beginners
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software etc. I want to build web api with golang, because big part of my work is backend of web or mobile applications.


### GVM
First step for work with golang is install golang.
Gvm - golang version manager, it's util to install and use different versions of golang.
```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

gvm install go1.15.6 -B
gvm use go1.15.6 [--default]
gvm list
```

### Go modules
To create something workable with any language you should have package manager to install libraries from other developers and organizations. Golang have thing called Go Modules - a module is a collection of Go packages stored in a file tree with a go.mod file at its root.
```
go mod init mini-seller
```

### Viper
When you develop any web product you have some configuration for web server, databases, file storages, api for integrations. Viper is comfortable and powerful package for manage configuration.
```
go get github.com/spf13/viper
```

### Mongo
To save data for my project i will use Mongo it is nosql database with simple syntax, big possibilities and good performance.
```
go get go.mongodb.org/mongo-driver
```

### Testing
When you develop something bigger than landing page you need to be sure that you project is working. Launch tests in golang is very simple.
```
go test -v ./...
```

### Gqlgen
And one of important things for web application is itself web application. I select gqlgen for this, it have generation of code and big possibilities to make graphql api.
```
go get github.com/99designs/gqlgen
```

Initialize web application
```
go run github.com/99designs/gqlgen init
```

Generate code for new model
```
go run github.com/99designs/gqlgen generate
```

At the top of our resolver.go, between package and import, add the following line:
```
//go:generate go run github.com/99designs/gqlgen
```
To run go generate recursively over your entire project, use this command:
```
go generate ./...
```

### Desing
After installation of packages need to plane design of application. Building a structure of application is not simple and before do it very helpful to read about solid rules, clean architecture, n-tier architecture. Very shortly central part of application is entities of business logic. Business logic of project should depended from entities. But logic should not be related with web frameworks or data storages like orm or web api. Need to get data from data storage from classes of repository, because of this will simples change one database to another or message broker or web api. And need to test logic uses mocks of repositories. If you have big difficult validation for your logic you should put this in classes of specifications. And you should not relate your logic with web framework, instead of this web application should depended from use cases, this make changing of web framework more simple.

In my case i made folders:

1. application - code for web application
    1. server - web server for project
    2. graph - folder with gqlgen application
        1. schemas - folder contain schemas for graphql
        2. model - folder contain entities for graphql
        3. resolvers - folder contain controllers for graphql

2. domain - business logic of project
    1. common - common data for all packages
        1. entities - here i save all entities, these could be entities for business logic, models for database and dto for communication between classes
    2. packages - here logic of project
        1. catalogpkg - logic for catalog of products
            1. usecase - interface for use case, it describe functionality available in package
            2. repository - interface describe functionality of data storage
            3. catalogusecase - contain use case - business logic of catalog and test for use case
            4. catalogrepository - contain repository with requests for database, integration tests for validation of repository, and mock of repository, it needed to test logic of use case without real connection to database
3. infrastructure - code without business logic, database driver, helper for configuration
    1. mongohelper - contain functionality for configure connection to mongo database and folder with test data. These files need to initialize empty project with demo data, and i use these for testing of logic.
    2. viperhelper - i use this helper to read configuration for project and for test. This helps me to change configuration from default to local and use environment variables.



## GraphQL
GraphQL is kind of web api, it provide a query language, understandable documentation for your requests and has pretty good playground.

With GraphQL you could select fields that you need from backend and you could create resolvers for related entities. Example for product:

```
query product($id: String!) {
    product(id: $id) {
        id
        name
        price
        count
        category{
            id
            name
        }
        organization {
            id
            name
        }
    }
}
```

Graphql have powerful query language, you could use Aliases and get different data in one request and you could use Fragment for common fields.
Query
```
query products($id_cola: String! $id_salad: String!) {
    cola: product(id: $id_cola) {
        ...detailFields
    }
  	salad: product(id: $id_salad) {
        ...detailFields
    }
}

fragment detailFields on Product {
    id
    name
    price
    count
    category{
        id
        name
    }
    organization {
        id
        name
    }
}
```
Variables
```
{
    "id_cola": "604497558ffcad558eb8e1f5",
    "id_salad": "604497558ffcad558eb8e1f4"
}
```

Result
```
{
    "data": {
        "cola": {
            "id": "604497558ffcad558eb8e1f5",
            "name": "Cola",
            "price": 100,
            "count": 100,
            "category": {
                "id": "604488100f719d9c76a28fe3",
                "name": "Drinks"
            },
            "organization": {
                "id": "6043d76e94df8de741c2c0d6",
                "name": "restaurant"
            }
        },
        "salad": {
            "id": "604497558ffcad558eb8e1f4",
            "name": "Salad Cesar",
            "price": 200,
            "count": 10,
            "category": {
                "id": "604488100f719d9c76a28fe6",
                "name": "Salad"
            },
            "organization": {
                "id": "6043d76e94df8de741c2c0d6",
                "name": "restaurant"
            }
        }
    }
}
```





### Validation
validator.v9 - package very helpful for validation of structures
```
go get gopkg.in/go-playground/validator.v9
```
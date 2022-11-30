# Driver

[[_TOC_]]


## Mongosh
```sh
$> mongosh "mongodb+srv://dhcluster001.c17aj.mongodb.net/myFirstDatabase" --username dhub
```



## Go 
### V.16
```go
mongodb+srv://dhub:<password>@dhcluster001.c17aj.mongodb.net/myFirstDatabase?retryWrites=true&w=majority
```

#### Example with driver
```go
import "go.mongodb.org/mongo-driver/mongo"
clientOptions := options.Client().
    ApplyURI("mongodb+srv://dhub:<password>@dhcluster001.c17aj.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}
```
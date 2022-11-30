# Development

## DarsHub-App-DEV

| Name            | Deployments | Users                 |
| --------------- | ----------- | --------------------- |
| DarsHub-App-DEV | 1           | 5 (Alle aus dem Team) |


### Deployed Databases

| Name         | Version | Region                           | CLUSTER TIER         | TYPE                  | BACKUPS  | Access                     |
| ------------ | ------- | -------------------------------- | -------------------- | --------------------- | -------- | -------------------------- |
| dhCluster001 | 5.0.14  | Azure / Netherlands (westeurope) | M0 Sandbox (General) | Replica Set - 3 nodes | Inactive | Allow Access from Anywhere |








### Connected with Mongosh
```sh
$> mongosh "mongodb+srv://dhcluster001.c17aj.mongodb.net/myFirstDatabase" --username dhub
```

### Go V.16
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


### DB User Access

| Name        | Password         | Auth Type | Resources     | ROLE                   |
| ----------- | ---------------- | --------- | ------------- | ---------------------- |
| dhub        | OiPe7pU8kxaIVhBx | SCRAM     | All Resources |                        |
| reports     | 9zVjpzk3EuORJhnO | SCRAM     | Read only     |                        |
| db_readOnly | z4SaweOAgLMvR4bg | SCRAM     | Read only     | Only read any database |



### Network Access

| IP-Address                                    | Comment           | Status |
| --------------------------------------------- | ----------------- | ------ |
| 0.0.0.0/0  (includes your current IP address) | Any IP can access | ACTIVE |

### Services
- NO ACTIVE TRIGGERS
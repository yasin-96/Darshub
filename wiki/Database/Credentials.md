# DarsHub-App-DEV

[[_TOC_]]

## DB User Access

| Name        | Password         | Auth Type | Resources     | ROLE                   |
| ----------- | ---------------- | --------- | ------------- | ---------------------- |
| dhub        | OiPe7pU8kxaIVhBx | SCRAM     | All Resources |                        |
| reports     | 9zVjpzk3EuORJhnO | SCRAM     | Read only     |                        |
| db_readOnly | z4SaweOAgLMvR4bg | SCRAM     | Read only     | Only read any database |

---
# DarsHub-App-Prod
## DB User Access

| Name        | Password         | Auth Type | Resources     | ROLE                   |
| ----------- | ---------------- | --------- | ------------- | ---------------------- |
| dhub        | NPfzTFefkrDb9B5R | SCRAM     | All Resources |                        |
| reports     | Ar1whW93Cwsx47Ge | SCRAM     | Read only     |                        |
| db_readOnly | ZqJ5Q6vhUTjp6nBS | SCRAM     | Read only     | Only read any database |



### Deployed Databases

| Name         | Version | Region                           | CLUSTER TIER         | TYPE                  | BACKUPS  | Access                     |
| ------------ | ------- | -------------------------------- | -------------------- | --------------------- | -------- | -------------------------- |
| dhCluster001 | 5.0.14  | Azure / Netherlands (westeurope) | M0 Sandbox (General) | Replica Set - 3 nodes | Inactive | Allow Access from Anywhere |


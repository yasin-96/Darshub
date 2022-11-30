# Production

[[_TOC_]]

## DarsHub-DB-PROD

| Name            | Deployments | Users                 |
| --------------- | ----------- | --------------------- |
| DarsHub-DB-PROD | 1           | 5 (Alle aus dem Team) |


### Deployed Databases

| Name         | Version | Region                           | CLUSTER TIER         | TYPE                  | BACKUPS  | Access                     |
| ------------ | ------- | -------------------------------- | -------------------- | --------------------- | -------- | -------------------------- |
| dhCluster001 | 5.0.13  | Azure / Frankfurt (eu-central-1) | M0 Sandbox (General) | Replica Set - 3 nodes | Inactive | Allow Access from Anywhere |



### DB User Access

| Name        | Password         | Auth Type | Resources     | ROLE                   |
| ----------- | ---------------- | --------- | ------------- | ---------------------- |
| dhub        | NPfzTFefkrDb9B5R | SCRAM     | All Resources |                        |
| reports     | Ar1whW93Cwsx47Ge | SCRAM     | Read only     |                        |
| db_readOnly | ZqJ5Q6vhUTjp6nBS | SCRAM     | Read only     | Only read any database |


### Network Access

| IP-Address                                    | Comment           | Status |
| --------------------------------------------- | ----------------- | ------ |
| 0.0.0.0/0  (includes your current IP address) | Any IP can access | ACTIVE |

### Services
- NO ACTIVE TRIGGERS
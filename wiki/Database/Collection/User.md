# User
| Propertyname | Description | Type |
| ------------ | ----------- | ---- |
| _id          |             |      |
| password     |             |      |
| first_name   |             |      |
| last_name    |             |      |
| birthday     |             |      |
| avatar       |             |      |
| email        |             |      |
| telNr        |             |      |
| company      |             |      |
| occupation   |             |      |
| school       |             |      |
| subject      |             |      |
| country      |             |      |
| isActive     |             |      |
| role         |             |      |

## Example 

```json
{
  "_id": {
    "$oid": "624855424b4c13366a80a362"
  },
  "password": {
    "$binary": {
      "base64": "JDJhJDEwJEcub29BNVlmVWFCM2pHSmVZb2lsbi52VDBnUUFJN3dWdExmdkZZSzBGOUlPN1EuamV6aXFx",
      "subType": "00"
    }
  },
  "first_name": "Alexander",
  "last_name": "Jajzler",
  "birthday": {
    "$date": "2022-02-02T00:00:00Z"
  },
  "avatar": "",
  "email": "asd",
  "telNr": "323",
  "company": "",
  "occupation": "",
  "school": "",
  "subject": "",
  "country": "Andorra",
  "isActive": false,
  "role": null
}
```
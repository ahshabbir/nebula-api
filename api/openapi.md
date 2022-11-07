To create openapi specification document:

```
    // cd ./api
    npx swagger-inline "./*/*.go" --base ./openapiBase.json > openapi.json
```
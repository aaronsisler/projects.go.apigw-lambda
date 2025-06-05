# projects.go.apigw-lambda

## TODO

- How do you create two handler functions in the same project
- How can I zip them up independently or do I need to make them one uber handler that routes internally

## Deployment Steps

```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap src/main.go
```

```bash
zip go-apigw-lambda.zip bootstrap
```

```bash
aws lambda create-function --function-name myFunction \
--runtime provided.al2023 --handler bootstrap \
--architectures arm64 \
--role arn:aws:iam::654918520080:role/service-role/go-apigw-lambda-role-28cwtfeh \
--zip-file fileb://go-apigw-lambda.zip
```

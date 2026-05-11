# RestAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationOauthRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```

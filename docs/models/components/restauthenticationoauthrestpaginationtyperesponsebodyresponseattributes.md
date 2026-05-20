# RestAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationOauthRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

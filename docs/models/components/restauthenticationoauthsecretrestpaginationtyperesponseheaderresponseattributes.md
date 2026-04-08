# RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationOauthSecretRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```

# RestAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationOauthSecretRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

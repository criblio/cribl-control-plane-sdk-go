# RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationBasicSecretRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```

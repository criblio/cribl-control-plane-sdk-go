# RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributes := components.CreateRestAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributes.Type {
	case components.RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributesTypeArrayOfAny:
		// restAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributesTypeStr:
		// restAuthenticationLoginSecretRestPaginationTypeResponseHeaderResponseAttributes.Str is populated
}
```

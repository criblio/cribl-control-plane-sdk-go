# RestAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationLoginSecretRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

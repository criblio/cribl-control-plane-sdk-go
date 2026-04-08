# RestAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationBasicSecretRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

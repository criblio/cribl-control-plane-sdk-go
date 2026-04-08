# RestAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationBasicRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

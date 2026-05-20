# RestAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restAuthenticationLoginRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

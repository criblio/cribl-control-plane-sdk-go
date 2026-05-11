# RestCollectMethodGetRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodGetRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodGetRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodGetRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodGetRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestCollectMethodGetRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restCollectMethodGetRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodGetRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restCollectMethodGetRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

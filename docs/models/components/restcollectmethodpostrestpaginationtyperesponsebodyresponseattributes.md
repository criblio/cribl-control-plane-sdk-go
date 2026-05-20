# RestCollectMethodPostRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodPostRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodPostRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodPostRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodPostRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestCollectMethodPostRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restCollectMethodPostRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodPostRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restCollectMethodPostRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

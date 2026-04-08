# RestCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restCollectMethodPostWithBodyRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

# RestCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributes

Names of attributes within the response that contain next-page information


## Supported Types

### 

```go
restCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributesArrayOfAny([]any{/* values here */})
```

### 

```go
restCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributes := components.CreateRestCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributesStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributes.Type {
	case components.RestCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributesTypeArrayOfAny:
		// restCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributes.ArrayOfAny is populated
	case components.RestCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributesTypeStr:
		// restCollectMethodOtherRestPaginationTypeResponseBodyResponseAttributes.Str is populated
}
```

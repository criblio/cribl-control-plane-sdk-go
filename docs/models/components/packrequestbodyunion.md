# PackRequestBodyUnion


## Supported Types

### PackRequestBody1

```go
packRequestBodyUnion := components.CreatePackRequestBodyUnionPackRequestBody1(components.PackRequestBody1{/* values here */})
```

### PackRequestBody2

```go
packRequestBodyUnion := components.CreatePackRequestBodyUnionPackRequestBody2(components.PackRequestBody2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch packRequestBodyUnion.Type {
	case components.PackRequestBodyUnionTypePackRequestBody1:
		// packRequestBodyUnion.PackRequestBody1 is populated
	case components.PackRequestBodyUnionTypePackRequestBody2:
		// packRequestBodyUnion.PackRequestBody2 is populated
}
```

# GetProductsGroupsByProductResponseBody

List of ConfigGroup objects.


## Supported Types

### CountedConfigGroup

```go
getProductsGroupsByProductResponseBody := operations.CreateGetProductsGroupsByProductResponseBodyCountedConfigGroup(components.CountedConfigGroup{/* values here */})
```

### PaginatedConfigGroup

```go
getProductsGroupsByProductResponseBody := operations.CreateGetProductsGroupsByProductResponseBodyPaginatedConfigGroup(components.PaginatedConfigGroup{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getProductsGroupsByProductResponseBody.Type {
	case operations.GetProductsGroupsByProductResponseBodyTypeCountedConfigGroup:
		// getProductsGroupsByProductResponseBody.CountedConfigGroup is populated
	case operations.GetProductsGroupsByProductResponseBodyTypePaginatedConfigGroup:
		// getProductsGroupsByProductResponseBody.PaginatedConfigGroup is populated
}
```

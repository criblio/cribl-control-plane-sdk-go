# LabelFieldsUnion


## Supported Types

### LabelFields1

```go
labelFieldsUnion := components.CreateLabelFieldsUnionLabelFields1(components.LabelFields1{/* values here */})
```

### LabelFields2

```go
labelFieldsUnion := components.CreateLabelFieldsUnionLabelFields2(components.LabelFields2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch labelFieldsUnion.Type {
	case components.LabelFieldsUnionTypeLabelFields1:
		// labelFieldsUnion.LabelFields1 is populated
	case components.LabelFieldsUnionTypeLabelFields2:
		// labelFieldsUnion.LabelFields2 is populated
}
```

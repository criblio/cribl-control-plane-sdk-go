# ChecksumBefore


## Supported Types

### 

```go
checksumBefore := components.CreateChecksumBeforeStr(string{/* values here */})
```

### 

```go
checksumBefore := components.CreateChecksumBeforeArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checksumBefore.Type {
	case components.ChecksumBeforeTypeStr:
		// checksumBefore.Str is populated
	case components.ChecksumBeforeTypeArrayOfStr:
		// checksumBefore.ArrayOfStr is populated
}
```

# GitDiffLines


## Supported Types

### DiffLineDelete

```go
gitDiffLines := components.CreateGitDiffLinesDelete(components.DiffLineDelete{/* values here */})
```

### DiffLineInsert

```go
gitDiffLines := components.CreateGitDiffLinesInsert(components.DiffLineInsert{/* values here */})
```

### DiffLineContext

```go
gitDiffLines := components.CreateGitDiffLinesContext(components.DiffLineContext{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch gitDiffLines.Type {
	case components.GitDiffLinesTypeDelete:
		// gitDiffLines.DiffLineDelete is populated
	case components.GitDiffLinesTypeInsert:
		// gitDiffLines.DiffLineInsert is populated
	case components.GitDiffLinesTypeContext:
		// gitDiffLines.DiffLineContext is populated
	default:
		// Unknown type - use gitDiffLines.GetUnknownRaw() for raw JSON
}
```

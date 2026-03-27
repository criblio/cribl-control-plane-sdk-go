# TypeOptions

Parser or formatter type to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TypeOptionsCsv

// Open enum: custom values can be created with a direct type cast
custom := components.TypeOptions("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `TypeOptionsCsv`   | csv                |
| `TypeOptionsElff`  | elff               |
| `TypeOptionsClf`   | clf                |
| `TypeOptionsKvp`   | kvp                |
| `TypeOptionsJSON`  | json               |
| `TypeOptionsDelim` | delim              |
| `TypeOptionsRegex` | regex              |
| `TypeOptionsGrok`  | grok               |
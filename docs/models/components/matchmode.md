# MatchMode

Specifies the matching method based on the format and logic used in the lookup file

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MatchModeExact

// Open enum: custom values can be created with a direct type cast
custom := components.MatchMode("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `MatchModeExact` | exact            |
| `MatchModeCidr`  | cidr             |
| `MatchModeRegex` | regex            |
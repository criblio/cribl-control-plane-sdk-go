# LookupDbLookupFalseMatchModeExactMatchMode

Specifies the matching method based on the format and logic used in the lookup file

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LookupDbLookupFalseMatchModeExactMatchModeExact

// Open enum: custom values can be created with a direct type cast
custom := components.LookupDbLookupFalseMatchModeExactMatchMode("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `LookupDbLookupFalseMatchModeExactMatchModeExact` | exact                                             |
| `LookupDbLookupFalseMatchModeExactMatchModeCidr`  | cidr                                              |
| `LookupDbLookupFalseMatchModeExactMatchModeRegex` | regex                                             |
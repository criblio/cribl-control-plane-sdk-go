# LookupDbLookupFalseMatchModeCidrMatchMode

Specifies the matching method based on the format and logic used in the lookup file

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LookupDbLookupFalseMatchModeCidrMatchModeExact

// Open enum: custom values can be created with a direct type cast
custom := components.LookupDbLookupFalseMatchModeCidrMatchMode("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `LookupDbLookupFalseMatchModeCidrMatchModeExact` | exact                                            |
| `LookupDbLookupFalseMatchModeCidrMatchModeCidr`  | cidr                                             |
| `LookupDbLookupFalseMatchModeCidrMatchModeRegex` | regex                                            |
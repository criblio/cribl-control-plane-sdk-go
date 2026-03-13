# MatchTypeOptionsLookupDbLookupFalseMatchModeCidr

Further defines how to handle multiple matches: return the first match, the most specific match, or all matches

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MatchTypeOptionsLookupDbLookupFalseMatchModeCidrFirst

// Open enum: custom values can be created with a direct type cast
custom := components.MatchTypeOptionsLookupDbLookupFalseMatchModeCidr("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `MatchTypeOptionsLookupDbLookupFalseMatchModeCidrFirst`    | first                                                      |
| `MatchTypeOptionsLookupDbLookupFalseMatchModeCidrSpecific` | specific                                                   |
| `MatchTypeOptionsLookupDbLookupFalseMatchModeCidrAll`      | all                                                        |
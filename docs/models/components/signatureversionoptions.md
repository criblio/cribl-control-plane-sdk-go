# SignatureVersionOptions

Signature version to use for signing MSK cluster requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SignatureVersionOptionsV2

// Open enum: custom values can be created with a direct type cast
custom := components.SignatureVersionOptions("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `SignatureVersionOptionsV2` | v2                          |
| `SignatureVersionOptionsV4` | v4                          |
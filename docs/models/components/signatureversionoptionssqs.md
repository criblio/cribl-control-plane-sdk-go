# SignatureVersionOptionsSqs

Signature version to use for signing SQS requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SignatureVersionOptionsSqsV2

// Open enum: custom values can be created with a direct type cast
custom := components.SignatureVersionOptionsSqs("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `SignatureVersionOptionsSqsV2` | v2                             |
| `SignatureVersionOptionsSqsV4` | v4                             |
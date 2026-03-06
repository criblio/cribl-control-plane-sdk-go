# SignatureVersionOptions1

Signature version to use for signing EC2 requests

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SignatureVersionOptions1V2

// Open enum: custom values can be created with a direct type cast
custom := components.SignatureVersionOptions1("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `SignatureVersionOptions1V2` | v2                           |
| `SignatureVersionOptions1V4` | v4                           |
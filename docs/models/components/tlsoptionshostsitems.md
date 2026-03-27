# TLSOptionsHostsItems

Whether to inherit TLS configs from group setting or disable TLS

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TLSOptionsHostsItemsInherit

// Open enum: custom values can be created with a direct type cast
custom := components.TLSOptionsHostsItems("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `TLSOptionsHostsItemsInherit` | inherit                       |
| `TLSOptionsHostsItemsOff`     | off                           |
# DiskSpaceProtectionOptions

How to handle events when disk space is below the global 'Min free disk space' limit

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DiskSpaceProtectionOptionsBlock

// Open enum: custom values can be created with a direct type cast
custom := components.DiskSpaceProtectionOptions("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `DiskSpaceProtectionOptionsBlock` | block                             |
| `DiskSpaceProtectionOptionsDrop`  | drop                              |
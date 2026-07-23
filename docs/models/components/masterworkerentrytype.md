# MasterWorkerEntryType

RPC message type reported by the node.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.MasterWorkerEntryTypeInfo

// Open enum: custom values can be created with a direct type cast
custom := components.MasterWorkerEntryType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `MasterWorkerEntryTypeInfo` | info                        |
| `MasterWorkerEntryTypeReq`  | req                         |
| `MasterWorkerEntryTypeResp` | resp                        |
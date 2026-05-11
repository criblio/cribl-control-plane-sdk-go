# CreateInputSystemByPackCheckpointStore

The backing store used to persist consumer checkpoints. Select "None" to disable checkpointing (consumers will restart from the configured start position).

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackCheckpointStoreNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackCheckpointStore("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CreateInputSystemByPackCheckpointStoreNone`      | none                                              |
| `CreateInputSystemByPackCheckpointStoreAzureBlob` | azureBlob                                         |
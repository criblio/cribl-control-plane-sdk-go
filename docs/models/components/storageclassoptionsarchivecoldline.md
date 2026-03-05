# StorageClassOptionsArchiveColdline

Storage class to select for uploaded objects

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.StorageClassOptionsArchiveColdlineStandard

// Open enum: custom values can be created with a direct type cast
custom := components.StorageClassOptionsArchiveColdline("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `StorageClassOptionsArchiveColdlineStandard` | STANDARD                                     |
| `StorageClassOptionsArchiveColdlineNearline` | NEARLINE                                     |
| `StorageClassOptionsArchiveColdlineColdline` | COLDLINE                                     |
| `StorageClassOptionsArchiveColdlineArchive`  | ARCHIVE                                      |
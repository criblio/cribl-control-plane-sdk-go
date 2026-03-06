# StorageClassOptions1

Storage class to select for uploaded objects

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.StorageClassOptions1Standard

// Open enum: custom values can be created with a direct type cast
custom := components.StorageClassOptions1("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `StorageClassOptions1Standard` | STANDARD                       |
| `StorageClassOptions1Nearline` | NEARLINE                       |
| `StorageClassOptions1Coldline` | COLDLINE                       |
| `StorageClassOptions1Archive`  | ARCHIVE                        |
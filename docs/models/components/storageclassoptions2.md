# StorageClassOptions2

Storage class to select for uploaded objects

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.StorageClassOptions2Standard

// Open enum: custom values can be created with a direct type cast
custom := components.StorageClassOptions2("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `StorageClassOptions2Standard`          | STANDARD                                |
| `StorageClassOptions2ReducedRedundancy` | REDUCED_REDUNDANCY                      |
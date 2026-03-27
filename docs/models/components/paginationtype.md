# PaginationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.PaginationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.PaginationType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `PaginationTypeNone`               | none                               |
| `PaginationTypeResponseBody`       | response_body                      |
| `PaginationTypeResponseHeader`     | response_header                    |
| `PaginationTypeResponseHeaderLink` | response_header_link               |
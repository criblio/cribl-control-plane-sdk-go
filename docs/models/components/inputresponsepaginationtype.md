# InputResponsePaginationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponsePaginationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponsePaginationType("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `InputResponsePaginationTypeNone`               | none                                            |
| `InputResponsePaginationTypeResponseBody`       | response_body                                   |
| `InputResponsePaginationTypeResponseHeader`     | response_header                                 |
| `InputResponsePaginationTypeResponseHeaderLink` | response_header_link                            |
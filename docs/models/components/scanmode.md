# ScanMode

Acceleration scan mode. <code>quick</code> collects object-level metadata; <code>detailed</code> also collects field-level statistics.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ScanModeDetailed

// Open enum: custom values can be created with a direct type cast
custom := components.ScanMode("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `ScanModeDetailed` | detailed           |
| `ScanModeQuick`    | quick              |
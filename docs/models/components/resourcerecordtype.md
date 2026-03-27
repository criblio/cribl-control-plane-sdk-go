# ResourceRecordType

The DNS record type (RR) to return. Defaults to 'A'.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ResourceRecordTypeA

// Open enum: custom values can be created with a direct type cast
custom := components.ResourceRecordType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ResourceRecordTypeA`     | A                         |
| `ResourceRecordTypeAaaa`  | AAAA                      |
| `ResourceRecordTypeAny`   | ANY                       |
| `ResourceRecordTypeCname` | CNAME                     |
| `ResourceRecordTypeMx`    | MX                        |
| `ResourceRecordTypeNaptr` | NAPTR                     |
| `ResourceRecordTypeNs`    | NS                        |
| `ResourceRecordTypePtr`   | PTR                       |
| `ResourceRecordTypeSoa`   | SOA                       |
| `ResourceRecordTypeSrv`   | SRV                       |
| `ResourceRecordTypeTxt`   | TXT                       |
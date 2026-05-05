# OutputResponseOutputAzureDataExplorerAuthenticationMethod

The type of OAuth 2.0 client credentials grant flow to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputAzureDataExplorerAuthenticationMethodClientSecret

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputAzureDataExplorerAuthenticationMethod("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `OutputResponseOutputAzureDataExplorerAuthenticationMethodClientSecret`     | clientSecret                                                                |
| `OutputResponseOutputAzureDataExplorerAuthenticationMethodClientTextSecret` | clientTextSecret                                                            |
| `OutputResponseOutputAzureDataExplorerAuthenticationMethodCertificate`      | certificate                                                                 |
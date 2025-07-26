# PreprocessCrowdstrike


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Disabled`                                                                   | **bool*                                                                      | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Command`                                                                    | **string*                                                                    | :heavy_minus_sign:                                                           | Command to feed the data through (via stdin) and process its output (stdout) |
| `Args`                                                                       | []*string*                                                                   | :heavy_minus_sign:                                                           | Arguments to be added to the custom command                                  |
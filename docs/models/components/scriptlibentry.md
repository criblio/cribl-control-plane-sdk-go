# ScriptLibEntry


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ID`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Command`                                                | *string*                                                 | :heavy_check_mark:                                       | Command to execute for this script                       |
| `Description`                                            | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `Args`                                                   | []*string*                                               | :heavy_minus_sign:                                       | Arguments to pass when executing this script             |
| `Env`                                                    | map[string]*string*                                      | :heavy_minus_sign:                                       | Extra environment variables to set when executing script |
| `AdditionalProperties`                                   | map[string]*any*                                         | :heavy_minus_sign:                                       | N/A                                                      |
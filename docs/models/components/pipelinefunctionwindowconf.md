# PipelineFunctionWindowConf


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `EventWindowID`                                                 | *float64*                                                       | :heavy_check_mark:                                              | Identifies the unique ID, used for a event window               |
| `RegisteredFunctions`                                           | []*string*                                                      | :heavy_check_mark:                                              | All window functions, tracked by this event window              |
| `TailEventCount`                                                | **float64*                                                      | :heavy_minus_sign:                                              | Number of events to keep before the current event in the window |
| `HeadEventCount`                                                | **float64*                                                      | :heavy_minus_sign:                                              | Number of events to keep after the current event in the window  |
# CreateOutputSystemTestByPackAndIDRequest


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ID`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The <code>id</code> of the Destination to send sample event data to.         |
| `Pack`                                                                       | *string*                                                                     | :heavy_check_mark:                                                           | The <code>id</code> of the Pack to create.                                   |
| `OutputTestRequest`                                                          | [components.OutputTestRequest](../../models/components/outputtestrequest.md) | :heavy_check_mark:                                                           | OutputTestRequest object                                                     |
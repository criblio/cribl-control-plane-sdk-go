# RestartResponse

Result of a restart request for a Worker or Edge Node.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ID`                                                                                         | `string`                                                                                     | :heavy_check_mark:                                                                           | Unique identifier for the Worker or Edge Node (GUID).                                        |
| `Message`                                                                                    | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Error message if the restart request failed for this Node.                                   |
| `Status`                                                                                     | [components.RestartResponseStatus](../../models/components/restartresponsestatus.md)         | :heavy_check_mark:                                                                           | Result of the restart request for this Node (<code>Restarting</code> or <code>Error</code>). |
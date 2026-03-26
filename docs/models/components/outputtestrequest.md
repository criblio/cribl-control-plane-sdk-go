# OutputTestRequest

Request body for testing a Destination by sending sample events.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Events`                                                       | []map[string]`any`                                             | :heavy_check_mark:                                             | Array of event objects to send to the Destination for testing. |
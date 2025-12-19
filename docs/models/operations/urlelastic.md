# URLElastic


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `URL`                                                                             | *string*                                                                          | :heavy_check_mark:                                                                | The URL to an Elastic node to send events to. Example: http://elastic:9200/_bulk  |
| `Weight`                                                                          | **float64*                                                                        | :heavy_minus_sign:                                                                | Assign a weight (>0) to each endpoint to indicate its traffic-handling capability |
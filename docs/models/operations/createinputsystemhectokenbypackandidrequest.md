# CreateInputSystemHecTokenByPackAndIDRequest


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ID`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Splunk HEC Source.                                  |
| `Pack`                                                                         | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Pack to create.                                     |
| `AddHecTokenRequest`                                                           | [components.AddHecTokenRequest](../../models/components/addhectokenrequest.md) | :heavy_check_mark:                                                             | AddHecTokenRequest object                                                      |
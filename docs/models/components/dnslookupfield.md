# DNSLookupField


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `InFieldName`                                                                      | **string*                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `ResourceRecordType`                                                               | [*components.ResourceRecordType](../../models/components/resourcerecordtype.md)    | :heavy_minus_sign:                                                                 | The DNS record type (RR) to return. Defaults to 'A'.                               |
| `OutFieldName`                                                                     | **string*                                                                          | :heavy_minus_sign:                                                                 | Name of field to add lookup results to. Leave blank to overwrite the lookup field. |
# UpdatePacksByIDRequest


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ID`                                                        | *string*                                                    | :heavy_check_mark:                                          | Pack name                                                   |
| `Source`                                                    | **string*                                                   | :heavy_minus_sign:                                          | body string required Pack source                            |
| `Minor`                                                     | **string*                                                   | :heavy_minus_sign:                                          | body boolean optional Only upgrade to minor/patch versions  |
| `Spec`                                                      | **string*                                                   | :heavy_minus_sign:                                          | body string optional Specify a branch, tag or a semver spec |
# ActiveHealthOverlayStatus


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `BinaryVersion`                                            | `string`                                                   | :heavy_check_mark:                                         | Binary version targeted by the active overlay.             | 4.12.0                                                     |
| `OverlayID`                                                | `string`                                                   | :heavy_check_mark:                                         | Active overlay identifier.                                 | **Example 1:** patch.0<br/>**Example 2:** sideload.a1b2c3d |
| `State`                                                    | `string`                                                   | :heavy_check_mark:                                         | Current overlay state.                                     | active                                                     |
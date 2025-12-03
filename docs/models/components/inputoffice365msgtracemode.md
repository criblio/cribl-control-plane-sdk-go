# InputOffice365MsgTraceMode

With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `InputOffice365MsgTraceModeSmart`  | smart                              |
| `InputOffice365MsgTraceModeAlways` | always                             |
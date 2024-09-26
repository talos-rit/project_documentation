# Talos API ICD

## Changelog

| Date | Name | Description |
|---|---|---|
| 2024-09-26 | Brooke Leinberger | Init |

Note: 
- Type names are given by the stdint.h header file in the C standard.
- All data is in big endian format

## Command Wrapper

| Arg           | Type | Description |
|---|---|---|
| Receipient    | UINT16    | Unique ID for device receiving the message |
| Sender        | UINT16    | Unique ID for device sending the message |
| Command       | UINT16    | Command for device to carry out |
| Length        | UINT16    | Length of Payload |
| Payload       | UINT8[]   | Command Info |
| CRC           | UINT16    | Checksum |

## Handshake

### Brief
This command is used by the Talos Command Interface (primarily the Director, but also the manual interface) on the Operator in order to prepare for giving further commands to the operator. Upon the Operator receiving this command, it responds with the version of the Talos Operator software it's running. The Command Interface can then use this information to determine the best commands to give this unit, or decide if it's incompatible.

### Send
**Command Value**: 0x0000

*No Payload*

### Recieve
**Command Value**: 0x8000

| Arg | Type | Description |
|---|---|---|
| Return Code   | UINT16 | |
| Major Version | UINT8  | |
| Minor Version | UINT8  | |


## Polar Pan (Min)

### Send
**Command Value**: 0x0001

| Arg | Type | Description |
|---|---|---|
| Delta Azimuth     | INT32 | Requested change in azimuth |
| Delta Altitude    | INT32 | Requested change in altitude |
| Delay (ms)        | INT32 | How long to wait until executing pan |
| Time              | INT32 | How long the pan should take to execute

### Recieve
**Command Value**: 0x8001

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error
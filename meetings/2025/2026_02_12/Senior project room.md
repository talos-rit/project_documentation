Zoom: https://rit.zoom.us/rec/share/0u9sA_kvXwq2tTPleAhQO1-1thgw0NnYbei7kconm12FFQac-N0deiZP1mfoD8H3.IO_Bw24RzFYo8uSE
## Key takeaways

- The team successfully refactored the config system using Pydantic for better validation and type safety
- Identified and fixed circular import issues in the codebase
- Discovered that ESP32S3 only has 4 independent pulse counter units, which is insufficient (need 8)
- Found an issue with RS232 communication where voltage levels were incorrect (11V vs expected 3.3V)
- Discussed the need to separate read and write operations for config files to prevent race conditions
- Identified that the bounding box prioritization PR doesn't work with YOLO model
- Decided to refactor Director pipeline in the future due to architectural issues
## Team updates


| **Speaker** |                               **Completed**                                |                                             **In Progress**                                              |                                               **Plan**                                               |                                **Blockers**                                |
|-------------|----------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------|
| **Hiroto**  | - PCMM685 implementation- Fixed RTSP buffering problem                     | - Working on hardware redevelopment- Refactoring driver on ESP                                           | - Simplifying user experience for flashing                                                           | - Hardware redevelopment is ongoing risk                                   |
| **Team**    | - Fixed bounding box prioritization- Implemented current sense calibration | - Refactoring config system with Pydantic- Working on Wi-Fi socket implementation- Working on ADC driver | - Implement encoders after current sensing is complete- Add object tracker with lighter weight model | - RS232 communication issues- ESP32S3 has insufficient pulse counter units |

## High risks

- Hardware redevelopment remains a significant risk
- RS232 communication issues with voltage levels (11V vs expected 3.3V)
- ESP32S3 only has 4 independent pulse counter units when 8 are needed
- Circular import issues in the codebase
- Director pipeline needs refactoring but may be too time-consuming
## Action items

- **Hiroto**
    - [ ] Review PR for bounding box prioritization
    - [ ] Test bounding box prioritization in another room
- **Ryan**
    - [ ] Send email to cancel ESP32S3 order and request original ESP32 instead
    - [ ] Work on Wi-Fi socket implementation
    - [ ] Complete ADC driver (Expected completion: March 24th)
- **Team**
    - [ ] Fix config system refactoring issues
    - [ ] Implement app settings for global configuration
    - [ ] Test multi-connection functionality
    - [ ] Fix TUI (terminal user interface) implementation
    - [ ] Add watchdog for monitoring processes
    - [ ] Purchase RS232 transceiver module for proper voltage level conversion
    - [ ] Back-merge main branch to fix latency issues

Zoom: https://rit.zoom.us/rec/share/m_t8QxwjoBXFLJznJ1XSiiYIvLNGH4TI98WNuoDnymvPGyOacndv82inrOEceD8T.6zB69jgVnZLH1T5M

## Key takeaways

- Wi-Fi driver works with WPA3, socket implementation is mostly complete
- Config commander refactor PR has been completed and merged
- Team decided on a new architecture for detection that separates the detector into its own process
- Video display will be removed from the UI to improve performance
- PCNT (pulse counter) component is complete but needs testing
- Team will use ESP32 instead of ESP32-S3 because it has 8 pulse counter channels vs 4
- Team is implementing unit testing framework for ESP components
## Team updates


| **Speaker** |                                  **Completed**                                  |                                      **In Progress**                                      |                       **Plan**                       | **Blockers** |
|-------------|---------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|------------------------------------------------------|--------------|
| **Hiroto**  | - Wi-Fi driver works with WPA3- Socket implementation connects and communicates | - Processing socket commands- Unit test framework for ESP                                 | - Implement API for motor control                    | N/A          |
| **Connor**  | - ADC driver connected to I2C bus                                               | - Implementing ADC driver using master bus ESP lib- Testing diodes for protection circuit | - Hook up at least one channel for testing           | N/A          |
| **Jay**     | - Pulse counter skeleton created                                                | - Writing unit tests for pulse counter                                                    | - Test with robot                                    | N/A          |
| **Team**    | - Config commander refactor- Finalized architecture for detection               | - YOLO/MediaPipe refactor- Director refactoring                                           | - Remove video from UI- Implement connection manager | N/A          |

## High risks

- Need to use ESP32 instead of ESP32-S3 due to pulse counter channel limitations
- Diodes for protection circuit need testing to ensure they can handle inductive kickback
- Current UI performance issues with video display
- Potential deadlocks in detection process termination
## Action items

- **Hiroto**
    - [ ] Complete socket command processing
    - [ ] Fix Wi-Fi SDK config
    - [ ] Implement unit test framework for ESP
- **Connor**
    - [ ] Complete ADC driver implementation
    - [ ] Test diodes for protection circuit
    - [ ] Send email to Malachowski about ordering more ESPs
- **Jay**
    - [ ] Complete pulse counter testing
    - [ ] Test with robot hardware
- **Team**
    - [ ] Remove video from UI to improve performance
    - [ ] Implement connection manager
    - [ ] Refactor director code
    - [ ] Add app settings for stream configuration
    - [ ] Fix TUI issues
    - [ ] Create midterm evaluation surveys (due today)

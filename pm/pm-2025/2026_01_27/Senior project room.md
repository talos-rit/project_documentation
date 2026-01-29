Zoom: https://rit.zoom.us/rec/share/dxZDAX55J2rmNuYVYWxyFbq1vZclVDSdBBgA7qlFsIvT2dh1Fum-ovjcS1YD-YBp.eL8j4YyPlIF7RAuo

## Key takeaways

- The team successfully tested and implemented a current sensing solution using an ADC to measure motor current
- PIV was found to be significantly faster than OpenCV for video processing, improving performance
- The team decided to use ESP32 for driver functionality, particularly for encoder handling
- Commander UI has recursion errors that were partially resolved by fixing logger issues
- The team discussed architecture options for the ESP implementation, deciding to keep some functionality on the Raspberry Pi
## Team updates


| **Speaker** |                             **Completed**                             |                             **In Progress**                              |                                      **Plan**                                       |               **Blockers**               |
|-------------|-----------------------------------------------------------------------|--------------------------------------------------------------------------|-------------------------------------------------------------------------------------|------------------------------------------|
| **Hiroto**  | - Fixed logger recursion errors- Implemented performance improvements | - Working on Commander UI improvements- Testing PIV for video processing | - Create PR for PIV implementation- Continue Commander refactoring                  | - QT DLL issues on Windows               |
| **Connor**  | - Set up ESP32 development environment- Created ESP driver project    | - Learning ESP32 programming- Testing ADC functionality                  | - Implement ESP driver for encoders- Port necessary functionality from Raspberry Pi | N/A                                      |
| **Kai**     | - Updated MediaPipe and YOLO model versions- Tested model performance | - Working on PySide 6 migration- Testing object detection models         | - Complete PySide 6 migration- Fix Windows compatibility issues                     | - PySide 6 DLL loading issues on Windows |
| **Ryan**    | - Tested ADC for current sensing- Verified current measurements       | - Working on ESP32 implementation                                        | - Determine architecture for ESP32 integration                                      | N/A                                      |

## High risks

- PySide 6 DLL loading issues on Windows preventing Commander from running properly
- Recursion errors in Commander UI that cause crashes
- Potential performance limitations with I2C communication between ESP32 and Raspberry Pi
- Time synchronization issues for video timestamps
## Action items

- **Hiroto**
    - [ ] Create PR for PIV implementation
    - [ ] Fix remaining recursion errors in Commander UI
    - [ ] Implement timestamp functionality for video frames
- **Connor**
    - [ ] Continue developing ESP driver project
    - [ ] Document ESP32 flashing process for team members
    - [ ] Test encoder functionality on ESP32
- **Kai**
    - [ ] Complete PySide 6 migration
    - [ ] Update MediaPipe and YOLO model dependencies
    - [ ] Fix Windows compatibility issues
- **Ryan**
    - [ ] Determine final architecture for ESP32 integration
    - [ ] Test current sensing with actual motors
    - [ ] Update documentation for ESP32 implementation
- **Team**
    - [ ] Decide on final architecture for ESP32 and Raspberry Pi communication
    - [ ] Update the Raspberry Pi's timezone settings
    - [ ] Investigate FastAPI for video streaming

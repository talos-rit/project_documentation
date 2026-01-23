zoom: https://rit.zoom.us/rec/share/_YkVFl0QPvmIsZyNTqBZVRasjTgdmd1dyT-mTfkLcpGM5o-aVcyb8VNSiwo8aPZi.ODCXNbpYkESnb99N
## Key takeaways

- Team is migrating from Makefiles to CMake for the build system
- Encoders may require a microcontroller (ESP32) instead of Raspberry Pi due to timing constraints
- Single speaker tracking feature is likely too complex to implement in the available timeframe
- Team will implement a simpler tracking solution based on bounding box size for demo purposes
- Static analysis tools (CPP check and ClangTidy) have been integrated into the build system
- Test coverage reporting has been set up with LCOV
## Team updates


| **Speaker** |                                                   **Completed**                                                    |                            **In Progress**                            |                                     **Plan**                                     |                       **Blockers**                       |
|-------------|--------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------|----------------------------------------------------------------------------------|----------------------------------------------------------|
| **Hiroto**  | - Converted build system from Makefiles to CMake- Set up test coverage reporting- Integrated static analysis tools | - Testing encoder timing on Raspberry Pi                              | - Potentially implement ESP32 solution for encoders- Update README documentation | - Raspberry Pi may not be fast enough for encoder timing |
| **Connor**  | - Added static analysis tools- Updated API tests                                                                   | - Working on operator refactoring- Researched single speaker tracking | - Implement simplified tracking solution for demo                                | N/A                                                      |
| **Ryan**    | - Merged driver code- Working on Pi research                                                                       | N/A                                                                   | N/A                                                                              | N/A                                                      |

## High risks

- Raspberry Pi may not be fast enough to handle encoder timing requirements
- Current ADC implementation has multiplexer issues that complicate the driver
- Single speaker tracking feature is too complex to implement in the available timeframe
## Action items

- **Hiroto**
    - [ ] Test encoder timing on Raspberry Pi
    - [ ] Research ESP32 as alternative for encoder handling
    - [ ] Update README documentation for new build system
- **Connor**
    - [ ] Implement simplified tracking solution based on bounding box size
    - [ ] Fix formatting issues in code files
- **Team**
    - [ ] Create research document on future single speaker tracking options
    - [ ] Consider applying for research computing access
    - [ ] Remove Makefiles now that CMake is implemented
    - [ ] Update time tracking spreadsheet

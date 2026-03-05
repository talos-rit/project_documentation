Zoom: https://rit.zoom.us/rec/share/DZGZfSEIg4mYvpJMiAZQhTBE36abfxxaN1tr5MTyKuwJlrOmnho3udrm5J4O4tJf.f3KWFJ0n0w9-6rM0
## Key takeaways

- ESP test framework was updated to support test groups, though it required writing a new test runner
- Checksum implementation was added between commander and operator
- ADS1015 driver was updated to handle four inputs instead of two
- PCNT driver tests are now passing
- Wi-Fi socket is complete and tested on multiple networks
- Directory refactoring PR was merged, allowing team members to work on subsections without conflicts
- The team aims to get basic motor movement (left, right, up, down) working for Milestone 2
- PiSight UI interface PR is ready to be created
- Virtual camera functionality was successfully tested in Zoom
## Team updates


| **Speaker** |                                                                                                 **Completed**                                                                                                  |                       **In Progress**                        |                                                  **Plan**                                                   |            **Blockers**            |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|------------------------------------|
| Hiroto      | - ESP unit test framework update- Checksum implementation between commander and operator- Wi-Fi socket implementation- Linux fix for commander- Directory refactor merged- Metrics tab enhancement showing FPS | - Glue code to link components together                      | - Start working on the application control code- Hardware schematics after PCNT is done                     | - Milestone 2 deadline approaching |
| Kai         | - PCNT driver tests passing                                                                                                                                                                                    | - ADS1015 driver updates to handle four inputs               | - Limit switch implementation- Hardware schematics                                                          | - Need to finish ADS1015 tests     |
| Ryan        | - Fixed Linux/Windows issues in commander                                                                                                                                                                      | - PiSight UI interface- Testing virtual camera functionality | - Create PR for PiSight UI- Refactor object model class- Implement app settings- Py Virtual Cam integration | N/A                                |

## High risks

- ESP implementation might not be completely done by Milestone 2 deadline
- The team needs to ensure safe motor movement for initial testing
- Glue code development is a substantial task that could impact timeline
## Action items

- **Hiroto**
    - [ ] Push PR for commander crash fix
    - [ ] Start working on glue code to link components
    - [ ] Test ESP with basic movement commands
- **Kai**
    - [ ] Finish ADS1015 driver review changes
    - [ ] Write tests for ADS1015
    - [ ] Implement limit switch functionality after ADS work
- **Ryan**
    - [ ] Create PR for PiSight UI interface
    - [ ] Work on app settings implementation
    - [ ] Refactor object model class to reduce spaghetti code
- **Team**
    - [ ] Test basic motor movement (left, right, up, down)
    - [ ] Finalize hardware schematics for at least one of each feature set
    - [ ] Add pull-up to the ADS
    - [ ] Find or create appropriate DB9 connector solution

Zoom: https://rit.zoom.us/rec/share/CPNU_LSHT-rWJWdbGRJXz09SkqqgU5ZGVr7MgnzisFXIFS6_VwC9xen0BKKIur0H.gA8tKdnYIcJPibnr

## Key takeaways

- The team is transitioning from refactoring to implementing features
- Hardware plan decided: using motor driver hat with current sensing for safety
- Multiple robot connection support is the next major feature to implement
- Commander architecture has been overhauled to improve performance and organization
- Time tracking sheet has been created and linked on the project website
## Team updates

| **Speaker** |                                                                              **Completed**                                                                              |             **In Progress**             |                                                         **Plan**                                                         |                 **Blockers**                 |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|--------------------------------------------------------------------------------------------------------------------------|----------------------------------------------|
| Hiroto      | - Overhauled Commander architecture- Separated manual interface logic from director/tracker- Improved UI performance with multi-processing- Added local config override | - Continuing Commander refactoring      | - Implement dropdown for object detection models- Stop director from pulling frames                                      | N/A                                          |
| Kai         | - Cleaned up issues and PRs- Merged/closed old PRs- Added issue and PR templates                                                                                        | - Working on Commander refactoring      | - Implement multiple robot connection support                                                                            | N/A                                          |
| J           | - Created inventory of existing parts- Installed Fritzing- Researched hardware options                                                                                  | - Researching current sensing solutions | - Draft message to sponsors about hardware needs- Create schematic for hardware implementation- Purchase necessary parts | - Need to finalize current sensing approach  |
| Connor      | - Removed AMQ dependency- Set up dev container for cross-platform development                                                                                           | - Working on Commander refactoring      | N/A                                                                                                                      | N/A                                          |
| Ryan        | - Created time tracking sheet- Added resources tab to website- Linked agenda and forups on main page                                                                    | - Working on documentation              | - Implement polar pan functionality- Implement homing capabilities                                                       | - Hardware implementation dependent on parts |

## High risks

- Scope risk: Driver implementation is tightly coupled with hardware decisions
- Refactoring taking too much time before implementing domain-specific features
- Current sensing implementation complexity for safety features
- Multiple robot connection support is a complex feature with many components
## Action items

- **J**
    - [ ] Create schematic for hardware implementation with current sensing
    - [ ] Draft email to sponsors about hardware purchase needs
    - [ ] Research GPIO extender chips for limit switches
    - [ ] Look into hardware emergency stop options
- **Hiroto**
    - [ ] Create dropdown for object detection models
    - [ ] Implement dropdown to switch between robots in manual interface
    - [ ] Work on supporting multiple robot connections
- **Kai**
    - [ ] Stop director from pulling frames unnecessarily
    - [ ] Implement per-robot configuration files
    - [ ] Work on supporting multiple robot connections
- **Ryan**
    - [ ] Implement polar pan functionality in the driver
    - [ ] Implement homing capabilities
    - [ ] Create issues for polar pan and homing features
- **Connor**
    - [ ] Document dev container setup
- **Team**
    - [ ] Look for network streaming application options (FFmpeg or OBS Studio)
    - [ ] Decide on approach for handling multiple video feeds
    - [ ] Implement branch protection rules for repositories

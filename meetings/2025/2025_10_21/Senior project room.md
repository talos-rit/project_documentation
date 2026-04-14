Zoom: https://rit.zoom.us/rec/share/zDCTc4Fw2iaM4n0Beoxd2cCRHEdQIAc50bwfQinvgWL3Z0Nb_JVIyBdu7OJNb3a2.vvswyhqghdJwYcjd
## Key takeaways

- Team is transitioning from strict spec-driven methodology to a more flexible, demo-driven approach with release dates and task tracking
- Hardware encoder issue resolved by using pull-up resistor; resin-printed parts allow UV light through, requiring either black resin printing or painting solution
- Commander's new manual interface with connection manager UI is in development, featuring multi-robot connection management
- Operator team discovered ACL manual mode limitations - cannot run other commands during polar pan continuous movement
- Team needs to schedule demo with Malachowski soon and prepare for video submission due November 25th
- Imaginary RIT registration planned for January; potential participation in Maker Faire end of semester as trial run
## Team updates

| **Speaker**       | **Completed**                                                                                                                                                                                                                                                                 | **In Progress**                                                                                                                                                                        | **Plan**                                                                                                                                                                   | **Blockers**                                                                                                                                                                         |
| ----------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Hiroto & Kai**  | - Modified shared memory structure from queue to buffer for video frames<br>- Fixed signal handling for multiple termination events<br>- Added feature task for adding performance metrics for Commander (frame rate counter, video stream frame rate, command response time) | - Debugging connection manager UI issues with dropdown menu<br>- Working on manual interface improvements<br>- Testing object detection performance integration                        | - Add connection manager functionality to save/load robot configs<br>- Implement multi-robot connection switching<br>- Fix termination handler to prevent concurrent calls | - Dropdown menu freezing on specific computer<br>- Application trying to close itself twice<br>- Register callback error on startup                                                  |
| **Connor & Ryan** | - Researched ACL manual mode vs direct mode functionality<br>- Identified polar pan continuous command blocking issue<br>- Tested encoder with pull-up resistor successfully                                                                                                  | - Investigating alternative methods to control robot without manual mode blocking<br>- Exploring ACL velocity commands and move commands<br>- Testing set speed command implementation | - Implement command queue protection during manual mode<br>- Find solution for non-blocking robot control<br>- Test shift command with move for single-axis movement       | - Manual mode blocks all other commands during polar pan<br>- Cannot guarantee command order with current serial implementation<br>- 9600 baud rate limitation for real-time control |
| **J**             | - Discovered resin-printed encoder discs allow UV light through<br>- Successfully tested encoder with metal disc and zip tie<br>- Confirmed encoder offset is mostly correct                                                                                                  | - Testing encoder solutions (black resin vs painting)<br>- Working on hardware assembly                                                                                                | - Contact Malachowski about black resin printing option<br>- Paint encoder discs if black resin doesn't work<br>- Complete encoder hardware integration                    | - Resin prints are UV transparent, causing encoder malfunction<br>- Need opaque material or coating for encoder discs                                                                |

## High risks

- Manual mode in ACL blocks all other commands during polar pan continuous movement, limiting concurrent command execution
- Resin-printed encoder discs allow UV light through, requiring material change or coating solution
- Video submission deadline November 25th approaching with limited time for preparation
- Dropdown menu UI issue on specific machines preventing proper testing
- Serial communication at 9600 baud may be too slow for real-time robot control requirements
## Action items

- **Hiroto and Kai**
    - [ ] Fix global lock on termination handler to prevent concurrent calls
    - [ ] Resolve dropdown menu freezing issue on specific computer
    - [ ] Complete connection manager UI implementation with config save/load functionality
    - [ ] Implement auto-connect to all saved robot configurations
    - [ ] Add performance counter for object detection process
- **Connor and Ryan**
    - [ ] Implement command queue protection to reject commands during manual mode
    - [ ] Research and test alternative ACL commands for non-blocking robot control
    - [ ] Test shift command with move command for single-axis movement
    - [ ] Verify manual mode check exists in command processing
    - [ ] Document ACL manual mode limitations and workarounds
- **J**
    - [ ] Contact Malachowski about printing encoder discs in black resin
    - [ ] Paint encoder discs if black resin option not viable
    - [ ] Complete encoder hardware integration and testing
    - [ ] Test encoder with different opacity materials
- **Team**
    - [ ] Schedule demo with Malachowski soon to show demo
    - [ ] Modify issue template to include possible completion dates
    - [ ] Create task for Imaginary RIT registration in January
    - [ ] Prepare video submission explaining project importance (Due: November 25th)
    - [ ] Research Maker Faire participation for end of semester
    - [ ] Complete midterm retrospective document and fill in rubric
    - [ ] Set release candidate dates for all repositories
    - [ ] Add product metrics tracking ticket for Operator
    - [ ] Review pending issues during Thursday meetings
    - [ ] Define clear goals and update conditions for all product metrics

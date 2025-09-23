Zoom: https://rit.zoom.us/rec/share/eNlWH1LZetQhRoIZpjG00CVz7EEK5uLoDv6En4DVFd6dXlDpMtltLtunCTsm5E0S.uKHGpmJFb5YVz85S?startTime=1758229487000
## Key takeaways

- The team verified that the gray robot (Encore) works on serial connection with camera tracking software
- The team met with Brooke from the previous team to understand the project better
- The team defined roles: Jay as hardware lead, Ryan as communications manager, Connor as Commander developer, Hero as Riker developer, and Franklin as driver developer/documentation lead
- The team identified key tasks: fixing encoders, implementing end stops, collecting multiple video streams, and controlling multiple robots simultaneously
- The project scope includes making both robots operational and creating a system to control them based on video feeds
## Team updates

| **Speaker** |                                                   **Completed**                                                    |                            **In Progress**                             |                                            **Plan**                                             |                                                                  **Blockers**                                                                   |
|-------------|--------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| Jay         | - Verified gray robot works on serial- Verified camera tracking software works- Added stabilization board to robot | - Testing all axes on Bluey robot                                      | - Verify functionality of all equipment- Find a new motor controller board with 6 motor drivers | - Camera tracking software crashes on his computer- One motor controller H-bridge is dead- Current controller only supports 4 motors but need 6 |
| Team        | - Met with Brooke to understand previous work- Determined project scope                                            | - Setting up GitHub project board- Defining roles and responsibilities | - Domain analysis- Project website setup- Determine project metrics- Research motor controllers | - Need to understand how to implement end stops- Need to research current sensing methods                                                       |

## High risks

- One side of the H-bridge is dead on a motor controller, limiting directional movement
- Current motor controller only supports 4 axes but the robot has 6 axes
- Need to implement end stops using current sensing since there are no physical limit switches
- Commander software needs major changes to support multiple robots simultaneously
## Action items

- **J**
    - [ ] Test all axes on Bluey robot
    - [ ] Research motor controller boards that can handle 6 axes
    - [ ] Meet with Professor Kaiser on Monday at 11am to get oscilloscope
- **Ryan**
    - [ ] Update the project website/GitHub organization page
    - [ ] Handle communications with sponsors and stakeholders
- **Connor**
    - [ ] Work on Commander software to handle multiple video streams
    - [ ] Research how to implement the "director" functionality for shot selection
- **Hero**
    - [ ] Work on Ichor software
    - [ ] Document the process for compiling Ichor
- **Kai**
    - [ ] Work on driver software
    - [ ] Lead documentation efforts
- **Team**
    - [ ] Research current sensing methods for end stops (Hall effect vs. series resistor)
    - [ ] Finalize project tools and methodologies
    - [ ] Complete project synopsis and website by tomorrow
    - [ ] Add tasks to GitHub project board
    - [ ] Determine preliminary roles and responsibilities

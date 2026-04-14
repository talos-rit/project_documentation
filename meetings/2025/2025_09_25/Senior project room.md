Zoom: https://rit.zoom.us/rec/share/agqAjzgovuFJhV8k6fCkcG97C04QR7tYfv4xD3L7ZnmsVTHu4F1Nl_GiNCNQ7Em9.kog9rojNd2q0YJse

## Key takeaways

- The team has finalized the project scope: getting both robots working in tandem with better framing
- The methodology plan includes 4 phases: planning (0.5 week), research and design (1 week), engineering (up to 3 weeks), and evaluation (0.5 week)
- The team will use GitHub issues and Project Board for tracking tasks
- The team is divided into two groups: Commander team (Kai and Hiro) and Driver/Hardware team (Ryan, Connor, and Jay)
- The team decided to continue with the custom driver approach rather than using the black box driver
- The Linux build is 3.9GB (vs 340MB for other platforms) and the team plans to discontinue it
## Team updates

| **Speaker** |                                           **Completed**                                           |                             **In Progress**                              |                                        **Plan**                                        | **Blockers** |
|-------------|---------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------|----------------------------------------------------------------------------------------|--------------|
| Hiroto      | - Updated the Commander code- Added frame limiting to the GUI- Organized the configuration        | - Cleaning up the Commander code- Working on connection classes          | - Continue refactoring- Start diagramming the architecture                             | N/A          |
| J           | - Updated project website- Created methodology plan- Defined metrics- Completed Pine Cellar model | - Working on hardware research- Exploring driver options                 | - Create schematic diagram- Put together bill of materials- Order necessary components | N/A          |
| Connor      | N/A                                                                                               | - Researching hardware components- Working on driver/hardware team tasks | - Work closely with Dominique on hardware integration                                  | N/A          |
| Ryan        | N/A                                                                                               | N/A                                                                      | - Work on removing ActiveMQ from Operator- Update Pi installer                         | N/A          |

## High risks

- Hardware development timeline: estimated 1-2 weeks to get the schematic done and parts ordered, with potentially 4 more weeks for debugging
- Commander code architecture is complex and difficult to understand, requiring significant refactoring
- Integration between the new hardware controller and the software will be challenging
## Action items

- **Hiroto**
    - [ ] Continue refactoring Commander code
    - [ ] Create architecture diagrams and state diagrams
    - [ ] Clean up the manual interface
    - [ ] Create issues for refactoring tasks
- **J**
    - [ ] Research hardware solutions for motor control
    - [ ] Create schematic diagram for the driver board
    - [ ] Put together bill of materials
    - [ ] Order necessary components
- **Ryan**
    - [ ] Remove ActiveMQ dependency from Operator
    - [ ] Update the Pi installer
    - [ ] Create issue templates for GitHub
- **Connor**
    - [ ] Work with J on hardware integration
- **Kai**
    - [ ] Go through branches and determine which ones to keep/delete
    - [ ] Clean up old issues from 2024
    - [ ] Update dependencies to address security issues

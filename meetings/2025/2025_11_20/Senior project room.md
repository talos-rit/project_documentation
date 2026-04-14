Zoom: https://rit.zoom.us/rec/share/6S87MJ6lZUD3t2a2dD3n-vXNeuGqSKkKviPifqrywH49WH4jlgx42LWJwheAgQY6.yx-Y_5iqqlYTllAV

## Key takeaways

- Milestone 1 is completed with most features implemented
- The team experienced hardware issues at Maker Faire when a motor driver board emitted smoke
- Current hypothesis is that back EMF from abrupt motor stopping caused a current spike
- Video for the milestone has been completed and is ready for submission
- Team is planning for Milestone 1.5 (documentation, presentations) and Milestone 2
- A new team member, Devar, has joined to help with AI model improvements
- Retrospective meeting scheduled for Friday 4-6 PM
## Team updates

| **Speaker** |                        **Completed**                        |                               **In Progress**                               |                                                    **Plan**                                                    |         **Blockers**          |
|-------------|-------------------------------------------------------------|-----------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|-------------------------------|
| **Hiroto**  | - Milestone 1- Rewrote driver to remove magic numbers       | - Investigating motor driver failure                                        | - Working on presentation- Creating incident report for hardware failure- Documentation                        | - Motor driver board failure  |
| **Ryan**    | - Initial version of unit tests for operator                | - Refactoring operator code- Implementing GitHub action for automated tests | - Re-architecting operator- Documentation- Changing serial communication approach                              | - Slow operator response time |
| **Connor**  | - Video recording and editing                               | - Documentation                                                             | - Presentation work                                                                                            | N/A                           |
| **Kaiser**  | - UI tweaks                                                 | - Documentation- Cleaning and refactoring commander                         | - Re-architecting commander- Implementing logging and performance counting- Exploring non-TKEnter UI solutions | N/A                           |
| **Team**    | - Maker Faire demonstration- Current measurement experiment | - Investigating Protobuff and gRPC implementation                           | - Milestone 1.5 completion- Planning for retrospective meeting- Exploring open source licenses                 | - Hardware issues             |

## High risks

- Motor driver board failure - one board emitted smoke during Maker Faire and is now permanently damaged
- Need to determine solution for back EMF issues with motors to prevent future hardware damage
- Current architecture causing performance issues with operator response time
- Need to decide on appropriate open source license for the project
## Action items

- **Hiroto**
    - [ ] Create incident report for motor driver failure
    - [ ] Submit completed video after team review
    - [ ] Research flyback diodes or other solutions for back EMF
- **Ryan**
    - [ ] Continue refactoring operator code
    - [ ] Implement GitHub action for automated tests
    - [ ] Document changes to operator
- **Connor**
    - [ ] Draft email to stakeholders about final meeting/demo
    - [ ] Work on presentation
- **Kaiser**
    - [ ] Document commander changes
    - [ ] Continue UI improvements
    - [ ] Research Protobuff and gRPC implementation
- **Devar**
    - [ ] Send GitHub username to get added to project
    - [ ] Review code and documentation
    - [ ] Explore simulation options for testing
- **Team**
    - [ ] Attend retrospective meeting (Friday 4-6 PM)
    - [ ] Research open source license options
    - [ ] Plan final demo/meeting with stakeholders (week after Thanksgiving)

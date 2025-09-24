## Key takeaways

- The team successfully got the Windows version of the commander working without WSL
- Discovered that the robot has limit switches that were previously unknown
- Removed PyBullet dependency which was causing installation issues
- Completed domain model and development methodology documents
- Implemented PyInstaller for easier distribution of the commander software
- Cleaned up code formatting and imports across multiple files
## Team updates

| **Speaker** |                                     **Completed**                                      |                             **In Progress**                              |                                      **Plan**                                       |                    **Blockers**                    |
|-------------|----------------------------------------------------------------------------------------|--------------------------------------------------------------------------|-------------------------------------------------------------------------------------|----------------------------------------------------|
| Hiroto      | - Implemented PyInstaller for commander- Created GitHub Actions workflow for builds    | - Cleaning up code formatting- Fixing file path handling for PyInstaller | - Improve documentation- Move code to source folder                                 | - N/A                                              |
| Kai         | - Got Windows version of commander working- Removed PyBullet dependency                | - Testing connection to robot                                            | - Test manual control of robot                                                      | - Connection issues between commander and operator |
| J           | - Investigated hardware interface- Found limit switches on robot                       | - N/A                                                                    | - Work on encoders- Get hardware specs with Professor Kaiser- Create breakout board | - N/A                                              |
| Team        | - Completed domain model- Completed development methodology- Completed project metrics | - Final plan document                                                    | - N/A                                                                               | - N/A                                              |

## High risks

- Connection between commander and operator is not working properly
- Current sensing may still be needed for the claw despite having limit switches
- Documentation for running the system is incomplete
- Multiple robots will require additional coordination
## Action items

- **Hiroto**
    - [x] Create issue for PyInstaller implementation
    - [ ] Create issue template for GitHub
    - [x] Update README with setup instructions
    - [ ] Implement GPU support for tracking algorithm
- **Kai**
    - [ ] Talk to Ryan about how he got the commander to connect to the operator
    - [ ] Test manual control with the robot
- **J**
    - [ ] Work on encoders
    - [ ] Get hardware specs with Professor Kaiser
    - [ ] Create breakout board for easier understanding of connections
- **Team**
    - [ ] Complete final plan document (Due tomorrow)
    - [ ] Meet online at 4pm tomorrow
    - [ ] Ask Brooke about limit switches and why current sensing was being pursued
    - [ ] Ask Kaiser about admin access to the repository
    - [ ] Consider adding unit testing as a stretch goal

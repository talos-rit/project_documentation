Zoom: https://rit.zoom.us/rec/share/Nmuhw42xpZOazl3yW8sCWIC6d3N7xEeybMYU5uHbF-Srzxbd7RfkVXu_fJztVd4K.-Vn8IZZc1wQBrPSS

## Key takeaways

- The team completed Milestone 1 and is planning for Milestone 1.5 (end of semester) and Milestone 2 (spring semester)
- A combined demo and retrospective meeting will be scheduled after Thanksgiving
- The team will create releases for their repositories to mark the completion of Milestone 1
- Hardware needs for the next milestone include motor drivers and encoders
- Software refactoring is needed for both Commander and Operator/Driver
- The team's methodology has evolved to be more milestone-driven, similar to Scrum
## Team updates

|   **Speaker**   |                               **Completed**                               |                             **In Progress**                              |                                **Plan**                                |                     **Blockers**                      |
|-----------------|---------------------------------------------------------------------------|--------------------------------------------------------------------------|------------------------------------------------------------------------|-------------------------------------------------------|
| **Hiroto, Kai** | - Implemented multi-connection support- Merged PR for multi-connection    | - Working on moving central data- Working on logic issues                | - Document architecture- Debug Commander issues                        | N/A                                                   |
| **Connor**      | - Created initial unit tests- Set up GitHub Actions for testing           | - Working on video editing for project- Setting up static analysis tools | - Create incident report for motor driver issues- Handle presentations | N/A                                                   |
| **Ryan**        | - Refactored driver code- Implemented MCP driver                          | - Working on PCA driver- Investigating motor driver issues               | - Focus on refactoring Operator- Implement proper C++ practices        | - Current PCA driver not working properly             |
| **J**           | - Created brickboard for limit switches- Completed hardware documentation | - Working on hardware schematics                                         | - Implement encoders- Create motor drivers                             | - Need 4-volt power supply on brickboard for encoders |

## High risks

- PCA driver is not working properly and the team doesn't fully understand why
- The team needs to figure out what's causing the object detection errors in Commander
- Random connection dropouts need to be addressed
- The architecture for Commander needs to be streamlined to fix the current roundabout data flow
## Action items

- **Hiroto, Kai**
    - [ ] Create architecture diagram for Commander
    - [ ] Debug object detection errors
    - [ ] Fix connection dropout issues
    - [ ] Create a release for Commander
- **J**
    - [ ] Complete video editing for project
    - [ ] Create incident report for motor driver issues
    - [ ] Update methodology documentation
    - [ ] Merge hardware documentation into project docs
    - [ ] Update hardware schematics
    - [ ] Design and implement encoders
    - [ ] Prepare for end-of-semester presentation
- **Ryan**
    - [ ] Schedule combined demo and retrospective meeting after Thanksgiving
    - [ ] Create agenda for the demo/retrospective meeting by Thursday
    - [ ] Continue investigating PCA driver issues
    - [ ] Refactor Operator to use proper C++ practices
    - [ ] Remove S-list implementation
    - [ ] Look into YAML parser implementation
    - [ ] Create a release for Driver/Operator
    - [ ] Create motor drivers
- Connor
    - [ ] Implement unit testing and static analysis
- **Team**
    - [ ] Opt in/out of Senior Project Excellence and Creativity Awards (overdue by 9 days)
    - [ ] Update timecards
    - [ ] Prepare for retrospective meeting

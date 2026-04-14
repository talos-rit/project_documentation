Zoom: https://rit.zoom.us/rec/share/5EVHuctHkkRZgvGagvg6c4n3fHsG4oKUnbZQLdASRSl-K3Hv31vdrwzNqQpHhtQ.xelR8wCbpy6JWuXT

## Key takeaways

- The team decided not to recruit an additional team member/co-op student as they don't have enough scope to provide 40 hours of work per week
- Multi-frame processing functionality has been implemented but needs integration with the Director component
- The team discussed refactoring the communication architecture to use synchronous processing for commands
- They identified issues with the current socket implementation for handling responses
- The team successfully submitted their rejection document for the co-op recruitment
- They discussed plans for the upcoming Maker Faire, including table staffing rotation
## Team updates

|    **Speaker**    |                                **Completed**                                 |                              **In Progress**                              |                                                          **Plan**                                                          |                              **Blockers**                              |
|-------------------|------------------------------------------------------------------------------|---------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------|
| **Ryan & Connor** | - Created and submitted the co-op rejection document                         | - Working on namespace implementation for request and response structures | - Refactor communication architecture to be synchronous                                                                    | - N/A                                                                  |
| **Hiro & Kai**    | - Implemented multi-frame processing function- Added connections to Director | - Working on Director integration- Working on fixing frame shape handling | - Continue integrating multi-frame processing with Director- Implement proper handling of multiple connections in Director | - Director implementation needs to be updated for multiple connections |
| J                 | - N/A                                                                        | - Working on script for video and presentation                            | - Continue working on hardware                                                                                             | - N/A                                                                  |
| **Hiro**          | - Working on camera mount for the robot                                      | - Creating another camera mount for the second robot                      | - N/A                                                                                                                      | - N/A                                                                  |

## High risks

- The current socket implementation doesn't handle responses well or needs custom implementation, which could cause issues when error handling is needed
- The e-stop button mounting requires drilling or recreate a container which is more complicated
- The team hasn't received confirmation about their Maker Faire application and may not have a spot
## Action items

- **Hiroto**
    - [ ] Contact Maker Faire organizers from personal email to check application status
    - [ ] Add APGL V3 license to repositories
    - [ ] Work on refactoring the communication architecture
    - [ ] Create activity diagram for user interaction flows
- **Kai**
    - [ ] Continue work on multi-frame processing
    - [ ] Test multi-frame processing with multiple cameras
    - [ ] Update Director to work with multiple connections
    - [ ] Fix the frame shape handling in Director
    - [ ] Implement proper handling of multiple connections in Director
    - [ ] Test multicam functionality using the same camera index
- **Ryan & Connor**
    - [ ] Define error codes for the communication protocol
    - [ ] Merge PR for refactoring to cpp implementation

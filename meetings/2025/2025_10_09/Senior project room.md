Zoom: https://rit.zoom.us/rec/share/D85IpKycMHSsmRx-JUGsnsfFQS0dZKuYi8IjZeGdswanX2LypLl9w2ow3pmisMub.LsQAD6gTXMEhm6ph

## Key takeaways

- The team successfully connected Commander to Operator and tested tracking functionality
- Dev container for Operator has been created with detailed setup instructions
- Hardware progress includes soldering the new optical encoder board with replacement switches
- Commander has been refactored to support multiple robots and configurations
- Documentation for Commander has been completed including architecture diagrams
- Mid-semester review needs to be scheduled with Professor Malachowski
## Team updates

| **Speaker** |                                   **Completed**                                    |                      **In Progress**                       |                                 **Plan**                                  |                                    **Blockers**                                    |
|-------------|------------------------------------------------------------------------------------|------------------------------------------------------------|---------------------------------------------------------------------------|------------------------------------------------------------------------------------|
| **Hiroto**  | - Soldered new optical encoder board with replacement switches                     | - Working on current sensing solution                      | - Label the 50-pin connector- Continue hardware work over break           | - Current sensing for bi-directional motors- Encoder alignment may need adjustment |
| **Jay**     | - Created dev container for Operator- Confirmed setup instructions are accurate    | - Config refactor for Commander- Architecture improvements | - Continue supporting multiple operator connections                       | - None mentioned                                                                   |
| **Ryan**    | - Commander documentation- UI enhancements (collapsed model buttons into dropdown) | - Working on optional dependency structure for models      | - Fix classes that don't work correctly- Continue UI improvements         | - Camera freezing issues during tracking                                           |
| **Team**    | - Fixed connection between Commander and Operator- Tested tracking functionality   | - Testing different tracking modes                         | - Add feedback from Operator to Commander- Implement polar pan and homing | - Camera tracking issues at certain distances                                      |

## High risks

- Current sensing solution for bi-directional motors needs to be resolved
- Camera freezing issues during tracking mode need investigation
- Operator doesn't receive feedback when robot reaches movement limits
- Encoder alignment may be off due to different component footprints
## Action items

- **Team**
    - [ ] Schedule mid-semester review with Professor Malachowski for early next week
    - [ ] Update time tracking sheet with actual hours worked (not just meetings)
    - [ ] Document the process for updating Operator
    - [ ] Consider moving all asset files into a single directory
    - [ ] Add hardware-based failsafe for when software fails to detect limits
- **Ryan**
    - [ ] Fix the issue with camera freezing during tracking
    - [ ] Continue UI enhancements for Commander
    - [ ] Implement support for multiple robot configurations
- **Jay**
    - [ ] Fix the connection issue between Commander and Operator
    - [ ] Continue working on supporting multiple operator connections
- **Hiroto**
    - [ ] Update the ICD to add open-close arm functionality to manual interface
    - [ ] Continue work on current sensing solution
    - [ ] Label the 50-pin connector


![[IMG_6143.mp4]]
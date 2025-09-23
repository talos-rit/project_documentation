Zoom https://rit.zoom.us/rec/share/erHmHT0qx8gjByMxLDCneePpuDrSvuy4GG-xUmDj6EyPupTFnPeadd0cCj_PcK4.cuN22MBasMrJb2gY

## Key takeaways

- Brook demonstrated the ER5+ robot arm is functional with a custom driver board she developed
- The team now has a clearer understanding of both robot arms' capabilities and limitations
- The ER4U robot arm has compatibility issues with the ER5+ but could potentially work with modifications
- The project scope includes getting both robots working together, with one potentially following the other
- The original controller has limitations (9600 baud, single-axis movement) that the custom solution addresses on the bluey
- Brook provided access to documentation including ACL commands and wiring diagrams
## Team updates

| **Speaker** |                                                                 **Completed**                                                                  |                              **In Progress**                               |                   **Plan**                    |                                      **Blockers**                                       |
|-------------|------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------|-----------------------------------------------|-----------------------------------------------------------------------------------------|
| Brook       | - Developed custom driver board for ER5+- Got basic motor movement working- Implemented encoder interrupt detection- Setting up Discord access | - Was working on current detection- Implementing ADC for current detection | N/A                                           | - H-bridge only works in one direction- Left the project before completing all features |
| Ryan        | - Got the ER5+ running in manual mode- Reset Raspberry Pi password to gain access                                                              | - Exploring project scope possibilities                                    | - Define team roles now that scope is clearer | - Previously couldn't access documentation- Raspberry Pi WiFi configuration issues      |

## High risks

- The H-bridge on the custom driver board is damaged and only works in one direction
- Serial console occasionally bugs out and requires system reset
- Communication between the two different robot models (ER5+ and ER4U) may be challenging due to incompatible pin configurations
- The project scope may be ambitious given the technical challenges
## Action items

- **Team**
    - [ ] Define individual roles based on the clarified project scope
    - [ ] Access and review the ACL documentation for command reference
    - [ ] Fix the Raspberry Pi WiFi configuration (currently set up as access point)
    - [ ] Determine if the ER4U and ER5+ pin configurations are compatible
    - [ ] Consider purchasing a new H-bridge (~$20) to fix the driver board
    - [ ] Manage the critical path to avoid progress bottlenecks

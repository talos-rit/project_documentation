Zoom: https://rit.zoom.us/rec/share/tGXXRSWFSs-FEsaHUTrDBPWXP-9-T1VJ3r-5WtsTgOX4p9upiHz7e2vrsNGS9Q7P.4bLax6g3R-F9hGki

## Key takeaways

- Brooke provided background on the Talos robot project, explaining that there are two robots: the gray ER5 (working) and the blue ER4 (partially working)
- The blue robot (Bluey) has encoder issues but can be operated with the Icker driver software that Brooke developed for the Raspberry Pi
- The digital twin feature was abandoned and ActiveMQ dependencies should be removed
- The Commander interface crashes when tracking mode is enabled, likely due to GPU/processing limitations
- Previous team struggled with poorly defined scope from Malikowsky (who was both coach and sponsor)
- Brooke will provide the missing Ichor driver code and additional hardware components
## Team updates

| **Speaker** |                                   **Completed**                                   |              **In Progress**              |                            **Plan**                            |                          **Blockers**                           |
|-------------|-----------------------------------------------------------------------------------|-------------------------------------------|----------------------------------------------------------------|-----------------------------------------------------------------|
| **J**       | - Got operator compiled- Tested manual robot control                              | - Working on Commander interface          | - Work on Bluey to get it operational- Focus on encoder issues | - Commander crashes in tracking mode- Missing Icker driver code |
| **Brooke**  | - Developed Icker driver for Raspberry Pi- Created custom hardware for blue robot | - Looking for hardware components at home | - Bring ADC board by Thursday- Share Icker driver code         | - Traffic preventing in-person attendance                       |
| **Ryan**    | - Tested robot functionality                                                      | - Working on webcam integration           | N/A                                                            | - WSL camera access issues                                      |
| **Connor**  | - Explored robot functionality                                                    | N/A                                       | N/A                                                            | N/A                                                             |

## High risks

- Commander interface crashes when tracking mode is enabled
- Missing encoder functionality for the blue robot (Bluey)
- Scope definition is vague, similar to previous team's experience
- Digital twin was abandoned and may not be recoverable
- ActiveMQ dependencies may cause compilation issues
## Action items

- **Brooke**
    - [ ] Provide Icker driver code from laptop
    - [ ] Bring ADC board (ADS1158/1258) by Thursday
    - [ ] Look for Malikowsky's Raspberry Pi before requesting new ones
    - [ ] Share additional hardware components for the blue robot
- **Team**
    - [ ] Schedule follow-up meeting with Alex to discuss Commander interface
    - [ ] Remove ActiveMQ dependencies from codebase
    - [ ] Define project scope more clearly than previous team
    - [ ] Review ICD documentation (described as "the Bible" for communications)
    - [ ] Investigate GPU options for running tracking functionality
    - [ ] Submit project synopsis and team website by Friday
- **J**
    - [ ] Work on getting Bluey operational with Ichor driver
    - [ ] Email Alex with Commander-specific questions

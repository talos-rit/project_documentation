Zoom: https://rit.zoom.us/rec/share/_94Kcf5n7c0KGtdGylVfYX7MAybuZoLCJ_uAVzCK_JgvbnR4t_JgJTuMdIixZgo4.yVRpUS_obWhJpKPm

## Key takeaways

- The team discussed refactoring the architecture to decouple the director and detector components
- Decided to use IPC (Inter-Process Communication) for communication between components rather than more complex solutions like Pyro5 or REST APIs
- Identified issues with the current SD card on the Raspberry Pi (only showing 6.8GB of 64GB) and fixed it using raspi-config
- Successfully merged PRs for fixing connection management and icon display in the title bar
- Discussed future features including documentation hosting on GitHub Pages, LLM-based video direction, and cloud-based model processing
## Team updates


| **Speaker** |                                                           **Completed**                                                           |                             **In Progress**                             |                                         **Plan**                                         | **Blockers** |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|------------------------------------------------------------------------------------------|--------------|
| **Hiroto**  | - Fixed connection management bug- Added icon to application title bar- Created document for future features- Added web viewer PR | - Working on architecture refactoring- Testing web viewer functionality | - Deprecate "keep away" logic- Refactor Director to be simpler- Remove person lost logic | N/A          |
| **J**       | - Set up unit testing framework- Implemented Wi-Fi functionality                                                                  | - Working on I2C implementation- Working on PCI driver                  | - Continue work on ESP32                                                                 | N/A          |

## High risks

- Current architecture makes it difficult to implement cloud-based detection or multi-robot support
- YOLO model only returns one bounding box, causing issues when tracking multiple people across different cameras
- Docker installation on Raspberry Pi taking up significant disk space
## Action items

- **Hiroto**
    - [ ] Refactor Director to decouple it from Detector
    - [ ] Deprecate "keep away" logic
    - [ ] Remove person lost logic from Director
    - [ ] Document troubleshooting steps for DNS cache issues
- **Team**
    - [ ] Implement IPC communication between components
    - [ ] Test performance of different architecture options
    - [ ] Consider implementing plugin architecture for models
    - [ ] Add authentication for component communication

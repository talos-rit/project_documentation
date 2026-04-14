Zoom: https://rit.zoom.us/rec/share/zx5VuLOKrlqD1mRr6IkukBsLQsniMxZbfC4bNuCzSNz0DlWW6_rOTpvzRX3u9S0.1WzFMdpyefg98DcI

## Key takeaways

- The team is working on a camera tracking system with separate control and filming streams
- Current focus is on getting basic functionality working before implementing advanced features
- Project organization needs improvement, including better documentation and release management
- The repository structure includes multiple components (Operator, Commander, Driver, Cam Streamer)
- The team is planning to implement a virtual camera system that could work with applications like Zoom
## Team updates

| **Speaker** |                                   **Completed**                                   |                                         **In Progress**                                          | **Plan** | **Blockers** |
|-------------|-----------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------|----------|--------------|
| Hiroto      | N/A                                                                               | N/A                                                                                              | N/A      | N/A          |
| Adam        | N/A                                                                               | Working on branch with multiple issues including per-robot config files and non-static publisher | N/A      | N/A          |
| Dave        | Created a new repository for streaming camera footage over network (Cam Streamer) | N/A                                                                                              | N/A      | N/A          |

## High risks

- Unclear documentation about repository structure and component relationships
- Submodule in the Driver repository causing potential issues
- Lack of clear release management process
- Incomplete implementation of agile methodology (tasks instead of proper user stories)
## Action items

- **Team**
    - [ ] Improve documentation to clarify repository structure and component relationships
    - [ ] Consider absorbing the Driver submodule into the main repository
    - [ ] Implement proper release management with ongoing documentation
    - [ ] Create README files for releases
    - [ ] Consider creating an installer that uses all components
    - [ ] Refine project management approach to better align with agile methodology
    - [ ] Develop plan for handling two video streams (decision point between cameras)
    - [ ] Work toward real-time functionality for use as a Zoom camera

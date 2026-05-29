# Onboarding Docs

Welcome to Talos-RIT. The goal of this document is to get you up to speed on the past, current, and future goals and progress of this project, and how to get setup. Also if there are any part of the project that is believed to be lacking from this document please feel free to add them, we should expect to be the living document for the entire project.

# Original Project Goals
This project started in Fall 2023, with the goal of taking the ER-V and getting it to track a human subject. The idea presented by the sponsor was to have smooth, interesting and engaging recorded lecture video that goes beyond the static webcam. It was to act as a autonomous human camera operator which handles subject framing, movement, slow transitions, and much more.

## History

The team from Fall 2023-Spring 2024: The goal of this iteration was to get basic manipulation functionality and setup project to a point of basic functionality. They are able to accomplish basic MVP on the ER-V and basic camera tracking on the robot. They were able to get the very basic goal of the project with commander running object tracking with several bottle necks in performance. The operator was able to do basic communication with ER-V via serial and communicate fully with commander using socket. In addition, the ER-4pc was a side project with more hardware requirements, and they accomplished basic motor manipulation.

The team from Fall 2025-Spring 2026: The goal of this iteration was to get 2 robots working in tandem. This included building/finishing the custom controller for the ER-4pc, and allowing the commander application to connect to 2 robots simultaneously. The outcome for this team was the ability to run the commander application using object recognition at rapid speed and accuracy and ability to connect multiple robots or operators. It also has wide support of operating systems and the ability export streams from commander. There were several scope cuts but was able to get partially working ER-4pc with maintainability improvements. Not many improvements were made for operator due to shift in focus for ER-4pc but was able to make non-functional features such as complete transfer to C++ and utilizing more C++ standard library.

# Project Demos
We demonstrated our project at both Makerfaire and Imagine RIT. Past application info can be found at [**Makerfaire**](makerfaire/README.md) and [**Imagine RIT**](imagine/README.md). 
Additionally, we made 4 end-of-milestone demos for our sponsor, Malachowsky, and other department staffs.

# [Project Architecture](technical/architecture/README.md)
Please read this document to get a grasp on the overall architecture of the project before continuing with setup. This will help determine which sections of the project you may care to work on, and how they may communicate with each other.

# Development Setup
Each application has different requirements to setup. Please view each repository link to get setup with each project.

- [Commander](https://github.com/talos-rit/commander/blob/main/README.md): Link to the Commander README to get setup on the commander app.
- [Operator](https://github.com/talos-rit/operator/blob/main/README.md): Link to the Operator README to get setup and how to build/run the Operator app on Linux.
- [ESP Driver](https://github.com/talos-rit/esp-driver/blob/main/README.md): Link to the ESP Driver README on how to setup ESP IDF.
- [Camera Streamer](https://github.com/talos-rit/cam_streamer/blob/main/README.md): Link to the Cam streamer on how to setup Cam streamer on Linux.

# Due Dates
1. The most important due dates to remember is Maker Faire(middle of Fall semester) and Imagine RIT(end of Spring semester). Again please review the [linked document](<#Project Demos>) for accurate dates and instruction for applications. 
2. Expect to schedule milestone demos outside of the event fairs with the sponsor and other department staffs. This should happen close to event faires mentioned above as well as two additional demo scheduled in middle of the two semester to make 4 total quarterly updates. This should include actual technical demos of the robots working with commander to show the full working systems.

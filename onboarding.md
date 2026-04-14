# On-boarding Docs

Welcome to Talos-RIT. The goal of this document is to get you up to speed on the past, current, and future goals and progress of this project, and how to get setup.

# Original Project Goals
This project started in Fall 2023, with the goal of taking the ER-V and getting it to track a human subject. The idea presented by the sponsor was to have smooth, interesting and engaging recorded lecture video that goes beyond the static webcam. It was to act as a autonomous human camera operator which handles subject framing, movement, slow transitions, and much more.

The first senior project team who started this from Fall 2023-Spring 2024 accomplished basic MVP on the ER-V and accomplished basic camera tracking on the robot, and although there were many bugs and bad code practices, they accomplished the very basic goal. In addition, the ER-4pc was a side project, and they accomplished getting the motors to move, though without encoders, limit switches, or without any features that would make it human safe.

The second senior project team took this from Fall 2025-Spring 2026. The goal of this iteration was to get 2 robots working in tandem. This included building/finishing the custom controller for the ER-4pc, and allowing the commander application to connect to 2 robots simultaneously. The end outcome due to scope cuts were the partially working ER-4pc, maintainability improvements, improvements in the tracking models, as well as the ability to connect to 2 robots simultaneously.

# [Project Architecture](technical/architecture/README.md)
Please read this document to get a grasp on the overall architecture of the project before continuing with setup. This will help determine which sections of the project you may care to work on, and how they may communicate with each other.

# Development Setup
Each application has different requirements to setup. Please view each to get setup with each project you would like to develop on.

## [Commander](https://github.com/talos-rit/commander/blob/main/README.md)
Link to the Commander README to get setup on the commander app.

## [Operator](https://github.com/talos-rit/operator/blob/main/README.md)

> [!NOTE] NOTE
> The Operator application can only be compiled on Linux, and is designed to run and be compiled on a Raspberry Pi. More details in the README show how to develop on WIndows/MacOS.

Link to the Operator README to get setup and how to build/run the Operator app.

## [ESP Driver](https://github.com/talos-rit/esp-driver/blob/main/README.md)
Link to the ESP Driver README on how to setup ESP IDF.
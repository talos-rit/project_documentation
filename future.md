
> [!WARNING] This is just some list of ideas that we think could be good to have people working on in the future, and it is not a required feature set or supported feature with the current architecture. For every feature please discuss with the sponsor and do an in depth analysis on architecture to make sure the feature is plausible and actually great to work on.

# General Efforts
1. Collaboration with graduate researchers and work with entities like research computing

# Commander

1. Plugin architecture
	1. We have a very unreliable dependency on object detection models where mediapipe is not the most reliable library and yolo is good, but could be problematic with licensing. 
	2. Any future object detection model will require change in the mechanisms for commander and would much rather benefit from a plugin architecture.
2. No interfaced option
	1. currently there are two interfaces for commander where we have tkinter's GUI and textual's TUI, but neither is great for if you just want to run a single command to listen to a video stream, detect objects, and send commands to operator.
	2. We could make a CLI command for running minimal amount of code.
3. Support for python notebooks
	1. While this could already be done, it might be more interesting to have development done in a notebook to develop our own object detection model or more advanced algorithms. The notebook could act as a playground/platform to support research in these areas.

# Operator

1. TBD

# Hardware

1. TBD

> [!WARNING] This is just some list of ideas that we think could be good to have people working on in the future, and it is not a required feature set or supported feature with the current architecture. For every feature please discuss with the sponsor and do an in depth analysis on architecture to make sure the feature is plausible and actually great to work on.

# General Efforts
1. Collaboration with graduate researchers and work with entities like research computing
2. Binocular Vision framing and capture

# Commander

1. Plugin architecture
	1. We have a very unreliable dependency on object detection models where mediapipe is not the most reliable library and yolo is good, but could be problematic with licensing. 
	2. Any future object detection model will require change in the mechanisms for commander and would much rather benefit from a plugin architecture.
2. No interfaced option
	1. currently there are two interfaces for commander where we have tkinter's GUI and textual's TUI, but neither is great for if you just want to run a single command to listen to a video stream, detect objects, and send commands to operator.
	2. We could make a CLI command for running minimal amount of code.
3. Support for python notebooks
	1. While this could already be done, it might be more interesting to have development done in a notebook to develop our own object detection model or more advanced algorithms. The notebook could act as a playground/platform to support research in these areas.
4. Port UI to another UI framework
	1. While tkinter is a great library for building simple UIs, it is very limited in terms of features and capabilities. Another framework, such as PyQt or PySide, would allow us to build more advanced UIs with better performance and more features. tkinter has shown to be very limiting when it comes to performance while including video streams. (See `pyside6-migration` branch for a start)

# Operator

1. Rewrite into embedded systems like esp32, stm32, or arduino.
	1. We don't need the full capability of raspberry pi and is more of an overhead for writing hardware components. With a caveat of video streaming device also required.
2. 

# Hardware

1. Simpler solution for ESP32 flasher using github pages(?)
2. Create comprehensive end user guide to setup all parts the system
	1. Include common tips like the 3D printed encoder [https://www.thingiverse.com/thing:3388595]
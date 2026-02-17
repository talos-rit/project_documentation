
> [!WARNING] This is just some list of ideas that we think could be good to have people working on in the future, and it is not a required feature set or supported feature with the current architecture. For every feature please discuss with the sponsor and do an in depth analysis on architecture to make sure the feature is plausible and actually great to work on.

# General Efforts
1. Collaboration with graduate researchers and work with entities like research computing
2. Binocular Vision framing and capture
3. Documentation hosted on github pages maybe using mkdocs to make it easier to maintain and update documentation. 
	- This would allow us to have a more comprehensive and organized documentation for the project, which would be helpful for both current and future developers. It would also make it easier for users to find the information they need to set up and use the system.

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
5. LLM as a video director
	1. The LLM could make decisions on which video feed to show or possibly even what commands to run based on the current situation. This could be especially useful in a multi-robot scenario where there are multiple video feeds or objects to track.
	2. Due to current real-time limitations of LLMs and resources required, this would likely be more of a high level decision maker and not something that could be used for real-time object tracking or control. Maybe do the processing of the video feeds after the fact. Deciding which video feed to show is probably more feasible.
	3. Could create an Model Context Protocol (MCP) server to make info about the robot or robots that are connected to commander available to the LLM.
		- Might be more helpful if a digital twin is implemented
	4. Possibly Relevant or Helpful Resources:
		- [VideoDirectorGPT: Consistent Multi-Scene Video Generation via LLM-Guided Planning](https://videodirectorgpt.github.io/)
		- [Dispider: Enabling Video LLMs with Active Real-Time Interaction via Disentangled Perception, Decision, and Reaction](https://arxiv.org/abs/2501.03218)
		- [VideoLLM-online: Online Large Language Model for Streaming Video](https://showlab.github.io/videollm-online/)
		- [Video Production Live Streaming: Integrating AI Director Feedback in Real-Time](https://reelmind.ai/blog/video-production-live-streaming-integrating-ai-director-feedback-in-real-time)
		- [LLM-Alignment Live-Streaming Recommendation](https://arxiv.org/html/2504.05217v1)
6. Local Organization-wide Database for robot configs
	1. Organizations could create a database of robot configurations (contains connection details utilized by commander), hosted locally by the organization, so that users can easily connect to new robots. Someone like an IT administrator could setup this database of configs based on the robots that they have already setup in different rooms. These configs could be easily added to the list of available connections in commander.
	2. This would likely fit best in a different repo as its own project, possibly with its own frontend, but could be integrated with commander.
	3. Some possible ideas for integrating with commander:
		- Create a front-end, likely some sort of web-app, for browsing the configs and allow users to download the config file. Then there could be a button in commander to import the robot config file and add it to the list of available connections.
		- Create an API for the database and allow commander to interface with the database and browse the available configs. This would allow for a more seamless integration and make it easier for users to find and connect to new robots.
7. Health monitoring
	1. Implement a health monitoring system that can notify the UI if there are any issues with the robot such as trying to connect to the socket. Currently, the tkinter GUI does not display any information about the connection status of the robot, so if there is an issue with the connection, the user would not know unless they checked the logs in the terminal. While this is available in the TUI, commander does not properly handle the case where the connection to the robot can't be established. Even though the system currently has a retry mechanism, it does nothing if the connection can't be established after 5 tries.
	2. This could also check if the object detection process is alive and restart it if it has died.
8. Yolo Related Improvements
	1. Optimize yolo model export format based on hardware
		1. Currently we just use the default `.pt` format for our yolo models, but this is not optimized for performance on all hardware. For example, on some hardware, using the OpenVino format (for Intel hardware), ONNX format, or TensorRT (for NVIDIA hardware) format could significantly improve performance. We could implement a system that detects the hardware that commander is running on and automatically converts the yolo model to the optimal format for that hardware. This would allow us to get better performance out of our object detection models without requiring users to manually convert the models themselves.
		2. Onnx format could be promising since AI says that yolo just figures out which runtime to use based on the hardware, but we would need to do some testing to see if this is actually the case and if it provides a significant performance boost. This could eliminate the need to switch between different formats for different hardware and make it easier for users to get good performance out of their object detection models regardless of the hardware they are using.
		- An issue could be installing the right onnx dependencies based on hardware since it is hard to detect which execution provider would be best. Also, installing all of the onnx dependencies we could possibly need can be a burden and bloat the application. Might benefit from targeting certain systems with specific hardware at first to make this easier to implement. Since it can be difficult to obtain all of the proper versions of dependencies for all hardware, we should keep an eye on [PEP 817](https://discuss.python.org/t/pep-817-wheel-variants-beyond-platform-tags/105860).
	2. Training on a people tracking/detection dataset
		1. The current yolo model is trained on the coco dataset which is a very general dataset and not specific to our use case. We could fine tune the model on a more specific dataset that is focused on people tracking/detection, which would likely improve the performance of the model for our use case. This would involve collecting a dataset of images that are relevant to our use case (e.g. images of people in a classroom environment) and then fine tuning the yolo model on this dataset. This could potentially improve the accuracy of the object detection and make it more reliable for use in commander.
		- Possible Datasets:
			- [ClassRoom-Crowd: A Comprehensive Dataset for Classroom Crowd Counting and Cross-Domain Baseline Analysis](https://www.mdpi.com/2673-4591/78/1/10)
		2. This could be combined with pruning to make the models smaller and hypothetically only detect people.
9. Running models in the cloud
	1. Running the object detection models in the cloud could potentially allow us to use more powerful models that would not be able to run on the local hardware. This would involve sending the video feed from the robot to a cloud server where the object detection model is running, and then sending the results back to commander. This could allow us to use more advanced models that require more computational resources, but it would also introduce latency and require a stable internet connection, so it may not be suitable for all use cases.

# Operator

1. Rewrite into embedded systems like esp32, stm32, or arduino.
	1. We don't need the full capability of raspberry pi and is more of an overhead for writing hardware components. With a caveat of video streaming device also required.
2. 

# Hardware

1. Simpler solution for ESP32 flasher using github pages(?)
2. Create comprehensive end user guide to setup all parts the system
	1. Include common tips like the 3D printed encoder [https://www.thingiverse.com/thing:3388595]
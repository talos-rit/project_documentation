
> [!WARNING] This is just some list of ideas that we think could be good to have people working on in the future, and it is not a required feature set or supported feature with the current architecture. For every feature please discuss with the sponsor and do an in depth analysis on architecture to make sure the feature is plausible and actually great to work on.

## Contents
- [[#General Efforts|General Efforts]]
- [[#Commander|Commander]]
- [[#Operator|Operator]]
- [[#ESP Driver|ESP Driver]]
- [[#Hardware|Hardware]]

## General Efforts
1. Collaboration with graduate researchers and work with entities like [research computing](https://www.rit.edu/researchcomputing/)
2. Binocular Vision framing and capture for depth awareness and enhanced tracking
3. Documentation hosted on github pages maybe using mkdocs to make it easier to maintain and update documentation. 
	- This would allow us to have a more comprehensive and organized documentation for the project, which would be helpful for both current and future developers. It would also make it easier for users to find the information they need to set up and use the system.
4. Authentication for connection between commander and a robot
	- Would require integration on both commander and operator/driver side
5. Dynamic Robot discovery
	- Commander should be able to discover robots available on the network. This could be done similarly to how computers find printers or like how some smart home apps find devices that can be controlled.
	- The robots themselves (operator) could possibly hold a config file for commander to use for that particular robot

## Commander

1. Plugin architecture
	1. We have a very unreliable dependency on object detection models where mediapipe is not the most reliable library and yolo is good, but could be problematic with licensing. 
	2. Any future object detection model will require change in the mechanisms for commander and would much rather benefit from a plugin architecture.
2. Turn commander into a importable package
	1. For easier distribution we should ship the commander application to PyPI.
3. Port UI to another UI framework
	1. Js framework like react.js or svelte.js could also be a good option. While I(Hiro) don't recommend it, you could look into electron.js for more local aproach.
	2. This will come with the limitation of support for pyvcam
4. LLM as a video director
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
5. Local Organization-wide Database for robot configs
	1. Organizations could create a database of robot configurations (contains connection details utilized by commander), hosted locally by the organization, so that users can easily connect to new robots. Someone like an IT administrator could setup this database of configs based on the robots that they have already setup in different rooms. These configs could be easily added to the list of available connections in commander.
	2. This would likely fit best in a different repo as its own project, possibly with its own frontend, but could be integrated with commander.
	3. Some possible ideas for integrating with commander:
		- Create a front-end, likely some sort of web-app, for browsing the configs and allow users to download the config file. Then there could be a button in commander to import the robot config file and add it to the list of available connections.
		- Create an API for the database and allow commander to interface with the database and browse the available configs. This would allow for a more seamless integration and make it easier for users to find and connect to new robots.
6. Health monitoring
	1. Implement a health monitoring system that can notify the UI if there are any issues with the robot such as trying to connect to the socket. Currently, the tkinter GUI does not display any information about the connection status of the robot, so if there is an issue with the connection, the user would not know unless they checked the logs in the terminal. While this is available in the TUI, commander does not properly handle the case where the connection to the robot can't be established. Even though the system currently has a retry mechanism, it does nothing if the connection can't be established after 5 tries.
	2. This could also check if the object detection process is alive and restart it if it has died.
7. Yolo Related Improvements
	1. Optimize yolo model export format based on hardware
		1. Currently we just use the default `.pt` format for our yolo models, but this is not optimized for performance on all hardware. For example, on some hardware, using the OpenVino format (for Intel hardware), ONNX format, or TensorRT (for NVIDIA hardware) format could significantly improve performance. We could implement a system that detects the hardware that commander is running on and automatically converts the yolo model to the optimal format for that hardware. This would allow us to get better performance out of our object detection models without requiring users to manually convert the models themselves.
		2. Onnx format could be promising since AI says that yolo just figures out which runtime to use based on the hardware, but we would need to do some testing to see if this is actually the case and if it provides a significant performance boost. This could eliminate the need to switch between different formats for different hardware and make it easier for users to get good performance out of their object detection models regardless of the hardware they are using.
		- An issue could be installing the right onnx dependencies based on hardware since it is hard to detect which execution provider would be best. Also, installing all of the onnx dependencies we could possibly need can be a burden and bloat the application. Might benefit from targeting certain systems with specific hardware at first to make this easier to implement. Since it can be difficult to obtain all of the proper versions of dependencies for all hardware, we should keep an eye on [PEP 817](https://discuss.python.org/t/pep-817-wheel-variants-beyond-platform-tags/105860).
	2. Training on a people tracking/detection dataset
		1. The current yolo model is trained on the coco dataset which is a very general dataset and not specific to our use case. We could fine tune the model on a more specific dataset that is focused on people tracking/detection, which would likely improve the performance of the model for our use case. This would involve collecting a dataset of images that are relevant to our use case (e.g. images of people in a classroom environment) and then fine tuning the yolo model on this dataset. This could potentially improve the accuracy of the object detection and make it more reliable for use in commander.
		- Possible Datasets:
			- [ClassRoom-Crowd: A Comprehensive Dataset for Classroom Crowd Counting and Cross-Domain Baseline Analysis](https://www.mdpi.com/2673-4591/78/1/10)
		2. This could be combined with pruning to make the models smaller and hypothetically only detect people.
8. Running models in the cloud
	1. Running the object detection models in the cloud could potentially allow us to use more powerful models that would not be able to run on the local hardware. This would involve sending the video feed from the robot to a cloud server where the object detection model is running, and then sending the results back to commander. This could allow us to use more advanced models that require more computational resources, but it would also introduce latency and require a stable internet connection, so it may not be suitable for all use cases.
	2. If we want production grade services to offload to a server, look into [Roboflow](https://roboflow.com/) as ideas
9. Digital Twin
	1. Creating a digital twin of the robot and its environment could allow us to simulate the robot's behavior and test out different algorithms and strategies in a virtual environment without needing the physical robot. This could be especially useful for testing out new features in commander without risking damage to the physical robot. The digital twin could be created using a simulation software, and it could be integrated with commander to allow for seamless switching between the physical robot and the digital twin.
	2. The original team attempted this for the ER 4 PC with pybullet but ran into some issues and ultimately ended up abandoning it. Could be worth revisiting. This repo might be able to help out with simulation of the ER V robot: [R2D3](https://github.com/ajnsit/r2d3). It has a `.blend` file for the ER V robot that could hopefully be used in a simulation environment. We would likely need to convert this file to a format that is compatible with the simulation software we choose to use, but it could be a good starting point for creating a digital twin of the ER V robot.
	3. THIS MUST BE ITS OWN APPLICATION
		1. The commander alone is somewhat bloated in features and extensions. Adding more large scale features such as this would create maintainability problems in the long term.
		2. This will require massive architectural rework to export hardware data from commander in realtime.
		3. Look into grpc, socket, message queues to list a few potential ideas
10. Manual or automated bbox prioritization feature
	1. Currently the process for person detection is the following:
	   Stream > detect multiple people > prioritize > return output
	2. Currently we have wide flexibility for detect multiple people which is great, but not very important in terms of the process because we can only detect so much people in the frame and don't have problems doing. But the area that commander lack currently is a post processing step, prioritize, where we can prioritize the detected people and select one instance to focus on. This is the most crucial step, even more than the person detection, where there are many active research going on in this area.
	3. The goal for prioritization is to rank the detected people and select the single best score to focus on but this ranking method can be done in multiple ways listed below.(easiest to hardest)
		1. We can manually select the person to focus on which could work until the person is lost out of sight or detection loses the person.
		2. BBox size: using the overall coverage of the frame we can prioritize bbox that takes either the largest or smallest or in between.
		3. Speed: tracking people overtime and their movement could allow us to reject fast moving object early on for people that can not be efficiently tracked using the robot.
		4. Time duration: using ID's in yolo's tracking model we can keep track of how long each person is in the frame allowing us to time for how long the person is in the frame for. 
		5. Priority by depth: using monocular vision we can estimate depth using a model. See https://github.com/DepthAnything/Depth-Anything-V2.
		6. Saliency ranking: using a model trained on saliency we can rank the bboxes. 
		   See https://arxiv.org/abs/2203.09416
		7. Language guided Saliency ranking: using a model with multimodal contexts we can rank the bboxes.
		   See https://arxiv.org/abs/2203.09416
11. Filtering based on user defined zone
	1. Some computer vision systems, like the trackman launch monitor, require the user to calibrate the launch monitor by defining the zone/area where detections should occur. This aids object tracking by allowing the model/system to filter out detections that are out of that range. 
	2. For commander, this could allow users to define a zone at the front of a classroom where valid detections should occur. The user would be able to define this zone upon starting up commander and connecting to the robot. Assuming the robots would not change rooms or positions super often, this zone could be persisted across different runs of the application. Like trackman, the user would still be prompted to calibrate the robot by redrawing the zone or review the existing zone when connecting to the robot for the first time since starting up commander.
	3. One possible way to do this could be using `supervision`. See these links below for some possibly helpful articles from the `supervision` documentation:
		- [How To Analyze Occupancy with Supervision](https://supervision.roboflow.com/develop/notebooks/occupancy_analytics/)
		- [Count Objects Crossing the Line](https://supervision.roboflow.com/develop/notebooks/count-objects-crossing-the-line/)
		- [Filter Detections](https://supervision.roboflow.com/develop/how_to/filter_detections/)
		- [Count in Zone](https://supervision.roboflow.com/develop/how_to/count_in_zone/)
	3. Please note more research is needed to determine if this will actually work and/or be helpful for our use case.

## Operator

1. Rewrite into embedded systems like esp32, stm32, or arduino.
	1. We don't need the full capability of raspberry pi and is more of an overhead for writing hardware components. With a caveat of video streaming device also required.

## ESP Driver
1. Be Human Safe; detect when the robots hit something
	1. By using a dynamic function comparing the velocity of the motor and the current of the motor, detect when the arm hits an object (like a human or an end stop), but DONT false trigger when a user isntantaneously starts motion of a motor, because motors have a spike of current when started.
2. Smooth Motion Profiling
	1. make it so motors start slow, ramp up in speed, and ramp down when done moving. look into PID (position; integral; differential) controllers for point to point motion, but for up-down-left-right control the motion wont be as snappy.

## Hardware

1. Simpler solution for ESP32 flasher using Github pages([ESP Web Tool](https://espressif.github.io/esptool-js/))
	1. This allows you to directly flash programs via the browser without
2. Create comprehensive end user guide to setup all parts the system
	1. Include common tips like the 3D printed encoder [https://www.thingiverse.com/thing:3388595]
3. Create an enclosure around the hardware
4. ESP 32 needs more pins for IO
5. Make the pendant controller work with the new ESP driver code
6. Custom PCB design!

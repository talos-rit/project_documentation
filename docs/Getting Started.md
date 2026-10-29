# Getting Started

This guide aims to explain how to set up an existing Scorbot robotic arm with the Talos system. It's primarily aimed at users who are new to the system.

> [!hint] Robot Name Guide
> As a point of reference: You may see some of our documentation referring to "Bluey" or "Bingo"/"Unc" - these are the names of our two robots. We'll try not to use these in this guide, but a few mentions may slip through the cracks. So it's worth noting that Bluey is our turquoise Scorbot ER-4pc, and Bingo (formerly named "Unc") is our grey ER-V. Bluey is the larger of the two bots, sits on a sliding base, and has twice the encoder resolution. (If you don't know what that is, don't worry about it too much!)

> [!hint] Tip
> If you're struggling to set up Talos, a great first step might be asking an AI chatbot about us! Make sure to send links to this documentation so it knows where to look and describe your situation in as much detail as you can. Some of the answers may not be in this setup guide, but AI chatbots can browse the rest of our docs and our code to figure out the best way to help you!

## Introduction
### What is this project?
Talos is a suite of open-source tools relating to the control of Scorbots (a line of robotic arms primarily intended for use in the education sector). The goal is to create an autonomous camera operator capable of tracking subjects, but we aren't forcing anyone to use the robots for this purpose.

In fact, a major focus has been making the stack as useful to as many people as possible. For instance, we have a project called [Operator](https://github.com/talos-rit/operator) whose sole purpose is exposing a web ICD API for interacting with the robot's controller. And, separately, we have ESP32 firmware and PCB schematics for replacing the controller. So if someone doesn't have the controller, but still just wants to use serial commands to manipulate their robot, they don't *have* to use Operator at all; they can just replace the need for a Scorbot controller with the ESP32 setup.

(If you don't know what any of that means, no worries! The simple answer is that we're creating a very flexible system to serve as many people as possible.)

> [!warning] 
> WARNING: Currently, we don't actually have up-to-date ESP32 firmware and our PCB design is still WIP. A custom branch of Operator (`twinning`) exists for controlling the robot directly, and we're using a complicated system of motor hats and ADCs and so on to mock out what the final PCB will accomplish. This guide is primarily written for future use, which is why I write about its (ideal/hopeful) future state!
> 
> (TODO: remove this warning once it's no longer relevant)

This project was started in 2024 as a capstone project at the [Rochester Institute of Technology](https://www.rit.edu), and has been continued by new capstone teams in 2025 and 2026 (classes of '25 thru '27).

### Project Components

The primary project components are as follows:
- 🎛️ **[Operator](https://github.com/talos-rit/operator):** Connects to a Scorbot over serial and sends commands to move the robot, plus makes frequent telemetry polls so that the client (e.g., Commander) knows what state the robot is in.
- 🖥️ **[Commander](https://github.com/talos-rit/commander):** A web app that connects to Operator and tracks a subject by moving the robot to keep them in frame.
- 🩻 **[Pi-Vision](https://github.com/talos-rit/pi-vision):** A server that runs subject tracking (including pose estimation), all done locally on the Raspberry Pi to minimize latency.
- 📽️ **[Cam Streamer](https://github.com/talos-rit/cam_streamer):** A collection of setup scripts to create an open video stream from a webcam automatically on system startup.
- 🧰 **[ESP-Driver](https://github.com/talos-rit/esp-driver):** Firmware for running a replacement controller on an ESP32 (with accompanying hardware).

We'll discuss each of these in more depth as we set them up.

### The Two Paths
While there are many possible configurations for a Talos system, this guide will mainly focus on one distinction: **Whether or not you have a controller.** This tends to take the form of a large, heavy box full of decades-old microchips. Perhaps unsurprisingly, those might not have survived the years - either due to electrical problems or people simply throwing them away (e.g., without realizing that they went along with a robot).

We (the Talos team at RIT) happen to have one of those broken controller boxes! Which is probably for the best, since it meant we needed to support even those without a controller.

These paths will be labeled "**Original Hardware**" (for those with a controller box) and "**Custom Controller**" (for those using the custom PCB) moving forward.

## What You Need
This guide assumes you want to set up the entire Talos stack to turn your Scorbot into a robotic camera operator. If you don't care 

**Original Hardware:** Since you already have a way to control the robot, you just need:
- **Raspberry Pi**
	- Runs Operator, Cam Streamer, and Pi-Vision
	- We use Raspberry Pi Model 4B (4GB RAM) and have had good results. Your mileage may vary with older models.
		- *(TODO: run `htop` or similar and figure out how much CPU & RAM usage we have; maybe a Pi 3 would be enough? I'm not sure how much worse those were though)*
	- Of course, you'll also need an MicroSD card and some means of powering it (e.g., a USB-C cable & power adapter) - there are many starter Pi kits available online that include this stuff!
- **Webcam**
	- This guide will assume the webcam is capable of 1080p30, but you can probably get away with a lower-quality (e.g. 720p) webcam
	- This webcam will be used for subject tracking *and* the output video feed
- **Commander Device**
	- This can be nearly any computer that you can install software on
	- Runs the Commander web UI
	- While the Commander UI *could* be accessed from another device, due to browser limitations it wouldn't be feasible to allow you to run the virtual camera on another device (e.g., for broadcasting the feed over Zoom or recording it in your device's webcam capture app; for more info see [Streaming and Capturing Video](Streaming%20and%20Capturing%20Video.md))
- **USB -> Serial cable**
	- *(TODO link to one)*
	- Connects the Raspberry Pi to the controller
- **Null modem adapter**
	- *(TODO link to one)*
	- Required to talk to the controller over serial. This goes at the end of the USB -> serial cable.

**Custom Controller:** In addition to the Raspberry Pi and webcam mentioned above, you'll also need:
- The Talos PCB
	- *(TODO PCBWay link or something?)*
- ESP32
	- *(TODO: link)*
- ATX power supply?
	- TODO: is this needed with the PCB
- TODO: anything else they'd need in addition to the basics mentioned here?

## Installation
TODO installation guide for all the components, with asides for OG and custom where they differ

TODO: do we also include setup info for the raspberry pi itself? or do we just link out to the raspberry pi official guide?



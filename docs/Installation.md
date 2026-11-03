# Installation
*(TODO installation guide for all the components, with asides for OG and custom where they differ)*

## Custom Controller Setup
For those of you without an original ScorBot controller, you'll have to do some work to get the custom controller board working & ready to go.

*(TODO: can't write this until we know what this setup would actually look like)*

## Raspberry Pi Setup
*TODO: do we also include setup info for the raspberry pi itself? or do we just link out to the raspberry pi official guide?*

### Connecting to the Pi
Once your Pi has an operating system, you'll need to find a way to connect to it over the network. There are a million ways to setup a Raspberry Pi, but at the core, all of them let you perform the same option: acquire a terminal.

If you already know the Pi's LAN IP address, open a terminal on another computer and SSH into the Pi using the username and password you configured:

```bash
ssh <username>@<ip_addr>
```

For example, if you know the Pi can be accessed at `192.168.0.14` and the default user is `pi`, you'd type:

```bash
ssh pi@192.168.0.14
```

After that, you'll be prompted to enter your password. It won't show up as you type it, but once you do (then hit enter), you should have a remote terminal session open on the Pi!

If you get network errors even with a known-good IP, make sure you're on the same network and that the Pi is plugged-in. If you're still getting issues, a good troubleshooting step is to connect an Ethernet cable directly between the Pi and your computer and trying to SSH into it like that. Doing so might require knowing the device's mDNS hostname, which you've probably already configured in the initial Pi setup process. If you know that the Pi is set up to be discovered on the network as `bingo`, for instance, you'd type:

```bash
ssh pi@bingo.local
```

If this still doesn't work, you might have some other issues afoot. *(TODO: expand this section with further troubleshooting?)*

> [!question] What if I don't know the Pi's IP address?
> 
> In that case, you'll want to figure that out. Hopefully you still have some way of opening a terminal on the Pi directly, e.g. via a keyboard & HDMI cable. If so, open a terminal and enter:
> 
> ```bash
> hostname -I
> ```
> 
> The output might look something like this:
> 
> ```bash
> $ hostname -I
> 192.168.1.42 fd12:3456:789a::42
> ```
> 
> Use the first address, the one with four numbers separated by dots (the second is an IPv6 address, which we don't care about).
> 
> Now that you have the Pi's IP address, open up a terminal on another computer and follow the instructions above to SSH into it.

### Connecting Pi Hardware
*(TODO: explain the serial connection and null modem stuff for ppl who aren't using the custom board (ideally **with pictures!! using er-v as an example**). and for those people, explain the process for them. not sure where is the best place to do this, but it needs to happen at **some** point!)*

### Installing the Pi Software
We have three pieces of software that run on the Pi: **Operator, Cam Streamer, and Pi-Vision**.

#### Operator
Operator is one of the key components of the Talos project, and it's relatively simple in its function: It is a translator. Commander figures out where the camera should point and tells Operator to make adjustments to the robot to achieve that change. Operator then **translates** those adjustments into ACL commands which it then sends to the ScorBot's controller over serial.

So basically: Commander specifies a result (e.g., "Look a bit to the left"), and Operator knows how to make this specific robot achieve that result. It speaks the robot's language (quite literally), handles its quirks, and even reports back with telemetry on where the arm actually is and whether something went wrong along the way.

Now that we know what Operator does, let's get it installed on our Pi. We'll start by downloading it:

```bash
mkdir talos
cd talos
sudo apt update
sudo apt install -y git
git clone https://github.com/talos-rit/operator
```

These commands create a directory called `talos` in your user's home directory (accessible at `~/talos`), moves into that directory, installs software called `git` if it isn't installed already (it probably is!), then downloads our code into a directory called `operator`.

> [!hint] 
> So the folder structure created looks something like this:
> - `~` (user home directory)
> 	- `talos`
> 		- `operator`

Now, you need to build it. First, navigate to the folder you just downloaded, then run the build:

```bash
cd operator
cmake -B build
./make erv
```

*(TODO: does this work? i've never tried building operator before, lol. also do u have to `chmod +x ./make` beforehand? also will we keep calling the binary `erv` even when it's (relatively) robot-agnostic?)*

To run Operator, you just have to invoke the `erv` binary that just got built:
```bash
./build/bin/erv
```

This starts Operator on port `61616`. If your robot isn't connected over serial, you might get errors. If the connection works, the robot should start a homing sequence! This is a series of commands that figures out all the robots' limits. It'll happen each time you boot Operator, but it should only take a couple of seconds.

#### Cam Streamer
Time to reveal a big secret: Cam Streamer isn't actually a piece of software - at least, not in the traditional sense. It's actually a small collection of scripts that work together to set your Pi up to stream its webcam over the local network.

In particular, it has two components:
- a MediaMTX installation script
- a setup script that runs a specific FFmpeg command on Pi startup

MediaMTX generates an [RTSP (Real-Time Streaming Protocol)](https://en.wikipedia.org/wiki/Real-Time_Streaming_Protocol) server on port `8554`, and FFmpeg pushes the webcam's video onto that server. To use a metaphor: FFmpeg is the chef turning the raw ingredients (the webcam video) into a dish, and MediaMTX is the waiter who gives that dish to anyone who asks.

> [!warning] Privacy Implications
> It's worth noting that there might be privacy implications here. Anyone on the same network as you will be able to view that stream if they know the IP address. The video is *not* encrypted. However, it's also worth noting that the RTSP server **will not be reachable from the public internet** unless your network forwards port 8554 or gives the Pi a public IP. You probably have nothing to worry about, but you should keep this in mind to be safe.

Setting these up is pretty simple. Begin by downloading the repo and installing dependencies:

```bash
cd ~/talos
git clone https://github.com/talos-rit/cam_streamer
cd cam_streamer
sudo apt-get update
sudo apt-get install -y ffmpeg
```

Now you'll want to install MediaMTX:

```bash
chmod +x install-mediamtx.sh
./install-mediamtx.sh
```

> [!important] 
> Before proceeding, make sure the webcam is connected to the Pi!

Now install and start the FFmpeg publisher:

```bash
sudo ./install.sh
sudo systemctl start camera-streamer
```

That's it! You're done.

For reference, you can also run these commands to manage/troubleshoot these components:
- MediaMTX status: `sudo systemctl status mediamtx`
- Publisher status: `sudo systemctl status camera-streamer`
- Publisher logs: `sudo journalctl -u camera-streamer -f`
- Stop the publisher: `sudo systemctl stop camera-streamer`

#### Pi-Vision
Next up is Pi-Vision, which is a piece of software that typically runs on the same Raspberry Pi as Operator and Cam-Streamer. It watches the camera feed, finds a person, and nudges the arm to keep them centered.

Originally this work was performed by Commander, and (at time of writing) Commander can still do some of it, but Pi-Vision allows the robot's tracking to be more responsive since it no longer needs to wait for the video feed to get sent to another computer on the network. The video feed is still sent, but we no longer wait for that. Commander acts as something akin to a supervisor now: it turns tracking on and off, adjusts settings, and (through its debug menu) can still drive the arm manually if needed. It does these things by communicating with Pi-Vision.

*(TODO: sequence diagram might be helpful here. i think things might be a little unclear because iirc we say multiple times that commander tells operator what to do, but that's technically not true. commander tells pi-vision to tell operator what to do.)*

We'll start as before by downloading the code. This time we'll also install the requirements:

```bash
cd talos
git clone https://github.com/talos-rit/pi-vision
cd pi-vision
sudo apt install -y python3 python3-venv python3-pip git ffmpeg
python3 -m venv .venv
.venv/bin/pip install --upgrade pip
.venv/bin/pip install -r requirements.txt onnxruntime
```

This downloads the code then creates a "**v**irtual **env**ironment" wherein we can freely install things without messing anything up for the rest of the system. (& then it installs things)

Importantly, whenever running Pi-Vision in the future, you should do it inside this virtual environment (venv), otherwise you won't have access to the things installed inside of it. Luckily, you probably won't be manually running it very often, since we'll be creating a system service that'll automatically run it on login.

Do this by running the following command:

```bash
sudo nano /etc/systemd/system/pi-vision.service
```

This should put you into a text editor, so paste in the following, remembering to replace `YOUR_USERNAME_HERE` with your Pi's username:

```ini
[Unit]
Description=Talos Pi-Vision edge tracker
Wants=network-online.target
After=network-online.target mediamtx.service camera-streamer.service

[Service]
Type=simple
User=YOUR_USERNAME_HERE
WorkingDirectory=/home/YOUR_USERNAME_HERE/talos/pi-vision
Environment=PYTHONUNBUFFERED=1
Environment=POSE_MODEL=/home/YOUR_USERNAME_HERE/talos/pi-vision/yolo11n-pose-320.onnx
# BELOW ARE EXAMPLES FOR AN ER-V; ADJUST THEM FOR DIFFERENT MODELS OF SCORBOT!
Environment=MAX_OBSERVATION_AGE=1.0
Environment=MIN_PULSE=0.14
Environment=MAX_BASE_PULSE=0.36
Environment=MIN_PITCH_PULSE=0.045
Environment=MAX_PITCH_PULSE=0.070
Environment=PITCH_MIN_COUNT=-250
Environment=PITCH_MAX_COUNT=250
ExecStart=/home/YOUR_USERNAME_HERE/talos/pi-vision/.venv/bin/python /home/YOUR_USERNAME_HERE/talos/pi-vision/main.py --edge-control
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

Then press Ctrl + X and hit Enter. You should now be back into your normal terminal.

Now that you're back, enable the service and start it:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now pi-vision.service
```

You can then check that it's running properly. The following should include "active (running)":

```bash
systemctl status pi-vision.service
```

And you can follow the logs if you wish:
```bash
journalctl -u pi-vision.service -f
```

## Commander Setup
Now that the Raspberry Pi is set up, we can set up Commander, which is the GUI for running the tracking. As a reminder, this is usually going to be on a separate computer than the Raspberry Pi.

### Installing Commander
First, we'll install dependencies: Python (via [`uv`](https://docs.astral.sh/uv/getting-started/installation/)), Git, and Bun. We also install OBS Studio or `v4l2loopback` for [Streaming and Recording Video](Streaming%20and%20Recording%20Video.md) with the *Virtual Camera* feature, although doing so is technically optional (though I'd imagine there'd be little point to using Commander if you can't capture the video).

These steps will differ on each operating system. Open a new terminal and enter the following commands:

> [!info] Windows
> Install via WinGet, which is Windows's built-in package manager (similar to `apt` on Linux).
> 
> ```powershell
> winget install -e --id Git.Git
> winget install -e --id astral-sh.uv
> winget install -e --id Oven-sh.Bun
> winget install -e --id OBSProject.OBSStudio
> ```
> 
> That last one, [OBS Studio](https://obsproject.com/download), is optional, but you'll want it for [Streaming and Recording Video](Streaming%20and%20Recording%20Video.md).

> [!info] macOS
> Install the [Homebrew](https://brew.sh) package manager if it isn't already, then use it install the packages
> 
> ```bash
> /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
> brew install git uv oven-sh/bun/bun
> brew install --cask obs
> ```
> That last one, [OBS Studio](https://obsproject.com/download), is optional, but you'll want it for [Streaming and Recording Video](Streaming%20and%20Recording%20Video.md). You'll need to open it and click "**Start Virtual Camera**" in the bottom right once so that the virtual camera is registered with your system.

> [!info] Linux (Debian/Ubuntu)
> 
> Generally:
> 
> ```bash
> sudo apt update
> sudo apt install -y git curl unzip libgl1 libegl1 libxkbcommon0
> curl -LsSf https://astral.sh/uv/install.sh | sh
> curl -fsSL https://bun.sh/install | bash
> sudo apt install v4l2loopback-dkms
> sudo modprobe v4l2loopback
> ```

Then close the terminal window and start a new one.

From now on, everything should be the same regardless of platform.

First, download our code:

```shell
git clone https://github.com/talos-rit/commander.git
cd commander
```

Now install the Python packages with `uv`:

```shell
uv sync --extra yolo
```

Next, we'll build the web UI:

```bash
cd web
bun install
bun run build
cd ..
```

Finally, we'll run the web UI:

```bash
uv run commander-web
```

Now open [http://localhost:8000](http://localhost:8000) in your web browser of choice. It may take a second to load, but you should be presented with the Commander UI!

### Adding Robots to Commander
*(TODO: write up guide, with screenshots)*

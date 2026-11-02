# Burn-In Testing
We have made a simple burn-in testing script based on Commander that makes small adjustments to the robot's various axes over a period of many hours. This is to make sure the robots are fit for long-horizon demos like MakerFaire.

**At time of writing, this SHOULD NOT be run on Bluey!** It works fine on Bingo, but on Bluey it would try to jog shoulders, elbow, and wrist, none of which work. The script would see this as an arm that isn't following commands, and subsequently quit.

Right now, it will stop on an impact and it only does VERY slow, VERY small moves. That can obviously be changed to make it more "realistic" if desired.

The source code for this is in Commander on my current `feat/elbow-tracking` branch.

## Limitations
The run stops after:
- 8 hours have passed
- an impact
- a command the controller doesn't recognize
- 3 faults in a row
- radio silence from telemetry
- a move that leaves the start pose

Additionally:
- speed cannot be set above 25%
- the step sizes are capped
- the script doesn't run homing on its own

## Running the Burn-In Test
All of the following steps should be performed over SSH:

1. Stop pi-vision since it occupies an exclusive socket that the burn-in needs to talk to: `sudo systemctl stop pivision-integration`
2. Stop Commander too (if it's running anywhere and is connected directly to operator[^1]) for the same reason
3. Start Operator in another terminal (if not already running) and wait for homing to finish
4. `cd ~/burn-in`
5. `python3 -m src.burn_in --real --host localhost --moves 4`

That particular command is just going to do 4 moves to test that it works. For an actual multi-hour run, you'd use tmux to run it in the background (that way it doesn't stop if/when you lose connection):

```bash
tmux new -s burnin 'python3 -m src.burn_in --real --host localhost --hours 3 2>&1 | tee ~/burn-in/burn_in.log'
```

You can detach from the `tmux` session w/ `Ctrl-b d`, and can check back on it with `tmux attach -t burnin` or by following the log: `tail -f ~/burn-in/burn_in.log`.

To stop it early, simply run this from any SSH session:
```bash
tmux send-keys -t burnin C-c
```

Then restart pi-vision:
```bash
sudo systemctl start pivision-integration
```

To check results, view the last line of the log:
```bash
tail -n 10 burn_in.log
```

### Configuration
You can adjust the moves:

Speed is `--speed`, the step sizes are `--azimuth-deg`, `--shoulder-counts`, `--elbow-counts` and `--pitch-counts`, and the interval is `--interval` (4-30 s).

## Developing the Burn-In Script
As mentioned above, the source code lives in Commander's repo since it reuses quite a bit of Commander's code.

After making changes to the burn-in script, move it to the Pi by running the following from `commander/` on your machine:

```bash
rsync -a --exclude __pycache__ src pi@bingo.local:~/burn-in/
```

Then in SSH, run:

```bash
sudo apt install python3-loguru # install the only real dep (you don't need to do uv sync for this!)
python3 -m src.burn_in --moves 2 # sends the commands to the digital twin (robot won't move; this is a good way to verify your changes and can be done on your machine as well, obviously)
```

From there you should be good to go.

[^1]: obviously you cannot have both pi-vision *and* commander talking directly to that socket, so steps 1 and 2 are more of an "either-or" situation

# Burn-In Testing
We have made a simple burn-in testing script based on Commander that makes small adjustments to the robot's various axes over a period of many hours. This is to make sure the robots are fit for long-horizon demos like MakerFaire.

Currently I have only tested this stuff on Bingo. It might try to jog shoulders/elbows/wrist pitch that doesn't function on Bluey (at time of writing) and IDK what it would do in that case.

Right now, it will stop on an impact and it only does VERY slow, VERY small moves. That can obviously be changed to make it more "realistic" if y'all want.

The source code for this is in Commander on my current `feat/elbow-tracking` branch.

## Running the Burn-In Test
All of the following steps should be performed over SSH:

1. Stop pi-vision since it occupies an exclusive socket that the burn-in needs to talk to: `sudo systemctl stop pivision-integration`
2. Stop Commander too (if it's running anywhere and is connected to pi-vision and/or operator) for the same reason
3. Start Operator in another terminal (if not already running)
4. `cd ~/burn-in`
5. `python3 -m src.burn_in --real --host localhost --moves 4`

That particular command is just going to do 4 moves to test that it works. For an actual multi-hour run, you'd use tmux to run it in the background (that way it doesn't stop if/when you lose connection):

```bash
tmux new -s burnin 'python3 -m src.burn_in --real --host localhost --hours 3 2>&1 | tee ~/burn-in/burn_in.log'
```

To stop it early, simply run this from any SSH session:

```bash
pkill -f src.burn_in
```

## Developing the Burn-In Script
As mentioned above, the source code lives in Commander's repo since it reuses quite a bit of Commander's code.

After making changes to the burn-in script, move it to the robot:

```bash
rsync -a --exclude __pycache__ src pi@bingo.local:~/burn-in/
```


# Updating
Once you've got all of Talos installed on your system, you may on occasion want to update. This guide will explain how to update each of Talos's components.

## Commander
*(TODO - should just be this, but add copy/info:)*

```bash
cd commander
git pull
uv sync --extra yolo
cd web
bun install
bun run build
cd ..
```

## Operator
TODO

## Cam Streamer
*(TODO: this one is probably the most involved, right? since you have to reinstall the service (i think)?)*

## Pi-Vision
*(TODO - should just be this:)*

```bash
cd ~/talos/pi-vision && git pull
.venv/bin/pip install -r requirements.txt onnxruntime
sudo systemctl restart pi-vision.service
```

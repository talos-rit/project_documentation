
> [!NOTE]
> 1. For mac users please install the driver from [here](https://www.wch.cn/downloads/CH34XSER_MAC_ZIP.html) for the specific esp-s3 hardware we used. Enable the driver via the Mac's Setting > General > Items & extensions > CH34...
> 2. run this installation command: `brew install libgcrypt glib pixman sdl2 libslirp dfu-util cmake python`

## VSCode Setup

1. Clone repository
2. Enable [esp idf extension](https://marketplace.visualstudio.com/items?itemName=espressif.esp-idf-extension)
3. Install the ESP IDF (command palette: `ESP-IDF: Open ESP-IDF Installation Manager`)
	1. "start Installation" > "Start Easy Installation"
4. Connect the esp32 to your device
5. Select port (command palette: `ESP-IDF: Select Monitor Port to Use (COM, tty, usbserial)`)
6. Set target (command palette: `ESP-IDF: Set Espressif Device Target`)
7. Build project (command palette: `ESP-IDF: Build Your Project`)
	1. or click the wrench icon at the status bar
8. Flash project via UART (command palette: `ESP-IDF: Flash (UART) Your Project`)
	1. or click the lightening icon at the status bar

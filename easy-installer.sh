#!/bin/bash

git clone https://github.com/frchocolate/cclip.git
cd cclip

set -e

# Dependencies
DEPENDENCIES=("go" "rofi" "wl-copy" "wl-paste")

# Check dependencies
for dep in "${DEPENDENCIES[@]}"; do
  if ! command -v "$dep" &> /dev/null; then
    echo "Error: $dep is not installed."
    exit 1
  fi
done

# Build the project
echo "Building the project..."
if ! go build -o cclip; then
  echo "Build failed."
  exit 1
fi

# Move to /usr/bin (or /usr/local/bin for macOS)
TARGET_PATH="/usr/bin/cclip"
if [[ "$(uname)" == "Darwin" ]]; then
  TARGET_PATH="/usr/local/bin/cclip"
fi

# Move the binary
sudo mv cclip "$TARGET_PATH"

# Verify installation
if command -v cclip &> /dev/null; then
  echo "Installation successful!"
else
  echo "Installation failed."
fi

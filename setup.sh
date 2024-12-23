#!/bin/bash

# Variables
REPO_URL="https://github.com/Duskmate/hyperledger-fabric-setup-files.git"
TARGET_DIR="$(pwd)/hyperledger-network"
EXPORT_PATH="$(pwd)/fabric-samples/bin"

# Clone the repository
echo "Cloning repository from $REPO_URL..."
git clone $REPO_URL $TARGET_DIR

# Check if clone was successful
if [ $? -ne 0 ]; then
    echo "Failed to clone repository. Exiting..."
    exit 1
fi

# Create bin directory if not exists
mkdir -p $EXPORT_PATH

# Move executable files (if any) to bin directory
find $TARGET_DIR -type f -executable -exec mv {} $EXPORT_PATH \;

# Set export path in .bashrc or .zshrc
if [ -f "$HOME/.bashrc" ]; then
    SHELL_RC="$HOME/.bashrc"
elif [ -f "$HOME/.zshrc" ]; then
    SHELL_RC="$HOME/.zshrc"
else
    SHELL_RC="$HOME/.bashrc"
fi

# Update PATH with absolute path
ABSOLUTE_EXPORT_PATH="$(pwd)/fabric-samples/bin"
if ! grep -q "export PATH=$ABSOLUTE_EXPORT_PATH:\$PATH" $SHELL_RC; then
    echo "export PATH=$ABSOLUTE_EXPORT_PATH:\$PATH" >> $SHELL_RC
    echo "Path updated in $SHELL_RC"
else
    echo "Path already set in $SHELL_RC"
fi

# Source the updated .bashrc or .zshrc
source $SHELL_RC

echo "Setup complete. Files downloaded and path exported."

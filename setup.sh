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

echo "Setup complete. Files downloaded."

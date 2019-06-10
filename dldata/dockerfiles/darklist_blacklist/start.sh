#!/bin/bash

# Change GOPATH
export GOPATH=/go

# Change working directory to GOPATH
cd $GOPATH/src/github.com/h0lysp4nk/darklist/blacklist

# Check if a previous build of Blacklist exists
if [ -e "/bootstrap/blacklist" ]; then
    echo "[!] Removing old build of Blacklist..."
    rm -rf /bootstrap/blacklist

    if [ $? -eq 0 ]; then
        echo "[!] Success! Removed old build of Blacklist."
    else
        echo "[!] Error! Couldn't remove old build of Blacklist."
        exit 1
    fi
fi

# Clean the Golang workspace
echo "[!] Cleaning the Golang workspace..."
go clean
if [ $? -eq 0 ]; then
    echo "[!] Success! Workspace has been cleaned."
else
    echo "[!] Error! Workspace couldn't be cleaned."
    exit 1
fi

# Compile the Golang application
go build -o /bootstrap/blacklist ./
if [ $? -eq 0 ]; then
    echo "[!] Success! Golang Application has been compiled."
else
    echo "[!] Error! Golang Application couldn't be compiled."
    exit 1
fi

# Run the application
/bootstrap/blacklist
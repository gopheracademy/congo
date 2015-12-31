#!/bin/bash

# Simple script to help bootstrap UI dev

# Check npm is installed
command -v npm >/dev/null 2>&1 || { echo >&2 "Please install node and npm.  Aborting."; exit 1; }

# If so setup node modules
npm install
if [ $? -ne 0 ]; then
    exit 1
fi

# And compile assets
npm run build

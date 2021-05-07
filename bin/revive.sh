#!/bin/bash
cd ./cache

export REVIVE_VERSION=$(curl --silent "https://api.github.com/repos/mgechev/revive/releases/latest" | jq '.tag_name' | sed -E 's/v//g' | sed -E 's/.*"([^"]+)".*/\1/')
echo "Installing Revive $REVIVE_VERSION"

curl -sL https://github.com/mgechev/revive/releases/download/v{$REVIVE_VERSION}/revive_{$REVIVE_VERSION}_Linux_x86_64.tar.gz | tar xz

chmod +x ./revive

#!/bin/bash

# currently only for macos
rm -rf ./dist

mkdir -p dist/cntr.app/Contents/MacOS/
mkdir -p dist/cntr.app/Contents/Resources/

cp build-assets/macos/Info.plist dist/cntr.app/
cp build-assets/macos/cntr.icns dist/cntr.app/Contents/Resources/

go build -o dist/cntr.app/Contents/MacOS/cntr .
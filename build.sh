#!/bin/bash

# currently only for macos
rm -rf ./dist

mkdir -p dist/cntr.app/Contents/MacOS/
mkdir -p dist/cntr.app/Contents/Resources/

cp build-assets/macos/Info.plist dist/cntr.app/Contents/
cp build-assets/macos/cntr.icns dist/cntr.app/Contents/Resources/

cd app 
npm run build 
cd ..

packr2
go build -o dist/cntr.app/Contents/MacOS/cntr .
packr2 clean
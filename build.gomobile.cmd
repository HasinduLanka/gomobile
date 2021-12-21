#!/bin/sh

echo "Building GoMobile..."

rm -rf ./FlutterSrc/spindly_app/android/app/src/main/libs/GoApp.aar
mkdir -p ./FlutterSrc/spindly_app/android/app/src/main/libs/

gomobile bind -o ./FlutterSrc/spindly_app/android/app/src/main/libs/gonativelib.aar

#!/bin/bash

## Simple script to cross-compile the project

# Build for linux x64
build_linux_x64() {
    echo "Building for Linux x64"
    mkdir -p build/linux/x64
    GOOS=linux GOARCH=amd64 go build -o build/linux/x64/xc
}

# Build for linux x86
build_linux_x86() {
    echo "Building for Linux x86"
    mkdir -p build/linux/x86
    GOOS=linux GOARCH=386 go build -o build/linux/x86/xc
}

# Build for windows x64
build_windows_x64() {
    echo "Building for Windows x64"
    mkdir -p build/windows/x64
    GOOS=windows GOARCH=amd64 go build -o build/windows/x64/xc.exe
}

# Build for windows x86
build_windows_x86() {
    echo "Building for Windows x86"
    mkdir -p build/windows/x86
    GOOS=windows GOARCH=386 go build -o build/windows/x86/xc.exe
}

# Build for darwin x64
build_darwin_x64() {
    echo "Building for Darwin x64"
    mkdir -p build/darwin/x64
    GOOS=darwin GOARCH=amd64 go build -o build/darwin/x64/xc
}

# Build for darwin x86
build_darwin_x86() {
    echo "Building for Darwin x86"
    mkdir -p build/darwin/x86
    GOOS=darwin GOARCH=386 go build -o build/darwin/x86/xc
}

# Build for linux arm
build_linux_arm() {
    echo "Building for Linux ARM"
    mkdir -p build/linux/arm
    GOOS=linux GOARCH=arm go build -o build/linux/arm/xc
}

# Build for linux arm64
build_linux_arm64() {
    echo "Building for Linux ARM64"
    mkdir -p build/linux/arm64
    GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/xc
}

# Build for android arm
build_android_arm() {
    echo "Building for Android ARM"
    mkdir -p build/android/arm
    GOOS=android GOARCH=arm go build -o build/android/arm/xc
}

# Build for android arm64
build_android_arm64() {
    echo "Building for Android ARM64"
    mkdir -p build/android/arm64
    GOOS=android GOARCH=arm64 go build -o build/android/arm64/xc
}

# Build for android x86
build_android_x86() {
    echo "Building for Android x86"
    mkdir -p build/android/x86
    GOOS=android GOARCH=386 go build -o build/android/x86/xc
}

# Build for android x64
build_android_x64() {
    echo "Building for Android x64"
    mkdir -p build/android/x64
    GOOS=android GOARCH=amd64 go build -o build/android/x64/xc
}

# Build for all platforms
build_all() {
    build_linux_x64
    # build_linux_x86
    # build_windows_x64
    # build_windows_x86
    # build_darwin_x64
    # build_darwin_x86
    # build_linux_arm
    # build_linux_arm64
    # build_android_arm
    # build_android_arm64
    # build_android_x86
    # build_android_x64
}

# Build for the current platform
build_current() {
    echo "Building for current platform"
    go build -o build/xc
}

# Build for specific platform
build_specific() {
    if "$1" == "linux" && "$2" == "x64"; then
        build_linux_x64
    elif "$1" == "linux" && "$2" == "x86"; then
        build_linux_x86
    elif "$1" == "windows" && "$2" == "x64"; then
        build_windows_x64
    elif "$1" == "windows" && "$2" == "x86"; then
        build_windows_x86
    elif "$1" == "darwin" && "$2" == "x64"; then
        build_darwin_x64
    elif "$1" == "darwin" && "$2" == "x86"; then
        build_darwin_x86
    elif "$1" == "linux" && "$2" == "arm"; then
        build_linux_arm
    elif "$1" == "linux" && "$2" == "arm64"; then
        build_linux_arm64
    elif "$1" == "android" && "$2" == "arm"; then
        build_android_arm
    elif "$1" == "android" && "$2" == "arm64"; then
        build_android_arm64
    elif "$1" == "android" && "$2" == "x86"; then
        build_android_x86
    elif "$1" == "android" && "$2" == "x64"; then
        build_android_x64
    else
        echo "Invalid platform"
    fi
}

# Check if the user has specified a platform
if [ $# -eq 0 ]; then
    build_current
elif [ $# -eq 1 ]; then
    if [ "$1" == "all" ]; then
        build_all
    else
        echo "Invalid platform"
    fi
elif [ $# -eq 2 ]; then
    build_specific "$1" "$2"
else
    echo "Invalid number of arguments"
fi

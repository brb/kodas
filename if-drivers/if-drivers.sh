#!/bin/sh

# List underlying drivers of network interfaces
# Taken from: http://unix.stackexchange.com/questions/41817/linux-how-to-find-the-device-driver-used-for-a-device

set -eu

for f in /sys/class/net/*; do
    dev=$(basename $f)
    driver=$(readlink $f/device/driver/module)
    if [ $driver ]; then
        driver=$(basename $driver)
    fi
    addr=$(cat $f/address)
    operstate=$(cat $f/operstate)
    printf "%10s [%s]: %10s (%s)\n" "$dev" "$addr" "$driver" "$operstate"
done

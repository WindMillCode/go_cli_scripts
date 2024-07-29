#!/bin/zsh

# Get running processes
processes=$(ps -axo pid,comm,tty,rss | tail -n +2)

# Format the output similar to tasklist
echo "PID        Name                     TTY        Mem Usage"
echo "$processes" | awk '{printf "%-10s %-25s %-10s %-15s\n", $1, $2, $3, $4}'


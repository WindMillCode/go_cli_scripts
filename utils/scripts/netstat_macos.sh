#!/bin/zsh

# Get TCP connections
tcp_connections=$(netstat -an -p tcp | awk 'NR>2 {print "TCP", $4, $5, $6, $9}')

# Get UDP listeners
udp_connections=$(netstat -an -p udp | awk 'NR>2 {print "UDP", $4, "*", "*", $7}')

# Combine TCP and UDP connections
all_connections=$(echo -e "$tcp_connections\n$udp_connections")

# Format the output similar to netstat -ano with more space between column headers
echo "Proto       Local Address            Foreign Address          State           PID"
echo "$all_connections" | awk '{printf "%-10s %-25s %-25s %-15s %s\n", $1, $2, $3, $4, $5}'

#!/bin/zsh

# Get TCP connections
tcp_connections=$(lsof -iTCP -sTCP:LISTEN,ESTABLISHED -nP | awk 'NR>1 {
  local_address = $9;
  foreign_address = $10;
  state = "ESTABLISHED";
  if (index(local_address, "->") > 0) {
    split(local_address, addresses, "->");
    local_address = addresses[1];
    foreign_address = addresses[2];
  }
  else if (!match(foreign_address, /^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+:[0-9]+$/) && !match(foreign_address, /\*:[0-9]+$/)) {
    foreign_address = "*";
    state = "LISTEN";
  }
  printf "TCP %-25s %-25s %-15s %s\n", local_address, foreign_address, state, $2
}')

# Get UDP listeners
udp_connections=$(lsof -iUDP -nP | awk 'NR>1 {
  local_address = $9;
  foreign_address = $10;
  state = "*";
  if (index(local_address, "->") > 0) {
    split(local_address, addresses, "->");
    local_address = addresses[1];
    foreign_address = addresses[2];
  }
  else if (!match(foreign_address, /^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+:[0-9]+$/) && !match(foreign_address, /\*:[0-9]+$/)) {
    foreign_address = "*";
  }
  printf "UDP %-25s %-25s %-15s %s\n", local_address, foreign_address, state, $2
}')

# Combine TCP and UDP connections
all_connections=$(echo "$tcp_connections\n$udp_connections")

# Format the output similar to netstat -ano with more space between column headers
echo "Proto   Local Address   Foreign Address    State    PID" | awk '{printf "%-10s %-55s %-55s %-25s %s\n", $1, $2" "$3, $4" "$5, $6, $7}'
echo "$all_connections" | awk '{printf "%-10s %-55s %-55s %-25s %s\n", $1, $2, $3, $4, $5}'

# Get TCP connections
$tcpConnections = Get-NetTCPConnection | Select-Object @{Name='Proto'; Expression={'TCP'}}, `
                                               @{Name='Local Address'; Expression={$_.LocalAddress}}, `
                                               @{Name='Local Port'; Expression={$_.LocalPort}}, `
                                               @{Name='Foreign Address'; Expression={$_.RemoteAddress}}, `
                                               @{Name='Foreign Port'; Expression={$_.RemotePort}}, `
                                               @{Name='State'; Expression={$_.State}}, `
                                               @{Name='PID'; Expression={$_.OwningProcess}}

# Get UDP listeners
$udpConnections = Get-NetUDPEndpoint | Select-Object @{Name='Proto'; Expression={'UDP'}}, `
                                             @{Name='Local Address'; Expression={$_.LocalAddress}}, `
                                             @{Name='Local Port'; Expression={$_.LocalPort}}, `
                                             @{Name='Foreign Address'; Expression={'*'}}, `
                                             @{Name='Foreign Port'; Expression={'*'}}, `
                                             @{Name='State'; Expression={'*'}}, `
                                             @{Name='PID'; Expression={$_.OwningProcess}}

# Combine TCP and UDP connections
$allConnections = $tcpConnections + $udpConnections

# Format the output similar to netstat -ano with more space between column headers
$allConnections | Format-Table -Property  @{Label="PID"; Expression={"$($_.PID)"}; Width=10},
                                          @{Label="Proto"; Expression={"$($_.Proto)"};Width=10},
                                          @{Label="Local Address"; Expression={"$($_.'Local Address'):$($_.'Local Port')"}; Width=25},
                                          @{Label="Foreign Address"; Expression={"$($_.'Foreign Address'):$($_.'Foreign Port')"}; Width=25},
                                          @{Label="State"; Expression={"$($_.State)"}; Width=15}


# Get running processes
$processes = Get-Process | Select-Object @{Name='Name'; Expression={$_.ProcessName}}, `
                                       @{Name='PID'; Expression={$_.Id}}, `
                                       @{Name='Session#'; Expression={($_ | Get-Process | Select-Object -ExpandProperty SessionId)}}, `
                                       @{Name='Mem Usage'; Expression={"{0:N0} K" -f ($_.WS / 1KB)}}


Format the output similar to tasklist
$processes | Format-Table -Property  @{Label="PID"; Expression={"$($_.PID)"}; Width=10},
                                     @{Label="Name"; Expression={"$($_.Name)"}; Width=25},
                                     @{Label="Session#"; Expression={"$($_.'Session#')"}; Width=10},
                                     @{Label="Mem Usage"; Expression={"$($_.'Mem Usage')"}; Width=15}

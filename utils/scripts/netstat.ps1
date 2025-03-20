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

# SIG # Begin signature block
# MIIFsgYJKoZIhvcNAQcCoIIFozCCBZ8CAQExCzAJBgUrDgMCGgUAMGkGCisGAQQB
# gjcCAQSgWzBZMDQGCisGAQQBgjcCAR4wJgIDAQAABBAfzDtgWUsITrck0sYpfvNR
# AgEAAgEAAgEAAgEAAgEAMCEwCQYFKw4DAhoFAAQUEIGmPdshGDfBvZtFXIbhyIg+
# dgWgggM2MIIDMjCCAhqgAwIBAgIQe5b4YokbOLZF9pP7k+K9BTANBgkqhkiG9w0B
# AQsFADAxMS8wLQYDVQQDDCZXaW5kbWlsbGNvZGUgUG93ZXJTaGVsbCBTY3JpcHQg
# U2lnbmluZzAeFw0yNTAzMTMwNTA3MThaFw0yNjAzMTMwNTI3MThaMDExLzAtBgNV
# BAMMJldpbmRtaWxsY29kZSBQb3dlclNoZWxsIFNjcmlwdCBTaWduaW5nMIIBIjAN
# BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuQE/dHljEczqSG18tmX5czKd/E0n
# EhK0zU6YetojlNacKnwZMRQ6YInXJIAAW8vyXI1M3S3KS+LvFuUrImbcTSu6Qeum
# Dy6jS/c7av/Y+1q/7nmsrc9ughY7/9gQlGc5XtJTHs7dIGOCSnnSMH+o+3EgYVdX
# WabtFqK/kNcWOLQvlXDgvXjIGo4dwIM+iw3y6+i2AWDEWSjQYC9X9kukJ+ESuQlW
# Z6fM0K96Zx7wKwWZiGd4NkXQNb9SUuSJpV5fqWwf9FceTW3F3lAzV90kZpZaiz3V
# ekve0Aa/XoJCimxoTPEDatBO09i14q2I3VcQ3kBX73xGFMil3H4h0HbhUQIDAQAB
# o0YwRDAOBgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwMwHQYDVR0O
# BBYEFNY/YuVqstrXH09kkfMFUyeqH07jMA0GCSqGSIb3DQEBCwUAA4IBAQBIeVeT
# g7QZqIoSpyesNyWXri2iQR/jJ4+0iAK0YGF9Ia8D0WYreEs6zi2LRdPBBMoKV40J
# 1tvj5KxrmfSI1slDtmG5IVhDa5fAmixtoEW2JOR6y5erYZh8lZguWZoL/ZguFIIH
# jQp8Tb+OmR4k4j1v1S0+QHqdP7H36Ua8U0a99zSlKOEbzdr3RfWwlCyVOGKHHRW3
# eI8A43J4RldY+ZV1eXjcqb2cc3tdGvAnvkaBKaCmRdVzwKtEbh6F8zDfacAzBEwY
# XR9ZSI0aUgCKd1aEtP8QMhBCO9pkEUzezgQVSJ/Wl79p4CCAh8yAXmrE3wA1T0Fv
# l/EEhRjjI2oy44ZiMYIB5jCCAeICAQEwRTAxMS8wLQYDVQQDDCZXaW5kbWlsbGNv
# ZGUgUG93ZXJTaGVsbCBTY3JpcHQgU2lnbmluZwIQe5b4YokbOLZF9pP7k+K9BTAJ
# BgUrDgMCGgUAoHgwGAYKKwYBBAGCNwIBDDEKMAigAoAAoQKAADAZBgkqhkiG9w0B
# CQMxDAYKKwYBBAGCNwIBBDAcBgorBgEEAYI3AgELMQ4wDAYKKwYBBAGCNwIBFTAj
# BgkqhkiG9w0BCQQxFgQUEiJy6QIbAvWc44xoQIBlC15cwwwwDQYJKoZIhvcNAQEB
# BQAEggEAZEOmxoVo5hxbnuEYyo0z+M1ISFPJyfG9V2UXWFPBfwVLSM6yM838fSAc
# ClZAGaGPMHq+HTSy3d2eSOxFHSIWq9DZvdFuk3ysXTHxd6azNZSBbkj6WXR2jmsd
# SHn4/LCMSSuU+FM/HEy0FcZyQd+OSFVP9VSIIpINBcZLuuiQv0zdF0awKJ+oKq3g
# DHhmkPjSFalSNIGRk6BaAP8HOr9sKntoQIZTwMNN3vDCt2RTykaBn77LiYkCauWo
# l/3gRS+kjMG1MEAji6trboYl+JFb//zV5oni0GMotil3CT2e6BYdt59uYBHvFdEM
# wwAp6vCK2i1tJYvoyFHhcvbnU7r6cQ==
# SIG # End signature block

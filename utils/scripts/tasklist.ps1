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

# SIG # Begin signature block
# MIIFsgYJKoZIhvcNAQcCoIIFozCCBZ8CAQExCzAJBgUrDgMCGgUAMGkGCisGAQQB
# gjcCAQSgWzBZMDQGCisGAQQBgjcCAR4wJgIDAQAABBAfzDtgWUsITrck0sYpfvNR
# AgEAAgEAAgEAAgEAAgEAMCEwCQYFKw4DAhoFAAQUOh1RQN7Kcr2CUTGCznfZJR7i
# /pugggM2MIIDMjCCAhqgAwIBAgIQe5b4YokbOLZF9pP7k+K9BTANBgkqhkiG9w0B
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
# BgkqhkiG9w0BCQQxFgQUtnC7huQVRRo+X88L83IKPyO6JCMwDQYJKoZIhvcNAQEB
# BQAEggEAgzPzZ7iwuOtXiZS2BTuotmgDQh/i0ELrUVpE/8AW55QukWlmmqaX+1gd
# fYUo8zH2Yvii1vS1xxbjmPzyzpj9holb0P+PjEBRjOJ2q+GYMxwpYu4E95WD9Ikq
# /Zny+1D4UkiX155DbWYJFj9i9h41PMMDXRGsqJY6i2oHak+lltaMJVkEFAZw+bDm
# Z8C2CD6O3SBYsQ/Vuq6U2n1gECR3HS4Q5EcdtHk2Wsz84UILgURp7AAgG5WGX8n+
# KZfJVlfrvFPPAH27TKeW1zUgZxTpYSetAkRBuyBZy30IR4BL40AOfmLoBat/SQXe
# D4S7V4LhflISCB7Am3p7U9D+5C3+9Q==
# SIG # End signature block

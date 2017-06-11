# Access control app written in golang for Raspberry PI
*consist of a scanner reader and udp server listener*
  * The scanner scans an image and send the raw ASCII to an access control server for a check
  * The opener is an UDP server listener which will trigger a pin output based on a string received on the socket

## Install binaries so that the systemd service manager runs them as a service
	GOBIN=/usr/local/bin go install opener.go
	GOBIN=/usr/local/bin go install scanner.go
## Restart the services to run the latest installed files
```
systemctl restart opener
systemctl restart scanner
```

## Service config files
```
/lib/systemd/system/scanner.service

[Unit]
  Description=Scanner Reader

  [Service]
  ExecStart=/usr/local/bin/scanner 127.0.0.1:2002
  ExecStartPre=/bin/sleep 5
  Restart=always
  RestartSec=10

  [Install]
  WantedBy=multi-user.target

/lib/systemd/system/opener.service

[Unit]
  Description=Door Control listener

  [Service]
  ExecStart=/usr/local/bin/opener 2001
  Restart=always
  RestartSec=1

  [Install]
  WantedBy=multi-user.target
```
## Service start on boot
```
systemctl enable opener
systemctl enable scanner

```


## Check the service logs

```
journalctl -u scanner -u opener -f --since "2017-06-01 17:15:00"
-u unit name(service) , -f follow
```
*log rotating and log size is all handled automaticaly by journald*

## Update dependancy packages
all dependant packages are vendored in the vendor folder using
`dep init`

to update a package to the latest version
`dep ensure github.com/krasi-georgiev/rpiGpio@^0.8.0`

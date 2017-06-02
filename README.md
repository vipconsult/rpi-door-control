	*** Install binaries so that the systemd service manager runs them as a service
		GOBIN=/usr/local/bin go install opener.go
		GOBIN=/usr/local/bin go install scanner.go
	*** Restart the services to run the latest installed files
		systemctl restart opener
		systemctl restart scanner

	*** Service config files
		/lib/systemd/system/scanner.service
		/lib/systemd/system/opener.service

	*** Check the service logs
		journalctl -u scanner -u opener -f --since "2017-06-01 17:15:00"
			-u unit name(service) , -f follow

		log rotating and space management is all handled automaticaly by journald :)

	*** Update dependancy packages
		all dependant packages are vendored in the vendor folder using
		`dep init`
		to update a package to the latest version
		`dep ensure github.com/krasi-georgiev/rpiGpio@^0.8.0`

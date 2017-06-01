// TO DO

- remove the timeout in the read loop

// Troubleshoot connection loss

// Think how to work together on the same files

// setup systemd service and logging filepath
// setup logrotate

// Health check
	Check connection with the reader and restart the pi on failed
		no way to test the reader is working- if no reading is taken in 3h , restart the usb ports?
	Check that the golang server reply handler can accept request

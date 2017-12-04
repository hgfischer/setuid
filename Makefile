all:
	go build
	sudo chown root: setuid
	sudo chmod 4755 setuid
	./setuid
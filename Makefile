all:
	go build
	sudo chmod 4755 setuid
	sudo chown root: setuid
	./setuid
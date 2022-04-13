all:
	go build ./src/main.go
	mv main dict

install:
	install -Dt /usr/local/bin -m 755 dict

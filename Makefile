all:
	go build -o dict ./src/main.go

clean:
	@if [ -f dict ] && [ -x dict ]; then rm dict; fi

install:
	install -Dt /usr/local/bin -m 755 dict

uninstall:
	rm /usr/local/bin/dict

all:
	go build -o build/dict ./main.go

clean:
	@if [ -f build ] && [ -x build ]; then rm build; fi

install:
	install -Dt /usr/local/bin -m 755 build/dict

uninstall:
	rm /usr/local/bin/dict

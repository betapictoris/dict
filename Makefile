all:
	mkdir -vp build
	make binary
	make deb

binary:
	go build -o build/dict ./main.go

deb:
	mkdir -vp build/dict-deb/usr/local/bin/
	go build -o build/dict-deb/usr/local/bin/
	mkdir -vp build/dict-deb/DEBIAN/
	cp -v dict/deb/control build/dict-deb/DEBIAN/
	cp -v dict/deb/postinst build/dict-deb/DEBIAN/
	chmod 775 build/dict-deb/DEBIAN/postinst
	dpkg-deb --build build/dict-deb

clean:
	rm -rv build

install:
	install -vDt /usr/local/bin -m 755 build/dict

uninstall:
	rm -v /usr/local/bin/dict

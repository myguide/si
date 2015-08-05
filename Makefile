WREN_DIR=src/wren
BUILD_DIR=build
DEPS_FILE=vendor/.deps
APPLICATION=si
INSTALL_PREFIX=/usr/local

all: si

si: deps maps wren bindata
	mkdir -p build
	cd src && go build -x -o ../${BUILD_DIR}/${APPLICATION} -v
	rm -f src/*.c
	rm -f src/*.h

# Build wren
wren:
	make -C ${WREN_DIR}

# Update and pull submodules
deps: vendor/
	@if [ -a ${DEPS_FILE} ] ; \
	then \
		echo "Skipping dependencies - run 'make clean' first if needed" ; \
	else \
		echo "Updating dependencies..." ; \
		git submodule init ; \
		git submodule update ; \
		git submodule foreach git pull origin master ; \
		go get -u github.com/myguide/cli ; \
		go get -u github.com/myguide/fsnotify ; \
		touch ${DEPS_FILE} ; \
	fi;

maps: maps/
	cp src/maps/* src/

# Generate API bindata
bindata: src/api
	go-bindata -o $^.go $^

test:
	go test ./src/classes

clean:
	rm -f ${BUILD_DIR}/${APPLICATION}
	cd src/wren && make clean
	rm -rf vendor/pkg vendor/src
	rm -f ${DEPS_FILE}
	rm -f src/*.c
	rm -f src/*.h

.PHONY: install
install: ${BUILD_DIR}/${APPLICATION}
	@install -m 0755 $^ ${INSTALL_PREFIX}/bin
	@echo "si has been successfully installed!"
	@echo "Execute 'si --version' to verify."

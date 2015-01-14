WREN_DIR=src/wren
BUILD_DIR=build
DEPS_FILE=vendor/.deps
APPLICATION=wrengo
INSTALL_PREFIX=/usr/local

all: wrengo

wrengo: deps wren bindata
	mkdir -p build
	cd src && go build -o ../${BUILD_DIR}/${APPLICATION} -v

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
		touch ${DEPS_FILE} ; \
	fi;

# Generate API bindata
bindata: src/api
	go-bindata -o $^.go $^

clean:
	rm -f ${BUILD_DIR}/${APPLICATION}
	cd src/wren && make clean
	rm -f ${DEPS_FILE}

.PHONY: install
install: ${BUILD_DIR}/${APPLICATION}
	@install -m 0755 $^ ${INSTALL_PREFIX}/bin
	@echo "Wrengo has been successfully installed!"
	@echo "Execute 'wrengo --version' to verify."

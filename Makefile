SHELL                  = /bin/bash

PREFIX                 = $(shell pwd)/install

static_flags           = --enable-shared=no --enable-static=yes

ffmpeg_version         = 4.1
ffmpeg_sha256          = b684fb43244a5c4caae652af9022ed5d85ce15210835bce054a33fb26033a1a5

ifeq ($(OS), Windows_NT)
	OSName = windows
else
	OSName = $(shell echo $(shell uname -s | tr '[:upper:]' '[:lower:]'))
	OSArch = $(shell echo $(shell uname -m | tr '[:upper:]' '[:lower:]'))
	ifeq ($(OSArch), armv7l)
		OSName = arm
	endif
endif

$(OSName): ffmpeg
	if [ -d $(PREFIX)/$@-lib ]; then \
		$(RM) -r $(PREFIX)/$@-lib; \
	fi
	cd $(PREFIX); mv lib $@-lib

ffmpeg:
	if [ -f $@-$($@_version).tar.bz2 ]; then \
		if [ $(firstword $(shell sha256sum $@-$($@_version).tar.bz2)) != $($@_sha256) ]; then \
			$(RM) $@-$($@_version).tar.bz2; \
			make $@; \
		fi; \
	else \
		wget https://ffmpeg.org/releases/$@-$($@_version).tar.bz2; \
	fi
	if [ ! -d $@-$($@_version) ]; then \
		tar xjvf $@-$($@_version).tar.bz2; \
	fi
	if [ -f $@-$($@_version)/Makefile ]; then \
		cd $@-$($@_version); make clean; \
	fi
	cd $@-$($@_version); ./configure --prefix=$(PREFIX) \
	--enable-gpl --enable-version3 --enable-nonfree --disable-programs \
	--disable-doc --disable-autodetect; \
	make -j8; make install

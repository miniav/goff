.PHONY: all
all:
	gcc -o decode.test -L../drivers/Linux-lib -I../drivers/include decode.c -lavformat -lavcodec -lavutil -lswresample -lswscale -Wno-deprecated-declarations
	strip decode.test
	$(RM) *.raw
	./decode.test ./test.ogg

.PHONY: resample
resample:
	gcc -o resample.test -L../drivers/Linux-lib -I../drivers/include resample.c -lavformat -lavcodec -lavutil -lswresample -lswscale -Wno-deprecated-declarations

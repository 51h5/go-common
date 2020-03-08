TARGET := libcrsa.a
HEADER_PATH := -I/usr/local/Cellar/openssl@1.1/1.1.1d/include

.PHONY: clean
clean:
	rm -f *.o *.so *.a $(TARGET)

.PHONY: build
build: clean
	$(CC)  -o rsa.o -c rsa.cpp $(HEADER_PATH)
	ar r $(TARGET) rsa.o

.PHONY: run
run: build
	$(CC) rsa.o -lcrypto -lssl  -lstdc++
	./a.out
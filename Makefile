install:
	go test
	go install

test: install
	audiovisio -c test/test.yaml -o test/test.html
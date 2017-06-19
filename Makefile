test:
	bash -c 'rm -f coverage.txt && cd src && ls -d */ | while read dir; do go test -coverprofile=$${dir:0:-1}.cover.out -covermode=atomic ./$$dir; done && cat *.cover.out >> ../coverage.txt && rm *cover.out'
cover:
	bash -c 'make test'
	bash -c 'go tool cover -html=coverage.txt -o coverage.html'
build:
	bash -c 'cd src && go build ./...'
install:
	glide install
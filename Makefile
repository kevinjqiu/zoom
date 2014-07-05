zoom:
	go build -o zoom github.com/kevinjqiu/zoom/cli
clean:
	rm zoom
serve: zoom
	./zoom serve
get-deps:
	go get github.com/codegangsta/cli
	go get github.com/gorilla/mux
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega
	go get code.google.com/p/go.tools/cmd/cover
	go get github.com/mattn/goveralls

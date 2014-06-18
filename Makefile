zoom:
	go build -o zoom github.com/kevinjqiu/zoom/cli
clean:
	rm zoom
serve: zoom
	./zoom serve

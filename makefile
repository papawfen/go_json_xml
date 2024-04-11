all: clean ex00 ex01 ex02

ex02:
	@go build -o compareFS ex02/*.go

ex01:
	@go build -o compareDB ex01/*.go

ex00: 
	@go build -o readDB ex00/*.go

clean:
	@rm -rf readDB compareDB compareFS
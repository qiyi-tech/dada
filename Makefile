# 执行单测并生成单测覆盖率
test:
	go test -v -cover -coverprofile=cover.out -covermode=count

race:
	go test -v -race

# 通过命令行查看测试覆盖率
cover:
	go tool cover -func=cover.out

# 通过浏览器查看测试覆盖率查看测试覆盖率
cover-html:
	go tool cover -html=cover.out
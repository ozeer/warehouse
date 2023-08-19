## Kratos Project Demo

### 安装Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
### 创建项目
```
kratos new warehouse
```
### 创建模块服务
#### 1. 用户模块
```bash
cd warehouse
# 创建user.proto
kratos proto add api/git/user.proto
# 创建PB文件（Generate the proto code）
kratos proto client api/git/user.proto
# 生成service服务（Generate the source code of service by proto file）
kratos proto server api/git/user.proto -t internal/service
```
#### 2. 仓库模块
```bash
# 创建repo.proto
kratos proto add api/git/repo.proto
# 创建PB文件
kratos proto client api/git/repo.proto
# 生成service服务
kratos proto server api/git/repo.proto -t internal/service
```

### 启动项目
```
go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
### Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
### Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

### Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```


# 使用一個輕量級的 Go 基礎映像
FROM golang:1.23-alpine

# 設置工作目錄
WORKDIR /app

# 複製 Go 模組文件
COPY go.mod go.sum ./

# 下載依賴
RUN go mod download

# 複製所有源碼
COPY . .

# 編譯成可執行文件
RUN go build -o main .

# 暴露端口（如果你的應用需要）
EXPOSE 8080

# 啟動應用
CMD ["./main"]
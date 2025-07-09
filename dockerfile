# Ubuntu指令
#你要以誰的Image為環境，事先幫你安裝好go這個環境

# 1. 建構環境
FROM golang:1.23 


# 2. 指定工作目錄
WORKDIR /work
# /work這個可以自己取
# 用途：我這個容器的預設位置在哪裡，有這個資料夾會直接進去，沒有的話會自動創建

# 3. 複製 go.mod 和 go.sum 檔案
COPY go.mod go.mod
COPY go.sum go.sum

# 4. 下載依賴
RUN go mod download

# 5. 複製GO檔案
COPY *.go .


# 以下兩句只是為了讓我們看到過程
# run echo "Before build" && sleep 2
# run ls -al && sleep 2


# 6. 進行編譯
RUN go build -v

# image一打開時會執行的指令
# 這裡的名稱要看你編譯完成出現什麼，現在出現gogo.exe，所以就是gogo，再來去docker build -t gogo .
# gogo這個名字是看你go.mod第一行打什麼 

# 5. 執行
ENTRYPOINT ["./gogo"]  

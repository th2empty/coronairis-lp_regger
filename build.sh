echo -e "\033[36m Building app"
go build -o bin/app cmd/main.go
echo -e "\033[36m Preparing after compilation..."
chmod +x bin/app
echo -e "\033[32m Building finished"
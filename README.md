go run ./secuential_wc.go ../files/pg84.txt wc.so

go run ./coordinator.go ../files/pg84.txt wc.so

go run ./worker wc.so

go build -buildmode=plugin -o wc.so ./plugins/wc.go
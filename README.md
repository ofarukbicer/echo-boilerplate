# echo-boilerplate

echo kütüphanesini kullanırken starter pack olması için yaptığım ufak bir başlangıç template'i

## Setup

```sh
# Setup

rm -rf go.mod go.sum
go mod init $PROJECT_NAME
go mod tidy

# Note: don't forget to update the invoked web packs in all files after doing this.


# Start 

go run main.go 
# or 
go run .
```


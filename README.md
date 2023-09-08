# 此项目在 GOPATH/src 目录下

# Ethereum Development with Go

https://goethereumbook.org/zh/

https://www.runoob.com/go/go-tutorial.html

https://geth.ethereum.org/docs

https://www.topgoer.com/

## Abigen

https://geth.ethereum.org/docs/tools/abigen

`solc --abi erc20.sol -o build`

`abigen --abi IERC20.abi --pkg erc20 --type ERC20 --out erc20.go`

`solc --abi erc721.sol -o build`

`abigen --abi IERC721.abi --pkg erc721 --type ERC721 --out erc721.go`

## Request

### Post(json)

`curl http://localhost:8888/postjson1 -X POST -H "Content-Type: application/json" -d "{\"message\":\"Hi Go\"}"`

`curl http://localhost:8888/postjson2 -X POST -H "Content-Type: application/json" -d "{\"message\":\"Hi Go\"}"`

### Get

# config

go version: 1.20.7

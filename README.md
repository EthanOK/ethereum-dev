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

### MAC System

```
solcjs --abi IERC6551Registry.sol -o build

abigen --abi IERC6551Registry.abi --pkg erc6551 --type ERC6551Registry --out erc6551Registry.go

<!-- ./solidity -->
solcjs --abi Multicall3.sol -o build
<!-- ./solidity/build -->
abigen --abi Multicall3.abi --pkg multicall3 --type Multicall3 --out multicall3.go

```

## Request

### Post(json)

`curl http://localhost:8888/postjson1 -X POST -H "Content-Type: application/json" -d "{\"message\":\"Hi Go\"}"`

`curl http://localhost:8888/postjson2 -X POST -H "Content-Type: application/json" -d "{\"message\":\"Hi Go\"}"`

### Get

go version: 1.24.0

# ethereum-dev doc

## utils

## filters

### mul topics filter

```go
// filters/mulTopicsEvent.go

	filterQuery := ethereum.FilterQuery{
		BlockHash: (*common.Hash)(&blockHash),
		Topics: [][]common.Hash{
			topic0s[:], // 多个 topic0，无需每个 topic0 调用一次
		},
	}
    // Topics: [][]common.Hash
    // Examples:
	// {} or nil          matches any topic list
	// {{A}}              matches topic A in first position
	// {{}, {B}}          matches any topic in first position AND B in second position
	// {{A}, {B}}         matches topic A in first position AND B in second position
	// {{A, B}, {C, D}}   matches topic (A OR B) in first position AND (C OR D) in second position

```

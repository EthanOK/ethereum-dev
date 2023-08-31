# Ethereum Development with Go

https://goethereumbook.org/zh/

https://www.runoob.com/go/go-tutorial.html

## Abigen

https://geth.ethereum.org/docs/tools/abigen

`solc --abi erc20.sol -o build`

`abigen --abi erc20.abi --pkg erc20 --type ERC20 --out erc20.go`

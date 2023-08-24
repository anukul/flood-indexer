package types

type Order struct {
	UserAddress string
	UsdValue    float64
	TxHash      string
	BlockTime   uint64
	BlockNumber uint64
}

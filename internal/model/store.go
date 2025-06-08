package model

type StoreMerchant struct {
    ID   int32
    Name string
    Url string
}

type UpdateStoreMerchant struct {
	Name *string
	Url  *string
}

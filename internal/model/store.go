package model

type StoreMerchant struct {
	ID   int32
	Name string
	Url  string
}

type MerchantPatch struct {
	Name *string
	Url  *string
}

type MerchantCreate struct {
	Name string
	Url  *string
}

type Store struct {
	ID       int32
	Name     *string
	Location *string
}

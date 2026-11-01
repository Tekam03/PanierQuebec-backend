package model

type Merchant struct {
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
	Name     string
	Location string
	Merchant Merchant
}

type StorePatch struct {
	Name     *string
	Location *string
	MerchantID *int32
}

type StoreCreate struct {
	Name     string
	Location *string
	MerchantID int32
}

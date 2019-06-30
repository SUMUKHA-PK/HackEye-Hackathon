package util

// AddToCartReq is the data to be added to a users cart
type AddToCartReq struct {
	UserID     string
	ItemList   []string
	ItemIDList []string
}

// StoreCheckoutReq is the /store/checkout req
type StoreCheckoutReq struct {
	UserID   string
	ItemList []string
}

type HomeCheckout struct {
	UserID   string
	ItemList []string
}

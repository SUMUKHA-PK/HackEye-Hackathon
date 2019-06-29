package util

// AddToCartReq is the data to be added to a users cart
type AddToCartReq struct {
	UserID     string
	ItemList   []string
	ItemIDList []string
}

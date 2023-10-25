package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

//
//func (app *Application) AddToCart() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		productQueryID := c.Query("id")
//		if productQueryID == "" {
//			log.Println("product id is empty")
//
//			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
//			return
//		}
//
//		userQueryID := c.Query("userID")
//		if userQueryID == "" {
//			log.Println("user id is empty")
//
//			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
//			return
//		}
//
//		productID, err := primitive.ObjectIDFromHex(productQueryID)
//		if err != nil {
//			log.Println(err)
//			_ = c.AbortWithError(http.StatusInternalServerError)
//			return
//		}
//
//		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
//		defer cancel()
//		err = database.AddProductTocart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
//		if err != nil {
//			c.IndentedJSON(http.StatusInternalServerError, err)
//		}
//
//		c.IndentedJSON(200, "successfully add to the cart")
//	}
//}
//
//func (app *Application) RemoveItem() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
//func GetItemFromCart() gin.HandlerFunc {
//
//}
//
//func BuyFromCart() gin.HandlerFunc {
//
//}
//
//func InstantBuy() gin.HandlerFunc {
//
//}

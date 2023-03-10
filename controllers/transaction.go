package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection *mongo.Collection = db.GetCollection(db.DB, "transaction")
var validate = validator.New()

func GenerateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()
		if err := c.BindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}
		if validationErr := validate.Struct(transaction); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}
		err := transactionCollection.FindOne(ctx, bson.M{"transaction_id" : transactionId})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
			return
		}
		result, err := transactionCollection.InsertOne(ctx, transaction)
		defer cancel()
		if err != nil {
			msg := Sprintf("Unable to add new transaction")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, transaction) 
	}
}

func GetTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()

		if err := c.JSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		objId, _ := primitive.ObjectIDFromHex(transactionId)
		result, err := transactionCollection.FindOne(ctx, bson.M{"id" : objId}).Decode(&transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}

func GetAllTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		results, err := transactionList.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			msg := Sprintf("An error while listing transactions")
			c.JSON(http.StatusBadRequest, gin.H{"error" : msg})
			return
		}
		var transaction []bson.M
		if err := results.All(ctx, &transactions); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, transaction)
	}
}

func GenerateTransactionToken() string {

}

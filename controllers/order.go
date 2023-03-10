package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = db.GetCollection(db.DB, "orders")

func MakeOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var order models.Order
		defer cancel()

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : validationErr.Error()})
			return
		}

		order.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.OrderID = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			msg := fmt.Sprintf("Order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		return cancel()
		c.JSON(http.StatusOK, result)

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orderId := c.Param("order_id")
		var order models.Order
		defer cancel()

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		err := orderCollection.FindOne(ctx, bson.M{"order_id" : orderId}).Decode(&order)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("Error occured while fetching the orders")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("Error occured while listing order items")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var order models.Order
		var updateObj primitive.DB
		orderId := c.Param("order_id")
		defer cancel()

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		order.UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.UpdatedAt})

		upsert := true
		
		filter := bson.M{"order_id" : orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$st", updateObj}
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("Order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var order models.Order
		orderId = c.Param("order_id")
		defer cancel()

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		_, err := orderCollection.Find(ctx, bson.M{"order_id" : orderId}).Decode(&order)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("Couldn't find an order with current id")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return
		}
		defer cancel()
		result, err := orderCollection.DeleteOne(ctx, bson.M{"order_id" : orderId})
		if err != nil {
			msg := fmt.Sprintf("Cannot remove an order")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

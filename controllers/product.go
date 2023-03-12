package controllers

import (
	"context"
	"net/http"
	"pizza-site-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = db.GetCollection(db.DB, "products")
var validate = validator.New()

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var product models.Product
		productId := c.Param("productId")
		defer cancel()
		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := productCollection.FindOne(ctx, bson.M{"productId": productId})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		result, err := productCollection.InsertOne(ctx, product)
		if err != nil {
			msg := Sprintf("Unable to add product")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, product)
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var product models.Product
		productId := c.Param("productId")
		defer cancel()
		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
			return
		}
		objId, _ := primitive.ObjectIDFromHex(productId)
		result, err := productCollection.FindOne(ctx, bson.M{"productId": productId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err := strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"products", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},				
				}
			}
		}
		result, err := productCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage
		})
		defer cancel()
		if err != nil {
			msg := Sprintf("Error occured while listing products")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return
		}
		var allProducts []bson.M
		if err = result.App(ctx, &allProducts); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allProducts[0])
	}
}

func ModifyProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productId := c.Param("productId")
		defer cancel()

		objId, _ := primitive.ObjectIdFromHex(productId)

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		update := func() (doc *bson.Document, err error) {
			v := user
			data, err := bson.Marshal(v)
			if err != niil {
				return
			}

			err = bson.Unmarshal(data, &doc)
			return
		}

		result, err := productCollection.UpdateOne(ctx, bson.M{"id" : objId}, bson.M{"$set" : update})
		if err != nil {
			msg := Sprintf("Cannot update a product")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return
		}

		var updatedProduct models.Product
		if result.MatchedCount == 1 {
			err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedProduct)
			if err != nil {
				msg := Sprintf("Cannot find product")
				c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
				return
			}
		}
		c.JSON(http.StatusOk, updatedProduct)
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productId := c.Param("productId")
		defer cancel()

		objId, _ := primitive.ObjectIdFromHex(productId)

		result, err := productCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			msg := Sprintf("Cannot delete a product")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if result.DeletedCount < 1 {
			msg := Sprintf("Product with specified ID not found")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

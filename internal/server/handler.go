package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RaphSku/micro/internal/database"
	"github.com/RaphSku/micro/internal/utils"
	"github.com/RaphSku/micro/src/product"
	"github.com/RaphSku/micro/src/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserHandler(db string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		mongodb_uri := utils.GetEnvVariable("MONGODB_URI")
		client, ctx := database.ConnectToMongoDB(mongodb_uri)
		defer client.Disconnect(ctx)

		db := client.Database(db)
		collection := db.Collection("user")

		if r.Method == "GET" {
			cursor, err := collection.Find(ctx, bson.D{})
			if err != nil {
				log.Fatal(err)
			}

			var decoded_documents []user.User
			for cursor.Next(context.TODO()) {
				var item user.User
				err := cursor.Decode(&item)
				if err != nil {
					log.Fatal(err)
				}
				decoded_documents = append(decoded_documents, item)
			}
			fmt.Println("The decoded documents are the following:\n", decoded_documents)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(decoded_documents)
		}

		if r.Method == "POST" {
			if r.Body == nil {
				http.Error(w, "A POST request needs a request body", 400)
				fmt.Println("Empty body")
				return
			}

			var item user.User
			err := json.NewDecoder(r.Body).Decode(&item)
			if err != nil {
				http.Error(w, err.Error(), 400)
				fmt.Println("Error:", err)
				return
			}

			insert_result, err := collection.InsertOne(ctx, item)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(insert_result)
			fmt.Println("The inserted document is the following:\n", item)

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			response := make(map[string]string)
			response["message"] = "Status Created"
			jsonResp, err := json.Marshal(response)
			if err != nil {
				log.Fatalf("Error during JSON Marhal: %s", err)
			}
			w.Write(jsonResp)
		}

		if r.Method == "DELETE" {
			if r.Body == nil {
				http.Error(w, "A DELETE request needs a request body", 400)
				fmt.Println("Empty body")
				return
			}

			var users user.UserList
			err := json.NewDecoder(r.Body).Decode(&users)
			if err != nil {
				http.Error(w, err.Error(), 400)
				fmt.Println("Error:", err)
				return
			}

			var ids []primitive.ObjectID
			for _, user := range users.Users {
				id := user.ID
				ids = append(ids, id)
			}

			filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}
			fmt.Println(filter)
			result, err_delete := collection.DeleteMany(context.TODO(), filter)
			if err_delete != nil {
				fmt.Println("Error", err_delete)
				return
			}

			fmt.Println("Number of documents deleted:", result)

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			response := make(map[string]string)
			response["message"] = "Status Created"
			jsonResp, err := json.Marshal(response)
			if err != nil {
				log.Fatalf("Error during JSON Marhal: %s", err)
			}
			w.Write(jsonResp)
		}

		if r.Method == "PUT" {
			if r.Body == nil {
				http.Error(w, "An UPDATE request needs a request body", 400)
				fmt.Println("Empty body")
				return
			}

			var users user.UserList
			err_decode := json.NewDecoder(r.Body).Decode(&users)
			if err_decode != nil {
				http.Error(w, err_decode.Error(), 400)
				fmt.Println("Error:", err_decode)
				return
			}

			var ids []primitive.ObjectID
			for _, user := range users.Users {
				id := user.ID
				ids = append(ids, id)
			}

			filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}
			update := bson.A{}
			for _, update_user := range users.Users {
				item := bson.M{
					"name":              update_user.Name,
					"username":          update_user.Username,
					"email":             update_user.Email,
					"address":           update_user.Address,
					"registration_date": update_user.RegistrationDate,
				}
				update = append(update, bson.D{{Key: "$set", Value: item}})
			}

			fmt.Println(filter, update)
			result, err := collection.UpdateMany(context.TODO(), filter, update)
			if err != nil {
				fmt.Println("Error", err)
				return
			}

			if result.ModifiedCount != 0 {
				fmt.Println("Number of updated documents:", result.ModifiedCount)

				w.WriteHeader(http.StatusCreated)
				w.Header().Set("Content-Type", "application/json")
				response := make(map[string]string)
				response["message"] = "Status Created"
				jsonResp, err := json.Marshal(response)
				if err != nil {
					log.Fatalf("Error during JSON Marhal: %s", err)
				}
				w.Write(jsonResp)
				return
			}
			fmt.Println("No matched document found!")

			http.Error(w, "404 not found", http.StatusNotFound)
		}
	}
}

func GetProductHandler(db string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/product" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		mongodb_uri := utils.GetEnvVariable("MONGODB_URI")
		client, ctx := database.ConnectToMongoDB(mongodb_uri)
		defer client.Disconnect(ctx)

		db := client.Database(db)
		collection := db.Collection("product")

		if r.Method == "GET" {
			cursor, err := collection.Find(ctx, bson.D{})
			if err != nil {
				log.Fatal(err)
			}

			var decoded_documents []product.Product
			for cursor.Next(context.TODO()) {
				var item product.Product
				err := cursor.Decode(&item)
				if err != nil {
					log.Fatal(err)
				}
				decoded_documents = append(decoded_documents, item)
			}
			fmt.Println("The decoded documents are the following:\n", decoded_documents)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(decoded_documents)
		}

		if r.Method == "POST" {
			if r.Body == nil {
				http.Error(w, "A POST request needs a request body", 400)
				fmt.Println("Empty body")
				return
			}

			var item product.Product
			err := json.NewDecoder(r.Body).Decode(&item)
			if err != nil {
				http.Error(w, err.Error(), 400)
				fmt.Println("Error:", err)
				return
			}

			insert_result, err := collection.InsertOne(ctx, item)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(insert_result)
			fmt.Println("The inserted document is the following:\n", item)

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			response := make(map[string]string)
			response["message"] = "Status Created"
			jsonResp, err := json.Marshal(response)
			if err != nil {
				log.Fatalf("Error during JSON Marhal: %s", err)
			}
			w.Write(jsonResp)
		}

		if r.Method == "DELETE" {
			if r.Body == nil {
				http.Error(w, "A DELETE request needs a request body", 400)
				fmt.Println("Empty body")
				return
			}

			var products product.ProductList
			err := json.NewDecoder(r.Body).Decode(&products)
			if err != nil {
				http.Error(w, err.Error(), 400)
				fmt.Println("Error:", err)
				return
			}

			var ids []primitive.ObjectID
			for _, user := range products.Products {
				id := user.ID
				ids = append(ids, id)
			}

			filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}
			fmt.Println(filter)
			result, err_delete := collection.DeleteMany(context.TODO(), filter)
			if err_delete != nil {
				fmt.Println("Error", err_delete)
				return
			}

			fmt.Println("Number of documents deleted:", result)

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			response := make(map[string]string)
			response["message"] = "Status Created"
			jsonResp, err := json.Marshal(response)
			if err != nil {
				log.Fatalf("Error during JSON Marhal: %s", err)
			}
			w.Write(jsonResp)
		}

		if r.Method == "PUT" {
			if r.Body == nil {
				http.Error(w, "An UPDATE request needs a request body", 400)
				fmt.Println("Empty body")
				return
			}

			var products product.ProductList
			err_decode := json.NewDecoder(r.Body).Decode(&products)
			if err_decode != nil {
				http.Error(w, err_decode.Error(), 400)
				fmt.Println("Error:", err_decode)
				return
			}

			var ids []primitive.ObjectID
			for _, user := range products.Products {
				id := user.ID
				ids = append(ids, id)
			}

			filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}
			update := bson.A{}
			for _, update_product := range products.Products {
				item := bson.M{
					"name":     update_product.Name,
					"price":    update_product.Price,
					"category": update_product.Category,
					"storage":  update_product.Storage,
				}
				update = append(update, bson.D{{Key: "$set", Value: item}})
			}

			fmt.Println(filter, update)
			result, err := collection.UpdateMany(context.TODO(), filter, update)
			if err != nil {
				fmt.Println("Error", err)
				return
			}

			if result.ModifiedCount != 0 {
				fmt.Println("Number of updated documents:", result.ModifiedCount)

				w.WriteHeader(http.StatusCreated)
				w.Header().Set("Content-Type", "application/json")
				response := make(map[string]string)
				response["message"] = "Status Created"
				jsonResp, err := json.Marshal(response)
				if err != nil {
					log.Fatalf("Error during JSON Marhal: %s", err)
				}
				w.Write(jsonResp)
				return
			}
			fmt.Println("No matched document found!")

			http.Error(w, "404 not found", http.StatusNotFound)
		}
	}
}

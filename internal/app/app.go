package app

import (
	"bootstrapservice/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Run(cfg *config.Config) {
	// Run application
	dbUrl := fmt.Sprintf("mongodb://%s:%s/%s", cfg.MONGO.HOST, cfg.MONGO.PORT, cfg.MONGO.DB)
	dbOptions := options.Client().ApplyURI(dbUrl)
	db, err := mongo.Connect(context.TODO(), dbOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		os.Exit(1)
	}
	fmt.Println("Connected to MongoDB")

	defer dbTidy(db)

}

func dbTidy(db *mongo.Client) {
	err := db.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("Error disconnecting from MongoDB:", err)
		os.Exit(1)
	}
}

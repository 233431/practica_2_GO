package infraestructure

import (
	"context"
	"practica_2/src/user/domain"
	"practica_2/src/core"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBRepository gestiona la colección de productos
type MongoDBRepository struct {
	collection *mongo.Collection
}

// NewMongoDBRepository crea un nuevo repositorio para la colección "productos"
func NewMongoDBRepository() *MongoDBRepository {
	collection := core.GetCollection("productos") // Asegúrate de que esta sea la colección correcta
	return &MongoDBRepository{collection: collection}
}

// Save inserta un nuevo producto en MongoDB
func (r *MongoDBRepository) Save(p *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, p)
	return err
}

// Delete elimina un producto por su nombre
func (r *MongoDBRepository) Delete(nombre string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"nombre": nombre})
	return err
}

// Update actualiza un producto en MongoDB
func (r *MongoDBRepository) Update(id string, p *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"nombre": p.Nombre, "precio": p.Precio}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

// GetAll obtiene todos los productos de la colección
func (r *MongoDBRepository) GetAll() ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []domain.Product
	for cursor.Next(ctx) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

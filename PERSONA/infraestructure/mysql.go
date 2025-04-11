package infraestructure

import (
	"recu/PERSONA/domain"
	"recu/PERSONA/core"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoPersonRepository struct {
	collection *mongo.Collection
}

func NewMongoPersonRepository() *MongoPersonRepository {
	client := core.GetMongoClient()
	collection := client.Database("recu").Collection("personas")
	return &MongoPersonRepository{collection: collection}
}

func (r *MongoPersonRepository) Save(p *domain.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		log.Printf("Error al insertar persona: %v", err)
	}
	return err
}

func (r *MongoPersonRepository) GetAll() ([]domain.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetLimit(5).SetSort(bson.D{{"_id", -1}})
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Printf("Error al obtener personas: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var people []domain.Person
	for cursor.Next(ctx) {
		var person domain.Person
		if err := cursor.Decode(&person); err != nil {
			log.Printf("Error al decodificar persona: %v", err)
			continue
		}
		people = append(people, person)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error al recorrer personas: %v", err)
		return nil, err
	}

	return people, nil
}



func (r *MongoPersonRepository) GetCountM() (int, error) {
	return r.countBySexo("M")
}

func (r *MongoPersonRepository) GetCountH() (int, error) {
	return r.countBySexo("H")
}

func (r *MongoPersonRepository) countBySexo(sexo string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{"sexo": sexo})
	if err != nil {
		log.Printf("Error al contar personas con sexo %s: %v", sexo, err)
		return 0, err
	}
	return int(count), nil
}

// Helper para puntero a int64
func int64Ptr(i int64) *int64 {
	return &i
}

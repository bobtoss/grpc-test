package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"libraryService/internal/model/book"
)

type BookRepository struct {
	db *mongo.Collection
}

func NewBookRepository(db *mongo.Collection) BookRepository {
	return BookRepository{
		db: db,
	}
}

func (s *BookRepository) List(ctx context.Context) (dest []*book.Entity, err error) {
	cursor, err := s.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &dest); err != nil {
		return nil, err
	}
	return
}

func (s *BookRepository) Add(ctx context.Context, req *book.Entity) error {
	_, err := s.db.InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookRepository) Get(ctx context.Context, id primitive.ObjectID) (*book.Entity, error) {
	res := book.Entity{}
	err := s.db.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *BookRepository) Update(ctx context.Context, req *book.Entity) error {
	options.Update().SetUpsert(true)
	pByte, err := bson.Marshal(req)
	if err != nil {
		return err
	}
	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", req.ObjectID}}
	_, err = s.db.UpdateOne(ctx, filter, bson.D{{"$set", update}})
	if err != nil {
		return err
	}
	return nil
}

func (s *BookRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.db.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

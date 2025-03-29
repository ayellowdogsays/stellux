package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/codepzj/Stellux/server/global"
	"github.com/codepzj/Stellux/server/internal/posts/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IPostsDao interface {
	FindById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error)
	FindListByCondition(ctx context.Context, skip int64, limit int64, keyword string, field string, order int) ([]*domain.Posts, error)
	FindAll(ctx context.Context) ([]*domain.Posts, error)
	GetAllCount(ctx context.Context) (int64, error)
	GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error)
	AdminFindListByCondition(ctx context.Context, skip int64, limit int64, keyword string, field string, order int) ([]*domain.Posts, error)
	AdminGetAllCountByKeyword(ctx context.Context, keyword string) (int64, error)
	AdminFindOneAndUpdateStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error
	AdminCreate(ctx context.Context, posts *domain.Posts) error
	AdminUpdate(ctx context.Context, posts *domain.Posts) error
	AdminResumePostById(ctx context.Context, id bson.ObjectID) error
	AdminDeleteSoftById(ctx context.Context, id bson.ObjectID) error
	AdminDeletePostById(ctx context.Context, id bson.ObjectID) error
}

type PostsDao struct {
	postColl *mongo.Collection
}

var _ IPostsDao = (*PostsDao)(nil)

func NewPostsDao() *PostsDao {
	return &PostsDao{postColl: global.DB.Collection("posts")}
}

// FindById 查询特定文章
func (p *PostsDao) FindById(ctx context.Context, id bson.ObjectID) (*domain.Posts, error) {
	var posts domain.Posts
	err := p.postColl.FindOne(ctx, bson.D{{Key: "_id", Value: id}, {Key: "is_publish", Value: true}}).Decode(&posts)
	if err != nil {
		return nil, errors.Wrap(err, "查询文章失败")
	}
	return &posts, nil
}

// FindListByCondition 分页，关键词查询文章列表，筛除删除，未发布的文章
func (p *PostsDao) FindListByCondition(ctx context.Context, skip int64, limit int64, keyword string, field string, order int) ([]*domain.Posts, error) {
	findOptions := options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.M{field: order})

	// 筛除删除，未发布的结果
	filter1 := bson.D{{Key: "is_publish", Value: true}, {Key: "deleted_at", Value: nil}}
	// 筛选搜索内容在title、description中出现
	filter2 := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: keyword}}}},
		bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: keyword}}}},
	}}}

	filter := bson.D{{Key: "$and", Value: bson.A{filter1, filter2}}}

	cursor, err := p.postColl.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	var posts []domain.Posts
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}
	return domain.ToPtr(posts), nil
}

// FindAll 获取所有文章
func (p *PostsDao) FindAll(ctx context.Context) ([]*domain.Posts, error) {
	cursor, err := p.postColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "查询文章失败")
	}
	var posts []*domain.Posts
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, errors.Wrap(err, "查询文章失败")
	}
	return posts, nil
}

// GetAllCount 获取文章总数
func (p *PostsDao) GetAllCount(ctx context.Context) (int64, error) {
	return p.postColl.CountDocuments(ctx, bson.M{})
}

// GetAllCountByKeyword 用户通过关键词获取文章总数
func (p *PostsDao) GetAllCountByKeyword(ctx context.Context, keyword string) (int64, error) {
	// 筛除删除，未发布的结果
	filter1 := bson.D{{Key: "is_publish", Value: true}, {Key: "deleted_at", Value: nil}}
	// 筛选搜索内容在title、description中出现
	filter2 := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: keyword}}}},
		bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: keyword}}}},
	}}}

	filter := bson.D{{Key: "$and", Value: bson.A{filter1, filter2}}}
	return p.postColl.CountDocuments(ctx, filter)
}

// AdminFindListByCondition 管理员分页，关键词查询文章列表，筛除删除的文章,不区分是否发布
func (p *PostsDao) AdminFindListByCondition(ctx context.Context, skip int64, limit int64, keyword string, field string, order int) ([]*domain.Posts, error) {
	findOptions := options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.M{field: order})

	filter1 := bson.D{{Key: "deleted_at", Value: nil}}
	filter2 := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: keyword}}}},
		bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: keyword}}}},
	}}}

	filter := bson.D{{Key: "$and", Value: bson.A{filter1, filter2}}}

	cursor, err := p.postColl.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	var posts []domain.Posts
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}
	return domain.ToPtr(posts), nil
}

// AdminGetAllCountByKeyword 管理员通过关键词获取文章总数
func (p *PostsDao) AdminGetAllCountByKeyword(ctx context.Context, keyword string) (int64, error) {
	// 筛除删除的文章
	filter1 := bson.D{{Key: "deleted_at", Value: nil}}
	// 筛选搜索内容在title、description中出现
	filter2 := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: keyword}}}},
		bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: keyword}}}},
	}}}
	filter := bson.D{{Key: "$and", Value: bson.A{filter1, filter2}}}
	return p.postColl.CountDocuments(ctx, filter)
}

// 管理员查看被软删除文章
// func (p *PostsDao) AdminFindListByDeleted(ctx context.Context, skip int64, limit int64, field string, order int) ([]*domain.Posts, error) {
// 	findOptions := options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.M{field: order})
// 	cursor, err := p.postColl.Find(ctx, bson.D{{Key: "deleted_at", Value: bson.M{"$ne": nil}}}, findOptions)
// 	if err != nil {
// 		return nil, err
// 	}
// }
// AdminCreate 管理员创建文章
func (p *PostsDao) AdminCreate(ctx context.Context, posts *domain.Posts) error {
	posts.ID = bson.NewObjectID()
	posts.CreatedAt = time.Now()
	posts.UpdatedAt = time.Now()
	_, err := p.postColl.InsertOne(ctx, posts)
	if err != nil {
		return errors.Wrap(err, "添加文章失败")
	}
	return nil
}

// AdminUpdate 管理员更新文章
func (p *PostsDao) AdminUpdate(ctx context.Context, posts *domain.Posts) error {
	posts.UpdatedAt = time.Now()
	fmt.Println("post", posts)
	result, err := p.postColl.UpdateOne(ctx, bson.M{"_id": posts.ID}, bson.M{"$set": posts})
	if result.ModifiedCount == 0 {
		return errors.New("更新条数为0，更新失败")
	}
	return err
}

// AdminFindOneAndUpdateStatus 管理员上下架文章
func (p *PostsDao) AdminFindOneAndUpdateStatus(ctx context.Context, id bson.ObjectID, isPublish *bool) error {
	result := p.postColl.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"is_publish": isPublish}})
	return result.Err()
}

// AdminDeleteSoftById 管理员软删除文章
func (p *PostsDao) AdminDeleteSoftById(ctx context.Context, id bson.ObjectID) error {
	result := p.postColl.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Local()}})
	return result.Err()
}

// AdminResumePostById 管理员恢复文章状态
func (p *PostsDao) AdminResumePostById(ctx context.Context, id bson.ObjectID) error {
	// 删除deleted_at字段
	result := p.postColl.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$unset": bson.M{"deleted_at": nil}})
	return result.Err()
}

// AdminDeletePostById 管理员硬删除文章
func (p *PostsDao) AdminDeletePostById(ctx context.Context, id bson.ObjectID) error {
	result, err := p.postColl.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrap(err, "删除失败")
	}
	if result.DeletedCount == 0 {
		return errors.Wrap(err, "删除条数为0，删除失败")
	}
	return nil
}

package adapter

import (
	"categorys/domain/category"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	facet = bson.D{{"$facet", bson.M{
		"snapshot": bson.A{
			bson.M{"$match": bson.M{"eventtype": category.Snapshotted}},
			bson.M{"$group": bson.M{
				"_id":       nil,
				"timestamp": bson.M{"$max": "$timestamp"},
			}},
		},
		"data": bson.A{bson.M{"$replaceRoot": bson.M{"newRoot": "$$ROOT"}}},
	}}}
	unwind1     = bson.D{{"$unwind", "$snapshot"}}
	unwind2     = bson.D{{"$unwind", "$data"}}
	match       = bson.D{{"$match", bson.M{"$expr": bson.M{"$gte": bson.A{"$data.timestamp", "$snapshot.timestamp"}}}}}
	replaceRoot = bson.D{{"$replaceRoot", bson.M{"newRoot": "$data"}}}

	limit1             = bson.D{{"$limit", 1}}
	sortTimestamp      = bson.D{{"$sort", bson.M{"timestamp": -1}}}
	matchNoSnapshotted = bson.D{{"$match", bson.M{"type": bson.M{"$ne": category.Snapshotted}}}}
)

func getCategoryQ(aggid uint32) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{"$match", bson.M{"aggregateid": aggid}}},
		facet,
		unwind1,
		unwind2,
		match,
		replaceRoot,
	}
}

func checkCategoryIsExistQ(aggid uint32) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{"$match", bson.M{"aggregateid": aggid, "eventtype": bson.M{"$in": bson.A{category.Snapshotted, category.Deleted}}}}},
		sortTimestamp,
		limit1,
		matchNoSnapshotted,
	}
}

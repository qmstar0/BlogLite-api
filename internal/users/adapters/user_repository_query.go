package adapters

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"users/domain/user"
)

var (
	facet = bson.D{{"$facet", bson.M{
		"baseline": bson.A{
			bson.M{"$match": bson.M{"eventtype": user.Login}},
			bson.M{"$group": bson.M{
				"_id":       nil,
				"timestamp": bson.M{"$max": "$timestamp"},
			}},
		},
		"event": bson.A{bson.M{"$replaceRoot": bson.M{"newRoot": "$$ROOT"}}},
	}}}
	unwind1     = bson.D{{"$unwind", "$baseline"}}
	unwind2     = bson.D{{"$unwind", "$event"}}
	match       = bson.D{{"$match", bson.M{"$expr": bson.M{"$gte": bson.A{"$event.timestamp", "$baseline.timestamp"}}}}}
	replaceRoot = bson.D{{"$replaceRoot", bson.M{"newRoot": "$event"}}}

	limit1             = bson.D{{"$limit", 1}}
	sortTimestamp      = bson.D{{"$sort", bson.M{"timestamp": -1}}}
	matchNoSnapshotted = bson.D{{"$match", bson.M{"type": bson.M{"$ne": user.Login}}}}
)

func getUserQ(aggid uint32) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{"$match", bson.M{"aggregateid": aggid}}},
		facet,
		unwind1,
		unwind2,
		match,
		replaceRoot,
	}
}

func checkUserIsExistQ(aggid uint32) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{"$match", bson.M{"aggregateid": aggid, "eventtype": bson.M{"$in": bson.A{user.Login, user.Logout}}}}},
		sortTimestamp,
		limit1,
		matchNoSnapshotted,
	}
}

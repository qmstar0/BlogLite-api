package eventstore

import (
	"common/domainevent"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var facet, unwind1, unwind2, sort, match, replaceRoot bson.D

func init() {
	facet = bson.D{{"$facet", bson.M{
		"snapshot": bson.A{
			bson.M{"$match": bson.M{"eventtype": domainevent.Snapshotted}},
			bson.M{"$group": bson.M{
				"_id":       nil,
				"timestamp": bson.M{"$max": "$timestamp"},
			}},
		},
		"data": bson.A{bson.M{"$replaceRoot": bson.M{"newRoot": "$$ROOT"}}},
	}}}
	unwind1 = bson.D{{"$unwind", "$snapshot"}}
	unwind2 = bson.D{{"$unwind", "$data"}}
	match = bson.D{{"$match", bson.M{"$expr": bson.M{"$gte": bson.A{"$data.timestamp", "$snapshot.timestamp"}}}}}
	replaceRoot = bson.D{{"$replaceRoot", bson.M{"newRoot": "$data"}}}
	sort = bson.D{{"$sort", bson.M{"timestamp": 1}}}
}

func DefaultEntityQuery(aggid uint32) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{"$match", bson.M{"aggregateid": aggid}}},
		facet,
		unwind1,
		unwind2,
		match,
		replaceRoot,
		sort,
	}
}

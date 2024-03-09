package adapters

import (
	"common/domainevent"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var facet, unwind1, unwind2, project, unwind3, replaceRoot bson.D

func initAggregateQuery() {
	facet = bson.D{{"$facet", bson.M{
		"snapshotted": bson.A{
			bson.M{"$match": bson.M{"eventtype": domainevent.Snapshotted}},
			bson.M{"$sort": bson.M{"timestamp": -1}},
			bson.M{"$limit": 1},
		},
		"data": bson.A{bson.M{"$replaceRoot": bson.M{"newRoot": "$$ROOT"}}},
	}}}

	unwind1 = bson.D{{"$unwind", "$snapshotted"}}

	project = bson.D{{"$project",
		bson.M{"result": bson.M{"$filter": bson.M{"input": "$data", "cond": bson.M{"$gte": bson.A{"$$this.timestamp", "$last1.lastTimestamp"}}}}}}}

	unwind3 = bson.D{{"$unwind", "$result"}}
	replaceRoot = bson.D{{"$replaceRoot", bson.M{"newRoot": "$result"}}}
}

func queryToGetLastSnapshotAndLastEvent(aggid uint32) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{"$match", bson.M{"aggregateid": aggid}}},
		facet,
		unwind1,
		unwind2,
		project,
		unwind3,
		replaceRoot,
	}
}

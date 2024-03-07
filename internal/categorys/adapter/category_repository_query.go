package adapter

import (
	"common/domainevent"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var facet, unwind1, unwind2, project, unwind3, replaceRoot bson.D

var (
	EntityNotFind = errors.New("the resource does not exist")
)

func initAggregateQuery() {
	facet = bson.D{{"$facet", bson.M{
		"last1": bson.A{bson.M{"$match": bson.M{"type": domainevent.Snapshotted}}, bson.M{"$group": bson.M{"_id": nil, "lastTimestamp": bson.M{"$max": "$timestamp"}}}},
		"last2": bson.A{bson.M{"$match": bson.M{"$or": bson.A{bson.M{"type": domainevent.Deleted}, bson.M{"type": domainevent.Created}}}}, bson.M{"$group": bson.M{"_id": nil, "lastTimestamp": bson.M{"$max": "$timestamp"}}}},
		"data":  bson.A{bson.M{"$replaceRoot": bson.M{"newRoot": "$$ROOT"}}},
	}}}

	unwind1 = bson.D{{"$unwind", "$last1"}}
	unwind2 = bson.D{{"$unwind", "$last2"}}

	project = bson.D{{"$project", bson.M{"result": bson.M{"$cond": bson.A{
		bson.M{"$gt": bson.A{"$last2.lastTimestamp", "$last1.lastTimestamp"}},
		//  如果 last8 大，则返回空数据
		bson.M{"$literal": []any{}},
		// 如果 last1 大，则返回 data 中所有时间戳大于 last1 的数据
		bson.M{"$filter": bson.M{"input": "$data", "cond": bson.M{"$gte": bson.A{"$$this.timestamp", "$last1.lastTimestamp"}}}},
	}}}}}

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

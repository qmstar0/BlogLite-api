package mongodb_test

import (
	"blog/pkg/mongodb"
	"categorys/adapter"
	"common/domainevent"
	"common/idtools"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand/v2"
	"testing"
	"time"
)

var testData *adapter.CategoryDomainEventStoreModel

func init() {
	data := map[string]any{"k": "v"}
	marshal, err := bson.Marshal(data)
	if err != nil {
		panic(err)
	}
	testData = &adapter.CategoryDomainEventStoreModel{
		EventID:     "",
		AggregateID: 1762045341,
		Type:        0,
		Event:       marshal,
		Timestamp:   time.Now(),
	}
}

func TestConnect(t *testing.T) {
	fn := mongodb.Init()
	defer fn(context.Background())
	_ = mongodb.GetDB()

}

func getRandTime() time.Duration {
	return time.Duration(rand.UintN(5)*100) * time.Millisecond
}

func getdb() *mongo.Collection {
	mongodb.Init()
	return mongodb.GetDB().Collection("Domain_EventStore_Cateogry")
}
func TestClean(t *testing.T) {
	ctx := context.Background()
	db := getdb()
	many, err := db.DeleteMany(ctx, bson.M{"aggregateid": 1762045341})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(many)
}

func testAggregateFind(t *testing.T, db *mongo.Collection, datalength int) {
	ctx := context.Background()
	var pipeline = mongo.Pipeline{
		bson.D{{"$match", bson.M{"aggregateid": 1762045341}}},

		bson.D{{"$facet", bson.M{
			"last1": bson.A{bson.M{"$match": bson.M{"type": 1}}, bson.M{"$group": bson.M{"_id": nil, "lastTimestamp": bson.M{"$max": "$timestamp"}}}},
			"last8": bson.A{bson.M{"$match": bson.M{"$or": bson.A{bson.M{"type": 8}, bson.M{"type": 2}}}}, bson.M{"$group": bson.M{"_id": nil, "lastTimestamp": bson.M{"$max": "$timestamp"}}}},
			"data":  bson.A{bson.M{"$replaceRoot": bson.M{"newRoot": "$$ROOT"}}},
		}}},

		bson.D{{"$unwind", "$last1"}},
		bson.D{{"$unwind", "$last8"}},
		bson.D{{"$project", bson.M{
			"result": bson.M{"$cond": bson.A{
				bson.M{"$gt": bson.A{"$last8.lastTimestamp", "$last1.lastTimestamp"}},
				//  如果 last8 大，则返回空数据
				bson.M{"$literal": []any{}},
				// 如果 last1 大，则返回 data 中所有时间戳大于 last1 的数据
				bson.M{"$filter": bson.M{"input": "$data", "cond": bson.M{"$gte": bson.A{"$$this.timestamp", "$last1.lastTimestamp"}}}},
			}},
		}}},
		bson.D{{"$unwind", "$result"}},
		bson.D{{"$replaceRoot", bson.M{"newRoot": "$result"}}},
	}
	// 执行聚合查询
	cursor, err := db.Aggregate(ctx, pipeline)
	if err != nil {
		t.Fatal(err)
	}
	var datas = make([]adapter.CategoryDomainEventStoreModel, 0)
	err = cursor.All(ctx, &datas)
	if err != nil {
		t.Fatal(err)
	}
	for i, data := range datas {
		t.Log(i, data)
	}
	if datalength != len(datas) {
		t.Fatal("数据量与预期结果不一致", len(datas))
	}
}

func TestAggFind(t *testing.T) {
	testAggregateFind(t, getdb(), 0)
}

// 测试创建数据
func TestData0(t *testing.T) {
	TestClean(t)
	ctx := context.Background()
	var (
		db = getdb()
	)

	for _, d := range []uint16{domainevent.Created, domainevent.Snapshotted, domainevent.Updated} {
		testData.Type = d
		testData.EventID = idtools.NewUUID()
		testData.Timestamp = testData.Timestamp.Add(getRandTime())

		one, err := db.InsertOne(ctx, testData)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(one)
	}
	testAggregateFind(t, db, 2)
}

// 测试删除后不再有数据
func TestData1(t *testing.T) {
	TestClean(t)
	ctx := context.Background()
	var (
		db = getdb()
	)

	for _, d := range []uint16{2, 1, 4, 4, 4, 8} {
		testData.Type = d
		testData.EventID = idtools.NewUUID()
		testData.Timestamp = testData.Timestamp.Add(getRandTime())

		one, err := db.InsertOne(ctx, testData)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(one)
	}
	testAggregateFind(t, db, 0)
}

// 测试删除后重新创建相同数据
func TestData2(t *testing.T) {
	TestClean(t)
	ctx := context.Background()
	var (
		db = getdb()
	)

	for _, d := range []uint16{2, 1, 4, 4, 4, 8, 2, 1, 4, 4} {
		testData.Type = d
		testData.EventID = idtools.NewUUID()
		testData.Timestamp = testData.Timestamp.Add(getRandTime())

		one, err := db.InsertOne(ctx, testData)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(one)
	}
	testAggregateFind(t, db, 3)
}

// 测试数据不存在
func TestData3(t *testing.T) {
	TestClean(t)
	ctx := context.Background()
	var (
		db = getdb()
	)

	for _, d := range []uint16{} {
		testData.Type = d
		testData.EventID = idtools.NewUUID()
		testData.Timestamp = testData.Timestamp.Add(getRandTime())

		one, err := db.InsertOne(ctx, testData)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(one)
	}
	testAggregateFind(t, db, 0)
}

// 测试存在时间长的数据
func TestData4(t *testing.T) {
	TestClean(t)
	ctx := context.Background()
	var (
		db = getdb()
	)

	for _, d := range []uint16{2, 1, 4, 4, 4, 1, 4, 1, 4, 4, 1, 4} {
		testData.Type = d
		testData.EventID = idtools.NewUUID()
		testData.Timestamp = testData.Timestamp.Add(getRandTime())

		one, err := db.InsertOne(ctx, testData)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(one)
	}
	testAggregateFind(t, db, 2)
}

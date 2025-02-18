package cache

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/core/utils"
	"configcenter/src/common/mapstr"
	"configcenter/src/source_controller/coreservice/core/operation"
	"configcenter/src/storage/driver/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/6
 * @Desc: syncunique.go
**/

var CacheObjMap *utils.SafeMap
var cacheMapTicker *time.Ticker
var CacheMapChan chan struct{} = make(chan struct{}, 0)

type syncCacheObj struct {
	//objList        []string
	cacheObjUnique mapstr.MapStr
}

func init() {
	CacheObjMap = utils.NewSafeMap()
}

func SyncObjCache() {
	cacheMapTicker = time.NewTicker(time.Hour * 1)
	// 初始化 objcachemap
	SyncObjCacheMap()

	// crontab objcachemap
	go func() {
		for {
			select {
			case <-cacheMapTicker.C:
				SyncObjCacheMap()
			}
		}
	}()

	// wait chan<-
	go func() {
		for {
			select {
			case <-CacheMapChan:
				SyncObjCacheMap()
			}
		}
	}()

}

func SyncObjCacheMap() {
	ts := new(syncCacheObj)
	//ts.sync()
	ts.matchUniqueDes()
	ts.writeCacheMap()
	//fmt.Printf("%+v\n", CacheObjMap)
	blog.Infof("sync cacheObj plan current:%v", CacheObjMap.M)
}

//func (s *syncCacheObj) filterCacheObj(ctx context.Context) (err error) {
//
//	var result []map[string]interface{}
//	//var result1 []interface{}
//	//err = mongodb.Client().Table(common.BKTableNameObjDes).Find(operation.M{"iscache": true}).
//	//Fields("bk_obj_id").All(ctx, &result)
//	err = mongodb.Client().Table(common.BKTableNameObjDes).Find(operation.M{}).
//		Fields("bk_obj_id").All(ctx, &result)
//	//if err != nil {
//	//	return err, result
//	//}
//	var rr []string = make([]string, 0, len(result))
//	for _, v := range result {
//		rr = append(rr, v["bk_obj_id"].(string))
//	}
//	//copy(rr, result)
//	s.objList = deepcopy.Copy(rr).([]string)
//	return err
//}

//func (s *syncCacheObj) sync() {
//	var ctx context.Context = context.Background()
//	err := s.filterCacheObj(ctx)
//	if err != nil {
//		blog.Errorf("sync host cacheObj err:%v", err)
//		return
//	}
//	if len(s.objList) == 0 {
//		blog.Errorf("sync host cacheObj len:%v", 0)
//		return
//	}
//	s.matchUniqueDes()
//
//}

func (s *syncCacheObj) matchUniqueDes() {
	//s.cacheObjUnique = make([]map[string]interface{})

	//var pipeline []operation.M
	var result []mapstr.MapStr
	var pipeline = []operation.M{
		//{"$match": operation.M{"must_check": true, "bk_supplier_account": "0"}}, // 属性为空时检验
		{"$match": operation.M{"must_check": false, "bk_supplier_account": "0"}}, // 属性为空时检验
		{"$unwind": "$keys"},
		{"$project": operation.M{"key_id": "$keys.key_id", "bk_obj_id": "$bk_obj_id"}},
		{"$lookup": operation.M{"from": "cc_ObjAttDes", "localField": "key_id", "foreignField": "id", "as": "AttDes"}},
		{"$unwind": "$AttDes"},
		{"$match": operation.M{"AttDes.bk_property_id": operation.M{"$ne": "bk_inst_name"}}}, // 不使用 默认实例
		{"$project": operation.M{"_id": 0, "key_id": "$key_id", "bk_obj_id": "$bk_obj_id", "bk_property_id": "$AttDes.bk_property_id"}},
		{"$group": operation.M{"_id": "$bk_obj_id", "items": operation.M{"$addToSet": "$bk_property_id"}}}, //排序后 取第一个 唯一字段
	}
	if err := mongodb.Client().Table(common.BKTableNameObjUnique).AggregateAll(context.Background(), pipeline, &result); err != nil {
		//blog.Errorf("sync obj :%v err:%v", obj, err)
	}
	//fmt.Println(result)
	if len(result) == 0 {
		blog.Errorf("sync obj unique len is %d", 0)
		return
	}

	s.cacheObjUnique = make(mapstr.MapStr, 0)
	for _, value := range result {
		s.cacheObjUnique[value["_id"].(string)] = value["items"]

	}

	//for _,v:=range result{
	//	if v["bk_obj_id"]
	//}
	//for _, obj := range s.objList {
	//	pipeline = []operation.M{
	//		//{"$match": operation.M{"must_check": true, "bk_supplier_account": "0"}}, // 属性为空时检验
	//		{"$match": operation.M{"must_check": false, "bk_supplier_account": "0"}}, // 属性为空时检验
	//		{"$unwind": "$keys"},
	//		{"$project": operation.M{"key_id": "$keys.key_id", "bk_obj_id": "$bk_obj_id"}},
	//		{"$lookup": operation.M{"from": "cc_ObjAttDes", "localField": "key_id", "foreignField": "id", "as": "AttDes"}},
	//		{"$unwind": "$AttDes"},
	//		{"$match": operation.M{"AttDes.bk_property_id": operation.M{"$ne": "bk_inst_name"}}}, // 不使用 默认实例
	//		{"$project": operation.M{"_id": 0, "key_id": "$key_id", "bk_obj_id": "$bk_obj_id", "bk_property_id": "$AttDes.bk_property_id"}},
	//		{"$sort": operation.M{"key_id": 1}},
	//		{"$group": operation.M{"_id": "$bk_obj_id", "first_unique": operation.M{"$first": "$bk_property_id"}}}, //排序后 取第一个 唯一字段
	//		{"$match": operation.M{"_id": obj}},
	//	}
	//	var result map[string]interface{}
	//	if err := mongodb.Client().Table(common.BKTableNameObjUnique).AggregateOne(context.Background(), pipeline, &result); err != nil {
	//		//blog.Errorf("sync obj :%v err:%v", obj, err)
	//		continue
	//	}
	//	if len(result) == 0 {
	//		blog.Errorf("sync obj :%v len is %d", obj, 0)
	//		continue
	//	}
	//	//fmt.Println(result)
	//	//cacheObjUnique[]
	//	s.cacheObjUnique[obj] = result["first_unique"]
	//}

}

func (s *syncCacheObj) writeCacheMap() {
	CacheObjMap.Clear()
	for key, value := range s.cacheObjUnique {
		valueArrs, ok := value.(primitive.A)
		//fmt.Println(key, value, valueArrs)
		if !ok {
			continue
		}
		valueMap := make(mapstr.MapStr, len(valueArrs))
		for _, subvalue := range valueArrs {
			valueMap[subvalue.(string)] = struct{}{}
		}
		CacheObjMap.Put(key, valueMap)
	}
	s.cacheObjUnique.Reset()
}

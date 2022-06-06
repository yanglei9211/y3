package db_model

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gopkg.in/mgo.v2/bson"
)

type TaggedItem struct {
	Id      bson.ObjectId `bson:"_id"`
	Simhash string        `bson:"simhash"`
	Data    struct {
		Type       int    `bson:"type"`
		Subtype    int    `bson:"subtype"`
		Difficulty int    `bson:"difficulty"`
		Stem       string `bson:"stem"`
		Qs         []struct {
			TagIds        []bson.ObjectId   `bson:"tag_ids"`
			SubTagIds     []bson.ObjectId   `bson:"tag_sub_ids"`
			OptsTagIds    [][]bson.ObjectId `bson:"opts_tag_ids"`
			OptsSubTagIds [][]bson.ObjectId `bson:"opts_sub_tag_ids"`
			Desc          string            `bson:"desc"`
			Ans           interface{}       `bson:"ans"`
			Qs            []struct {
				Desc          string            `bson:"desc"`
				TagIds        []bson.ObjectId   `bson:"tag_ids"`
				SubTagIds     []bson.ObjectId   `bson:"tag_sub_ids"`
				OptsTagIds    [][]bson.ObjectId `bson:"opts_tag_ids"`
				OptsSubTagIds [][]bson.ObjectId `bson:"opts_sub_tag_ids"`
			}
		}
		TagIds    []bson.ObjectId `bson:"tag_ids"`
		SubTagIds []bson.ObjectId `bson:"tag_sub_ids"`
	} `bson:"data"`
	UserAdd   bool          `bson:"user_add"`
	RefItemId bson.ObjectId `bson:"ref_item_id"`
}

var TaggingFileds = bson.M{
	"_id":             1,
	"user_add":        1,
	"ref_item_id":     1,
	"data.difficulty": 1, "data.tag_ids": 1,
	"data.stem":        1,
	"data.qs.desc":     1,
	"data.qs.qs.desc":  1,
	"data.qs.ans":      1,
	"data.tag_sub_ids": 1, "data.qs.tag_ids": 1,
	"data.qs.tag_sub_ids": 1, "data.qs.opts_tag_ids": 1,
	"data.qs.opts_sub_tag_ids": 1, "data.type": 1,
	"data.subtype": 1,
}

func ItemFindOneById(subject string, item_id bson.ObjectId) TaggedItem {
	dbName, found := SubjectName2DbName[subject]
	if !found {
		logs.Error(fmt.Sprintf("无效的学科: %s", subject))
		// 异常处理
	}
	db := MongoDbClient.GetDb(dbName)
	var ret TaggedItem
	_ = db.C("items").FindId(item_id).Select(TaggingFileds).One(&ret)
	return ret
}

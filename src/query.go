package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Query struct {
	db *gorm.DB
}

func (q *Query) InsertData(data Data) {
	q.db.Create(&data)
}

func (q *Query) InsertExperiment(ex Experiment) {
	q.db.Create(&ex)
}

func (q *Query) InsertUser(user User) {
	q.db.Create(&user)
}

func (q *Query) InsertOriginData(data OriginData) {
	q.db.Create(&data)
}

func (q *Query) InsertFeature(f Feature) {
	q.db.Create(&f)
}

func (q *Query) SelectDataFromId(id int) Data {
	var data Data
	q.db.First(&data, id)
	return data
}

func (q *Query) SelectDataFromExperimentId(id int) []Data {
	var datas []Data
	q.db.Find(&datas, "ex_id = ?", id)
	return datas
}

func (q *Query) SelectExperimentFromModelName(name string) Experiment {
	var ex Experiment
	q.db.First(&ex, "model_name = ?", name)
	return ex
}

func (q *Query) SelectExperimentFromIdToModelName(id int) string {
	var ex Experiment
	q.db.First(&ex, id)
	return ex.ModelName
}

func (q *Query) SelectExperimentFromModelNameToId(name string) int {
	var ex Experiment
	q.db.First(&ex, "model_name = ?", name)
	return int(ex.ID)
}

func (q *Query) SelectUserFromIdToName(id int) string {
	var user User
	q.db.First(&user, id)
	return user.Name
}

func (q *Query) SelectUserFromNameToId(name string) int {
	var user User
	q.db.First(&user, "name = ?", name)
	return int(user.ID)
}

func (q *Query) SelectOriginDataFromIdToData(id int) string {
	var data OriginData
	q.db.First(&data, id)
	return data.Data
}

func (q *Query) SelectOriginDataFromDataToId(data string) int {
	var origin OriginData
	q.db.First(&origin, "data = ?", data)
	return int(origin.ID)
}

func (q *Query) SelectFeatureFromIdToName(id int) string {
	var f Feature
	q.db.First(&f, id)
	return f.Name
}

func (q *Query) SelectFeatureFromNameToId(name string) int {
	var f Feature
	q.db.First(&f, "name = ?", name)
	return int(f.ID)
}

func (q *Query) UpdateExperimentDataNum(name string, num int) {
	var ex Experiment
	q.db.Model(&ex).Where("model_name = ?", name).Update("data_num", num)
}

func (q *Query) IsExperimentModelName(name string) bool {
	var count int
	var ex Experiment

	q.db.Where("model_name = ?", name).Find(&ex).Count(&count)

	if count != 0 {
		return true
	}
	return false
}

func (q *Query) IsUser(name string) bool {
	var count int
	var user User

	q.db.Where("name = ?", name).Find(&user).Count(&count)

	if count != 0 {
		return true
	}
	return false
}

func (q *Query) IsOriginData(origin string) bool {
	var count int
	var data OriginData
	q.db.Where("data = ?", origin).Find(&data).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func (q *Query) IsFeature(name string) bool {
	var count int
	var f Feature
	q.db.Where("name = ?", name).Find(&f).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func Init() {
	var db *gorm.DB
	db, err := gorm.Open("mysql", "user:pass@tcp(mysql:3306)/db")
	if err != nil {
		panic(err)
	}

	d := &Data{}
	o := &OriginData{}
	u := &User{}
	e := &Experiment{}
	f := &Feature{}

	db.CreateTable(d)
	db.CreateTable(o)
	db.CreateTable(u)
	db.CreateTable(e)
	db.CreateTable(f)

	defer db.Close()
}

package main

import (
	"time"

	"github.com/Yamashou/elm"
)

func (q *Query) SaveData(d Learn) {
	q.SaveUser(d)
	q.SaveOriginDatas(d)
	q.SaveFeature(d)
	q.AddDataNum(d.ModelName, len(d.OriginDatat))
	data := q.NewData(d)
	for _, v := range data {
		q.InsertData(v)
	}
}

func (q *Query) SaveExperiment(d Learn) {
	q.SaveUser(d)
	q.SaveOriginDatas(d)
	q.SaveFeature(d)
	if !q.IsExperimentModelName(d.ModelName) {
		var ex Experiment
		ex.ModelName = d.ModelName
		ex.Description = d.Description
		ex.DataNum = 0
		ex.H = d.H
		ex.Seed = time.Now().UnixNano()
		q.InsertExperiment(ex)
	}
}

func (q *Query) SaveUser(d Learn) {
	if !q.IsUser(d.User) {
		var user User
		user.Name = d.User
		q.InsertUser(user)
	}
}

func (q *Query) SaveOriginDatas(d Learn) {
	od := removeDuplicate(d.OriginDatat)
	for _, v := range od {
		if !q.IsOriginData(v) {
			var data OriginData
			data.Data = v
			q.InsertOriginData(data)
		}
	}
}

func (q *Query) SaveFeature(d Learn) {
	if !q.IsFeature(d.FeatureName) {
		q.InsertFeature(Feature{Name: d.FeatureName})
	}
}

func (l *Learning) SaveModel(e elm.ELM) {
	e.MarshalBinaryTo(l.ModelName)
}

func removeDuplicate(args []string) []string {
	results := make([]string, 0, len(args))
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}

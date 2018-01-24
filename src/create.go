package main

import (
	"github.com/Yamashou/elm"
	"github.com/Yamashou/floatListtoBinary"
)

func (q *Query) NewData(r Learn) []Data {
	datas := make([]Data, len(r.OriginDatat))
	oData := q.GetOriginDataId(r.OriginDatat)
	for i, v := range r.Feature {
		datas[i].Source = floatListtoBinary.MarshalBinary(v)
		datas[i].FeatureID = q.GetFeatureId(r.FeatureName)
		datas[i].FeatuerNum = r.FeatureNum
		datas[i].CorrectAnswerNum = r.CorrectAnswerNum
		datas[i].OriginDataID = oData[i]
		datas[i].UID = q.GetUserId(r.User) // koko

		datas[i].ExID = q.GetExperimentId(r.ModelName)
	}

	return datas
}

func (l *Experiment) NewDataSet(d []Data) elm.DataSet {
	datas := make([][]float64, len(d))
	for i, v := range d {
		datas[i] = v.Unmarshal()
	}
	e, err := elm.NewDataSet(datas, d[0].FeatuerNum, d[0].CorrectAnswerNum)
	if err != nil {
		panic(err)
	}
	return e
}

func (d *Data) Unmarshal() []float64 {
	return floatListtoBinary.UnmarshalBinary(d.Source)
}

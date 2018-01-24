package main

import (
	"github.com/Yamashou/elm"
)

func (r *Experiment) Training(d []Data) elm.ELM {
	trainingDataSet := r.NewDataSet(d)
	e := elm.ELM{}
	e.Fit(&trainingDataSet, r.H, r.Seed)
	return e
}

func (r *Experiment) Result(d []float64) int {
	model, err := elm.UnmarshalBinaryFromName(r.ModelName)
	if err != nil {
		panic(err)
	}
	ans := model.GetResult(d)
	return ans
}

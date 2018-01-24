package main

func (q *Query) AddDataNum(name string, num int) {
	if q.IsExperimentModelName(name) {
		ex := q.SelectExperimentFromModelName(name)
		q.UpdateExperimentDataNum(name, num+ex.DataNum)
	}
}

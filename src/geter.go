package main

func (q *Query) GetOriginDataId(d []string) []int {
	idList := make([]int, len(d))
	for i, v := range d {
		idList[i] = q.SelectOriginDataFromDataToId(v)
	}
	return idList
}

func (q *Query) GetUserId(name string) int {
	return q.SelectUserFromNameToId(name)
}

func (q *Query) GetFeatureId(name string) int {
	return q.SelectFeatureFromNameToId(name)
}

func (q *Query) GetAllExperiment() []Experiment {
	return q.SelectAllExperiment()
}

func (q *Query) GetExperimentId(name string) int {
	return q.SelectExperimentFromModelNameToId(name)
}

func (q *Query) GetExperimentFromModelName(name string) Experiment {
	return q.SelectExperimentFromModelName(name)
}

func (q *Query) GetDatasFromExperimentId(exId int) []Data {
	return q.SelectDataFromExperimentId(exId)
}

func (q *Query) GetDatasFromExperimentModelName(name string) []Data {
	return q.SelectDataFromExperimentId(q.SelectExperimentFromModelNameToId(name))
}

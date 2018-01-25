package main

import "encoding/json"

type Learn struct {
	ModelName        string      `json:"model_name"`
	Description      string      `json:"description"`
	FeatureName      string      `json:"feature_name"`
	User             string      `json:"user"`
	OriginDatat      []string    `json:"origin_data"`
	Feature          [][]float64 `json:"feature"`
	FeatureNum       int         `json:"feature_num"`
	CorrectAnswerNum int         `json:"correct_answer_num"`
	H                int         `json:"H"`
}

type Learning struct {
	ModelName string `json:"model_name"`
}

type Message struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Ans struct {
	Ans int `json:"ans"`
}

type ExperimentJson struct {
	ModelName   string `json:"model_name"`
	Description string `json:"description"`
	DataNum     int    `json:"data_num"`
	H           int    `json:"h"`
}

func UnmarshalLearn(data []byte) (Learn, error) {
	var r Learn
	err := json.Unmarshal(data, &r)
	return r, err
}
func UnmarshalLearning(data []byte) (Learning, error) {
	var r Learning
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Learn) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Learning) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

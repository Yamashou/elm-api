package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Data struct {
	gorm.Model
	Source           []byte
	FeatureID        int
	FeatuerNum       int
	CorrectAnswerNum int
	OriginDataID     int
	UID              int
	ExID             int
}

type Experiment struct {
	gorm.Model
	ModelName   string
	Description string
	DataNum     int
	H           int
	Seed        int64
}

type User struct {
	gorm.Model
	Name string
}

type OriginData struct {
	gorm.Model
	Data string
}

type Feature struct {
	gorm.Model
	Name string
}

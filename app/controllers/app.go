package controllers

import (
	"encoding/gob"
	"fmt"
	"os"

	. "github.com/jbrukh/bayesian"
	"github.com/revel/revel"
)

var (
	lastname       string
	lastprediction int
	maleNames      []string
	femaleNames    []string
	crar           string
	likely         int
)

const (
	male   Class = "Male"
	female Class = "Female"
	file         = "names.gob"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "GenderAssumer"
	return c.Render(greeting)
}

func (c App) Classify(myName string) revel.Result {
	classifier := NewClassifier(male, female)
	classifier.Learn(maleNames, male)
	classifier.Learn(femaleNames, female)
	lastname = myName
	_, likely, _ = classifier.LogScores([]string{myName})
	if likely == 1 {
		crar = "female"
	} else {
		crar = "male"
	}
	return c.Render(myName, crar)
}

func (c App) TrainYes() revel.Result {
	return c.Render()
}

func (c App) TrainNo() revel.Result {
	if likely == 0 {
		femaleNames = append(femaleNames, lastname)
	} else {
		maleNames = append(maleNames, lastname)
	}
	fmt.Println(maleNames)
	fmt.Println(femaleNames)
	return c.Render()
}

// Encode via Gob to file
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

package link

import (
	"MiRolls/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

/* type data struct { //Return's data type
	Questions      []*Question `json:"questions"`
	Title          string      `json:"title"`
	AnswerOfNumber int         `json:"answerOfNumber"`
}

type Question struct { //Question. return data.
	Type string `json:"type"`
	//OptionsNumber int    `json:"optionsNumber,omitempty"`
	//Placeholder   string `json:"placeholder,omitempty"`
	Title string `json:"title"`
	//Option       []*Options `json:"options,omitempty"`
	//Answer        []string   `json:"answer,omitempty"`
	Answer interface{} `json:"answer"`
}

type Options struct {
	Title  string `json:"title"`
	Option string `json:"option"`
	//Selectivity []int    `json:"selectivity,omitempty"`
	NumberOfSelect int `json:"numberOfSelect"`
}*/ //old data
type data struct {
	Title          string `json:"title"`
	AnswerOfNumber int    `json:"answerOfNumber"`
	Questions      []struct {
		Type   string        `json:"type"`
		Title  string        `json:"title"`
		Answer []interface{} `json:"answer"`
	} `json:"questions"`
}

type radio struct {
	Option         string `json:"option"`
	NumberOfSelect int    `json:"numberOfSelect"`
}

type roll struct {
	//Id string `db:"`
	Roll string `db:"roll"`
}

type dataFromFront struct {
	Link string `json:"link"`
}

type Answer struct {
	Answer string `db:"answer"`
}

type dbRoll struct {
	Title string `json:"title"`
	Quest []struct {
		Type          string   `json:"type"`
		OptionsNumber int      `json:"optionsNumber,omitempty"`
		Title         string   `json:"title"`
		Options       []string `json:"options,omitempty"`
		Placeholder   string   `json:"placeholder,omitempty"`
	} `json:"quest"`
}

type AnswerFromJson struct {
	Link   string `json:"link"`
	Answer []struct {
		Title  string      `json:"title"`
		Type   string      `json:"type"`
		Answer interface{} `json:"answer"`
	} `json:"answer"`
}

//goland:noinspection ALL
func QueryRoll(r *gin.Engine) {
	r.POST("/query/roll", func(c *gin.Context) {
		body, _ := c.GetRawData()
		dataFront := new(dataFromFront)
		_ = json.Unmarshal(body, &dataFront)

		mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
		sql, err := sqlx.Open("mysql", mysql)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error(), "error": err.Error(), "errorType": "database connect error"})
			//log.Fatal("[FATAL ERROR]Cannot connect database")
			return
		}

		var roll roll
		err = sql.Get(&roll, "SELECT `roll` FROM `rolls` WHERE `link`=?")
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}
		//handle data

		data := new(data)
		//create return data

		//init data
		rollStruct := new(dbRoll)
		json.Unmarshal([]byte(roll.Roll), &rollStruct)
		data.Title = rollStruct.Title

		var answers []Answer
		err = sql.Select(&answers, "SELECT `answer` FROM `answer` WHERE `link`=?", dataFront.Link)
		data.AnswerOfNumber = len(answers)

		//Add questions to data
		for i := 0; i < len(rollStruct.Quest); i++ {
			data.Questions[i].Title = rollStruct.Quest[i].Title
			data.Questions[i].Type = rollStruct.Quest[i].Type
			if data.Questions[i].Type == "radio" || data.Questions[i].Type == "multipleChoice" {
				//append(data.Questions[i].Answer)
				for i2 := 0; i2 < len(rollStruct.Quest[i].Options); i2++ {
					ra := radio{
						Option:         rollStruct.Quest[i].Options[i2],
						NumberOfSelect: 0,
					}
					data.Questions[i].Answer = append(data.Questions[i].Answer, ra)
				}
			}
		}

		// Use for(){} to loop database.
		for i := 0; i < len(answers); i++ {
			// loop length for answer.
			//dbMap := make(map[string]interface{})
			dbStruct := new(AnswerFromJson)
			err := json.Unmarshal([]byte(answers[i].Answer), &dbStruct)
			if err != nil {
				c.JSON(500, gin.H{
					"error":   "Can't unmarshal json. " + err.Error(),
					"message": "error",
				})
				return
			} // Use json.Unmarshal to parse.
			data.Questions[i].Title = dbStruct.Answer[i].Title
			data.Questions[i].Type = dbStruct.Answer[i].Type
		}
		// to link rolls

		c.JSON(200, gin.H{
			"message": "success",
			"data":    *data,
		})
		_ = sql.Close()
	})
}

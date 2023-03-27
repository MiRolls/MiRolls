package link

import (
	"MiRolls/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type data struct {
	Title          string         `json:"title"`
	AnswerOfNumber int            `json:"answerOfNumber"`
	Questions      []dataQuestion `json:"questions"`
}

type dataQuestion struct {
	Type   string        `json:"type"`
	Title  string        `json:"title"`
	Answer []interface{} `json:"answer"`
}

type radio struct {
	Option         string `json:"option"`
	NumberOfSelect int    `json:"numberOfSelect"`
}

type roll struct {
	//Id string `db:"`
	Roll string `db:"roll"`
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
		Title  string        `json:"title"`
		Type   string        `json:"type"`
		Answer []interface{} `json:"answer"`
	} `json:"answer"`
}

type dataFromFront2 struct {
	//Link string `json:"link,omitempty"`
	Code string `json:"code"`
}

func (ra *radio) AddNumberOfSelect() {
	ra.NumberOfSelect++
}

//goland:noinspection ALL
func QueryRoll(r *gin.Engine) {
	r.POST("/query/roll", func(c *gin.Context) {
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error(), "error": err.Error()})
			//log.Fatal("[FATAL ERROR]Cannot connect database")
			return
		}
		dataFront := new(dataFromFront2)
		err = json.Unmarshal(body, &dataFront)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error(), "error": err.Error()})
			//log.Fatal("[FATAL ERROR]Cannot connect database")
			return
		}

		mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
		sql, err := sqlx.Open("mysql", mysql)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error(), "error": err.Error(), "errorType": "database connect error"})
			//log.Fatal("[FATAL ERROR]Cannot connect database")
			return
		}

		var roll roll
		err = sql.Get(&roll, "SELECT `roll` FROM `rolls` WHERE `code`=?", dataFront.Code)
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
		err = sql.Select(&answers, "SELECT `answer` FROM `answer` WHERE `link`=?", dataFront.Code)
		data.AnswerOfNumber = len(answers)

		//Add questions to data
		for i := 0; i < len(rollStruct.Quest); i++ {
			data.Questions = append(data.Questions, dataQuestion{
				Type:  rollStruct.Quest[i].Type,
				Title: rollStruct.Quest[i].Title,
			})
			//data.Questions[i].Title = rollStruct.Quest[i].Title
			//data.Questions[i].Type = rollStruct.Quest[i].Type
			if data.Questions[i].Type == "radio" || data.Questions[i].Type == "multipleChoice" {
				//append(data.Questions[answerNumber].Answer)
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
		for answerNumber := 0; answerNumber < len(answers); answerNumber++ {
			// loop length for answer.
			//dbMap := make(map[string]interface{})
			dbStruct := new(AnswerFromJson)
			err := json.Unmarshal([]byte(answers[answerNumber].Answer), &dbStruct)
			if err != nil {
				c.JSON(500, gin.H{
					"error":   "Can't unmarshal json. " + err.Error(),
					"message": "error",
				})
				return
			} // Use json.Unmarshal to parse.

			//Loop every question
			for questionNumber := 0; questionNumber < len(dbStruct.Answer); questionNumber++ {
				//2 condition
				if data.Questions[questionNumber].Type == "radio" || data.Questions[questionNumber].Type == "multipleChoice" {
					for optionNumber := 0; optionNumber < len(dbStruct.Answer[questionNumber].Answer); optionNumber++ {
						//dbStruct..[questionNumber].Answer[optionNumber]
						if dbStruct.Answer[questionNumber].Answer[optionNumber] == true {
							//data.Questions[questionNumber].Answer[optionNumber].numberOfSelect++
							radio, isOK := data.Questions[questionNumber].Answer[optionNumber].(radio)
							if isOK {
								radio.AddNumberOfSelect()
							} else {
								c.JSON(500, gin.H{
									"error":   "data has error!",
									"message": "error",
								})
							}
						}
					}
				} else {
					// is blank
					data.Questions[questionNumber].Answer = append(data.Questions[questionNumber].Answer, dbStruct.Answer[questionNumber].Answer[0])
				}
			}
			// to link rolls

			c.JSON(200, gin.H{
				"message": "success",
				"data":    data,
			})
			fmt.Println("success")
			err = sql.Close()
			if err != nil {
				c.JSON(500, gin.H{"message": err.Error(), "error": err.Error()})
				//log.Fatal("[FATAL ERROR]Cannot connect database")
				return
			}
		}
	})
}

package link

import (
	"MiRolls/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type data struct { //Return's data type
	Questions      []*Question `json:"questions"`
	Title          string      `json:"title"`
	AnswerOfNumber int         `json:"answerOfNumber"`
}

type Question struct { //Question. return data.
	Type          string     `json:"type"`
	OptionsNumber int        `json:"optionsNumber,omitempty"`
	Placeholder   string     `json:"placeholder,omitempty"`
	Title         string     `json:"title"`
	Options       []*Options `json:"options,omitempty"`
	Answer        []string   `json:"answer,omitempty"`
}

type Options struct {
	Options string `json:"options"`
	//Selectivity []int    `json:"selectivity,omitempty"`
	NumberOfSelect int `json:"NumberOfSelect"`
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
	Quest any    `json:"quest"`
}

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
		//data.Title =

		var answers []Answer
		err = sql.Select(&answers, "SELECT `answer` FROM `answer` WHERE `link`=?", dataFront.Link)
		// Use for(){} to loop database.
		for i := 0; i < len(answers); i++ {
			// loop length for answer.
		}
		// to link rolls

		c.JSON(200, gin.H{
			"message": "success",
			"data":    *data,
		})
		_ = sql.Close()
	})
}

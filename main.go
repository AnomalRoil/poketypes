package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

const (
	Normal int = iota
	Fire
	Water
	Electric
	Grass
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
	Fairy
)

type float float32

var chart [][]float
var typeName map[int]string
var fromName map[string]int

func init() {
	chart = make([][]float, 18)
	typeName = make(map[int]string, 18)
	fromName = make(map[string]int, 18)

	chart[Normal] = []float{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.5, 0, 1, 1, 0.5, 1}
	chart[Fire] = []float{1, 0.5, 0.5, 1, 2, 2, 1, 1, 1, 1, 1, 2, 0.5, 1, 0.5, 1, 2, 1}
	chart[Water] = []float{1, 2, 0.5, 1, 0.5, 1, 1, 1, 2, 1, 1, 1, 2, 1, 0.5, 1, 1, 1}
	chart[Electric] = []float{1, 1, 2, 0.5, 0.5, 1, 1, 1, 0, 2, 1, 1, 1, 1, 0.5, 1, 1, 1}
	chart[Grass] = []float{1, 0.5, 2, 1, 0.5, 1, 1, 0.5, 2, 0.5, 1, 0.5, 2, 1, 0.5, 1, 0.5, 1}
	chart[Ice] = []float{1, 0.5, 0.5, 1, 2, 0.5, 1, 1, 2, 2, 1, 1, 1, 1, 2, 1, 0.5, 1}
	chart[Fighting] = []float{2, 1, 1, 1, 1, 2, 1, 0.5, 1, 0.5, 0.5, 0.5, 2, 0, 1, 2, 2, 0.5}
	chart[Poison] = []float{1, 1, 1, 1, 2, 1, 1, 0.5, 0.5, 1, 1, 1, 0.5, 0.5, 1, 1, 0, 2}
	chart[Ground] = []float{1, 2, 1, 2, 0.5, 1, 1, 2, 1, 0, 1, 0.5, 2, 1, 1, 1, 2, 1}
	chart[Flying] = []float{1, 1, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 2, 0.5, 1, 1, 1, 0.5, 1}
	chart[Psychic] = []float{1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 0.5, 1, 1, 1, 1, 0, 0.5, 1}
	chart[Bug] = []float{1, 0.5, 1, 1, 2, 1, 0.5, 0.5, 1, 0.5, 2, 1, 1, 0.5, 1, 2, 0.5, 0.5}
	chart[Rock] = []float{1, 2, 1, 1, 1, 2, 0.5, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 0.5, 1}
	chart[Ghost] = []float{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 1, 1}
	chart[Dragon] = []float{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 0.5, 0}
	chart[Dark] = []float{1, 1, 1, 1, 1, 1, 0.5, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 1, 0.5}
	chart[Steel] = []float{1, 0.5, 0.5, 0.5, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 0.5, 2}
	chart[Fairy] = []float{1, 0.5, 1, 1, 1, 1, 2, 0.5, 1, 1, 1, 1, 1, 1, 2, 2, 0.5, 1}

	typeName[Normal] = "Normal"
	typeName[Fire] = "Fire"
	typeName[Water] = "Water"
	typeName[Electric] = "Electric"
	typeName[Grass] = "Grass"
	typeName[Ice] = "Ice"
	typeName[Fighting] = "Fighting"
	typeName[Poison] = "Poison"
	typeName[Ground] = "Ground"
	typeName[Flying] = "Flying"
	typeName[Psychic] = "Psychic"
	typeName[Bug] = "Bug"
	typeName[Rock] = "Rock"
	typeName[Ghost] = "Ghost"
	typeName[Dragon] = "Dragon"
	typeName[Dark] = "Dark"
	typeName[Steel] = "Steel"
	typeName[Fairy] = "Fairy"
	fromName["Normal"] = Normal
	fromName["Fire"] = Fire
	fromName["Water"] = Water
	fromName["Electric"] = Electric
	fromName["Grass"] = Grass
	fromName["Ice"] = Ice
	fromName["Fighting"] = Fighting
	fromName["Poison"] = Poison
	fromName["Ground"] = Ground
	fromName["Flying"] = Flying
	fromName["Psychic"] = Psychic
	fromName["Bug"] = Bug
	fromName["Rock"] = Rock
	fromName["Ghost"] = Ghost
	fromName["Dragon"] = Dragon
	fromName["Dark"] = Dark
	fromName["Steel"] = Steel
	fromName["Fairy"] = Fairy
}

func Vulnerable(t int) []string {
	ret := []string{}
	for i, v := range chart {
		if v[t] > 1 {
			ret = append(ret, typeName[i])
		}
	}
	return ret
}

func NotVulnerable(t int) []string {
	ret := []string{}
	for i, v := range chart {
		if v[t] < 1 && v[t] > 0 {
			ret = append(ret, typeName[i])
		}
	}
	return ret
}

func Invulnerable(t int) []string {
	ret := []string{}
	for i, v := range chart {
		if v[t] == 0 {
			ret = append(ret, typeName[i])
		}
	}
	return ret
}

func Ineffective(t int) []string {
	ret := []string{}
	for i, v := range chart[t] {
		if v == 0 {
			ret = append(ret, typeName[i])
		}
	}
	return ret
}

func NotEffective(t int) []string {
	ret := []string{}
	for i, v := range chart[t] {
		if v < 1 && v > 0 {
			ret = append(ret, typeName[i])
		}
	}
	return ret
}

func VeryEffective(t int) []string {
	ret := []string{}
	for i, v := range chart[t] {
		if v > 1 {
			ret = append(ret, typeName[i])
		}
	}
	return ret
}

type PokeType struct {
	Name string `uri:"name" binding:"required"`
	Lang string `uri:"lang" binding:"required"`
}

func GetTypes() []string {
	ret := []string{}
	for i := 0; i < len(typeName); i++ {
		ret = append(ret, typeName[i])
	}
	return ret
}

func main() {
	route := gin.Default()
	route.SetTrustedProxies(nil)
	route.SetFuncMap(template.FuncMap{
		"sliceContains": slices.Contains[string],
	})
	route.LoadHTMLFiles("./home.tmpl", "./entry.tmpl")
	route.Static("/img", "./types")
	route.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/en")
	})
	route.GET("/fr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"valid": GetTypes(),
			"lang":  "fr",
		})
	})
	route.GET("/en", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"valid": GetTypes(),
			"lang":  "en",
		})
	})
	route.GET("/:lang/:name", func(c *gin.Context) {
		var t PokeType
		if err := c.ShouldBindUri(&t); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		if v, ok := fromName[t.Name]; !ok || (t.Lang != "en" && t.Lang != "fr") {
			c.JSON(400, gin.H{
				"valid_types":    typeName,
				"valid_language": []string{"en", "fr"},
				"got_lang":       t.Lang,
				"got_type":       t.Name,
			})
			return
		} else {
			c.HTML(http.StatusOK, "entry.tmpl", gin.H{
				"lang":          t.Lang,
				"name":          t.Name,
				"veryEffective": VeryEffective(v),
				"notEffective":  NotEffective(v),
				"ineffective":   Ineffective(v),
				"vulnerable":    Vulnerable(v),
				"invulnerable":  Invulnerable(v),
				"resistant":     NotVulnerable(v),
			})
		}
	})
	route.Run(":8088")
}

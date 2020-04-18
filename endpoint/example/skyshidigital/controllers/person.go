package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-API/endpoint/example/skyshidigital/structs"
)

// GetPerson : to get one data with {id}
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// GetPersons : get all data in person
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// CreatePerson : create new data to the database
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	person.FirstName = firstName
	person.LastName = lastName
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

// UpdatePerson : update data with {id} as query
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newPerson.FirstName = firstName
	newPerson.LastName = lastName
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// DeletePerson : delete data with {id}
func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}

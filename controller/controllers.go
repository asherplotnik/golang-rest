package controller

import (
	"encoding/json"
	"errors"
	"github.com/asherplotnik/golang-rest/model"
	"github.com/asherplotnik/golang-rest/service"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	result, err:= service.GetAllTodos()
	if err!= nil {
		if err.(model.TodoError).Code == 500 {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		} else {
			context.AbortWithStatus(204)
			return
		}
	}
	context.IndentedJSON(http.StatusOK, result)
}

func DeleteAllTodo(context *gin.Context) {
	deleteCount, err := service.DeleteAllTodos()
	if err != nil {
		context.IndentedJSON(err.(model.TodoError).Code, gin.H{"message": err.(model.TodoError).Message})
		return
	}
	context.IndentedJSON(http.StatusOK, deleteCount)
}

func PostTodo(context *gin.Context) {
	var newTodo model.Todo
	err := context.BindJSON(&newTodo)
	if  err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	result, err := service.AddOneTodo(newTodo)
	if err != nil {
		context.IndentedJSON(err.(model.TodoError).Code, gin.H{"message": err.(model.TodoError).Message})
		return
	}
	context.IndentedJSON(http.StatusCreated, result)
}

func GetTodo(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	result, err := service.GetTodo(id)

	if err != nil {
		context.IndentedJSON(err.(model.TodoError).Code, gin.H{"message": err.(model.TodoError).Message})
		return
	}

	webResponse, _ := makeWebRequest(result.Item)
	if webResponse != nil {
		context.IndentedJSON(http.StatusOK, webResponse)
		return
	}

	context.IndentedJSON(http.StatusOK, result)

}

func UpdateTodo(context *gin.Context) {
	var newTodo model.Todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest , gin.H{"message": "bad request"})
		return
	}

	result, err := service.UpdateTodo(newTodo)
	
	if err != nil {
		context.IndentedJSON(err.(model.TodoError).Code, gin.H{"message": err.(model.TodoError).Message})
		return
	}

	context.IndentedJSON(http.StatusOK, result)
}

func DeleteTodo (context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	result, err := service.DeleteTodo(id)

	if err != nil {
		context.IndentedJSON(err.(model.TodoError).Code, gin.H{"message": err.(model.TodoError).Message})
		return
	}

	context.IndentedJSON(http.StatusOK, result)
}

func makeWebRequest(url string) (interface{}, error) {
	response, err := http.Get(url)
	
	if err != nil {
		return nil, errors.New("error calling url")
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("error read response")
	}
	var iot interface{}
	mErr := json.Unmarshal(bodyBytes, &iot)
	
	if mErr != nil {
        return nil, errors.New("error Unmarshal response")
    }
	return iot, nil
}

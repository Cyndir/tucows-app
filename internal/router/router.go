package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Cyndir/tucows-app/internal/database"
	"github.com/Cyndir/tucows-app/internal/messagequeue"
	"github.com/Cyndir/tucows-app/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Router struct {
	db     database.Database
	mq     messagequeue.MessageQueue
	Engine *gin.Engine
}

func New(db database.Database, mq messagequeue.MessageQueue) Router {
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	return Router{db: db, mq: mq}
}
func (r *Router) Setup() {
	engine := gin.Default()
	engine.GET("/order/:id", r.GetOrder)
	engine.POST("/order", r.PostOrder)
	engine.PATCH("/order", r.UpdateOrder)
	r.Engine = engine
}
func (r Router) Run() {
	r.Engine.Run(":8081")
}

func (r Router) GetOrder(c *gin.Context) {
	log.Println("received get request")
	order, err := r.db.GetOrder(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, order)
}

func (r Router) PostOrder(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	log.Println("received post request")

	var order model.Order
	dec := json.NewDecoder(c.Request.Body)
	err := dec.Decode(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	order.Status = "pending"
	order.ID = xid.New().String()

	err = r.db.InsertOrder(&order)
	if err != nil {
		log.Printf("db call failed: %s", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	err = r.mq.Publish(order)
	if err != nil {
		log.Println("message queue send failed")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": order.ID,
	})

}

func (r Router) UpdateOrder(c *gin.Context) {
	log.Println("received patch request")
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	var order model.Order
	dec := json.NewDecoder(c.Request.Body)
	err := dec.Decode(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err = r.db.UpdateOrder(order.ID, order.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
}

package controllers

import (
	"VisitorsManagementSystem/dao"
	"VisitorsManagementSystem/logic"
	"VisitorsManagementSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
func AddVisitor(c *gin.Context) {
/*	visitorName := c.PostForm("visitor_name")
	sex := c.PostForm("sex")
	phone := c.PostForm("phone")
	visitId := c.PostForm("visit_id")
	event := c.PostForm("event")
*/
	var visitor visitors

	if err := c.Bind(&visitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	if len(visitor.VisitorName) == 0{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message" : "用户名不能为空",
		})
		return
	}

	if visitor.Sex != "男" && visitor.Sex != "女"{
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message" : "性别只能为男女",
			})
			return
		}

	if len(visitor.Phone) != 11{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "电话号码有误",
		})
		return
	}


	if len(visitor.VisitId) != 18{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "身份证有误",
		})
		return
	}

	if len(visitor.Event) == 0 {
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":    422,
			"message": "事由不能为空",
		})
		return
	}

	newVisitor := models.Visitor{
		VisitorName: visitor.VisitorName,
		Sex: visitor.Sex,
		Phone: visitor.Phone,
		VisitId: visitor.VisitId,
		Event: visitor.Event,
	}
	dao.InitDB().Create(&newVisitor)


	//返回结果
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"success",
	})
}



func AllVisitors(c *gin.Context) {
	result := logic.AllData()

	var users []User

	for i:=0;i<=len(result)-1;i++{
		visistorName := result[i].VisitorName
		createdAt := result[i].CreatedAt
		event := result[i].Event
		var u User
		u = User{VisitorName: visistorName,CreatedAt: createdAt,Event: event}

		users = append(users,u)
	}


	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"results":users,
		"message":"success",
	})
}

type visitors struct {

	//Id uint `gorm:"primary key"`
	VisitorName string `gorm:"varchar(20);not null" form:"visitor_name" json:"visitor_name"`
	Sex string `gorm:"enum('男','女');not null" form:"sex" json:"sex"`
	Phone string `gorm:"varchar(20);not null" form:"phone" json:"phone"`
	VisitId string `gorm:"varchar(20);not null" form:"visit_id" json:"visit_id"`
	Event string `gorm:"varchar(100);not null" form:"event" json:"event"`
	//VisitTime time.Time
}

type User struct {
	VisitorName string `form:"visitor_name" json:"visitor_name"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	Event string `form:"event" json:"event"`

}






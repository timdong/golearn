package main

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"net/http"
	"strconv"
)

// Model Struct
type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123465@/test?charset=utf8", 30)
}

func main() {

	orm.Debug = true
	router := gin.Default()

	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	if err != nil {
		router.GET("/error", func(c *gin.Context) {
			c.String(http.StatusOK, err.Error())
		})
	}

	// update
	user.Name = "astaxie"
	id, err = o.Update(&user)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)

	// delete
	//num, err = o.Delete(&u)

	s := strconv.FormatInt(id, 16)
	//fmt.Printf("%T, %v\n", s, s)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello slene id = "+s+"\n Name = "+u.Name)
	})

	//fmt.Print(string(id))

	router.Run(":8080")
}

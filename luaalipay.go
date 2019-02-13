package main

import (
	"github.com/gin-gonic/gin"
	"luaalipay/controllers"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			//fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())

	index := new(controllers.IndexController)

	//不要token的接口
	public := r.Group("/")
	{
		public.GET("SaveImei", index.SaveImei)
		public.GET("SaveImei2", index.SaveImei2)
		public.GET("GetPhone", index.GetPhone)
		public.GET("GetPhone2", index.GetPhone2)
		public.GET("GetNeedAddFriend", index.GetNeedAddFriend)
		public.GET("UpdateNeedAddfriend", index.UpdateNeedAddfriend)
		public.GET("UpdateCheckRule", index.UpdateCheckRule)
		public.GET("SaveCheckPhone", index.SaveCheckPhone)
	}

	r.Run("0.0.0.0:8905")
}

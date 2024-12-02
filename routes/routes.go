package routes

import (
	"delivery/models"
	"github.com/gin-gonic/gin"
)


func InitializeRoutes(r *gin.Engine) {

	r.GET("/products", func(c *gin.Context) {

		products := models.GetAllProducts()


		c.HTML(200, "products.html", products)
	})
	r.GET("/form", func(c *gin.Context) {
		c.HTML(200, "form.html", nil)
	})
	
	r.POST("/submit-product", func(c *gin.Context) {
		model := c.PostForm("name")
		company := c.PostForm("company")
		price := c.PostForm("price")
		
		err := models.AddProduct(model, company, price)
		if err != nil {
			c.String(500, "Failed to submit product: %v", err)
			return
		}
		
		c.JSON(200, "Product submitted successfully!")
	})
	

}
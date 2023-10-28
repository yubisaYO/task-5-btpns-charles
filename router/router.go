package router

import (
	"task-5-pbi-btpns-charles/controllers"
	"task-5-pbi-btpns-charles/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterSetUp() *gin.Engine {
    r := gin.Default()

    // Route publik (untuk pendaftaran dan masuk)
    public := r.Group("/api")
    {
        public.POST("/users/register", controllers.Register) // Ganti "User" dengan "Register"
        public.POST("/users/login", controllers.Login)
    }

    // Route yang dilindungi oleh middleware otentikasi
    protected := r.Group("/api")
    protected.Use(middlewares.Authenticate())
    {
        // Pengguna
        protected.GET("/users/:id", controllers.GetUserByID)
        protected.PUT("/users/:id", controllers.UpdateUser)
        protected.DELETE("/users/:id", controllers.DeleteUser)

        // Foto-foto
        protected.GET("/photos", controllers.GetAllPhotos)
        protected.GET("/photos/:id", controllers.GetPhotoByID)
        protected.POST("/photos", controllers.CreatePhoto)
        protected.PUT("/photos/:id", controllers.UpdatePhoto)
        protected.DELETE("/photos/:id", controllers.DeletePhoto)
    }

    return r
}

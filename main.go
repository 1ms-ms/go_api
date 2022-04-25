package main
import (
    "net/http"

    "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"context"
    "log"
    "time"
)
type Note struct {
	Id     string
	Title    string
	Body     string
	Date    int
}

func getNotes(c *gin.Context) {
    clientOptions := options.Client().ApplyURI("mongodb+srv://")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	notesCollection := client.Database("notesDB").Collection("notes")
	cursor, err := notesCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var notes []Note
	if err = cursor.All(context.TODO(), &notes); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, notes)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	
}
func addNotes(c *gin.Context) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	notesCollection := client.Database("notesDB").Collection("notes")
	cursor, err := notesCollection.Find(context.TODO(), bson.D{{}})
	_ = cursor
	if err != nil {
		log.Fatal(err)
	}
	body:=Note{}
	if err:=c.ShouldBindJSON(&body);err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"error": "VALIDATEERR-1",
			"message": "Invalid inputs. Please check your inputs"})
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, insertErr := notesCollection.InsertOne(ctx, body)
	_ = result
	_ = insertErr
    c.JSON(http.StatusAccepted,gin.H{"status": "success",})
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
  }

func deleteNotes(c *gin.Context) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	b := c.Param("id")
	if err != nil {
		log.Fatal(err)
	}
	notesCollection := client.Database("notesDB").Collection("notes")
	cursor, err := notesCollection.Find(context.TODO(), bson.D{{}})
	_ = cursor
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := notesCollection.DeleteOne(ctx, bson.M{"id": b})	
	_ = result
	c.JSON(http.StatusAccepted,gin.H{"status": "success",})
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {   
    router := gin.Default()
    router.GET("/", getNotes)
	router.POST("/create", addNotes)
	router.DELETE("/delete/:id", deleteNotes)
    router.Run("localhost:8080")
	
}

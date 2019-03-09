type User struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	SurName string `bson:"surname" json:"surname"`
	Address string `bson:"address" json:"address"`
}

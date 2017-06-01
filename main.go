package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.LoadHTMLGlob("templates/*")
	// db := NewDB("wallet.db")

	// defer db.Close()
	// data := make(map[string]models.Account)
	// data := models.Account{
	// 	UserName:  "mague",
	// 	Email:     "maguedeveloper@gmail.com",
	// 	Password:  "enmanuel",
	// 	WebSite:   "enmanuelmolina.com",
	// 	CreatedAt: time.Now(),
	// }
	// err = db.Save(&data)
	// var rAccount models.Account
	// err = db.One("Email", "enmanueldavidmolina@gmail.com", &rAccount)
	// fmt.Println(rAccount)
	// var rAccounts []models.Account
	// err = db.AllByIndex("", &rAccounts)
	// fmt.Println(rAccounts)
	// err = db.All(&rAccounts)
	// fmt.Println(rAccounts)
	// db.Set("answer", &data, "accounts")
	// fmt.Println("Resultado: " + db.Get("1", "accounts"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			})
	})
	initializeRoutes(router)

	// router.GET("/accounts", apiCtrl.All)
	router.Run()
}

//Database
// type Store interface {
// 	Set(key string, values *models.Account, table string) error
// 	Get(key string, table string) string
// 	// Len() int
// 	Close() error
// }
// type DB struct {
// 	db *bolt.DB
// }
//
// var _ Store = &DB{}
// var (
// 	tableAccounts = []byte("accounts")
// )
//
// func openDatabase(name string) *bolt.DB {
// 	db, err := bolt.Open(name, 0600, nil)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	var tables = [...][]byte{
// 		tableAccounts,
// 	}
//
// 	db.Update(func(tx *bolt.Tx) (err error) {
// 		for _, table := range tables {
// 			_, err = tx.CreateBucketIfNotExists(table)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 		return
// 	})
// 	return db
// }
//
// func NewDB(name string) *DB {
// 	return &DB{
// 		db: openDatabase(name),
// 	}
// }
//
// func (d *DB) Close() error {
// 	return d.db.Close()
// }
// func (d *DB) Set(key string, value *models.Account, table string) error {
// 	d.db.Update(func(tx *bolt.Tx) error {
// 		tableb := []byte(table)
// 		b := tx.Bucket(tableb)
// 		jsonResult, err := json.Marshal(value)
// 		if err != nil {
// 			fmt.Println("Error al generar el json")
// 			// return
// 		}
// 		return b.Put([]byte(value.Id), jsonResult)
// 	})
// 	return nil
// }
//
// func (d *DB) Get(key string, table string) (value string) {
// 	keyb := []byte(key)
// 	d.db.View(func(tx *bolt.Tx) error {
// 		tableb := []byte(table)
// 		b := tx.Bucket(tableb)
// 		c := b.Cursor()
//
// 		for k, v := c.First(); k != nil; k, v = c.Next() {
// 			if bytes.Equal(keyb, k) {
// 				value = string(v)
// 				break
// 			}
// 		}
//
// 		return nil
// 	})
// 	return
// }

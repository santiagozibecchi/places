// Gestion de la coneccion a la DB

package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBData struct {
	Host string
	Port string
	BdName string
	RolName string
	RolPassword string
}

// func init() {
// 	fmt.Println("Iniciando coneccion con la DB")
// 	EstablishDBConnection()
// }

func LoadEnv() (DBData, error) {
	var err error
	if err = godotenv.Load(".env"); err !=nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
		return DBData{}, err
	}

	return DBData{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		BdName: os.Getenv("DB_NAME"),
		RolName: os.Getenv("ROL_NAME"),
		RolPassword: os.Getenv("ROL_PASSWORD"),
	}, nil
}

var Db *sql.DB

func EstablishDBConnection() error {
	dbData, err := LoadEnv()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbData.Host, dbData.Port, dbData.RolName, dbData.RolPassword, dbData.BdName)

	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	Db = dbConn
	fmt.Println("Conectado! ", Db)
	
	if err = Db.Ping(); err != nil {
		Db.Close()
		return err
	}
	
	return nil
}
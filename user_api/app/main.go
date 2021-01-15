package main

import (
	"database/sql"
	"efishery_test/user_api/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	Database struct {
		User     string `env:"DATABASE_USER,required"`
		Password string `env:"DATABASE_PASSWORD,required"`
		Host     string `env:"DATABASE_HOST,default=127.0.0.1"`
		Port     string `env:"DATABASE_PORT,default=3306"`
		Name     string `env:"DATABASE_NAME,default=sejastip"`
		Pool     int    `env:"DATABASE_POOL,default=5"`
		Charset  string `env:"DATABASE_CHARSET,default=utf8"`
	}

	Port string `env:"PORT,required"`

	JWTPrivateKey string `env:"JWT_PRIVATE_KEY,required"`
}

var config Config

func init() {
	gotenv.Load()
	err := envdecode.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func NewMysqlConnection(c *Config) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.Charset,
	)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(c.Database.Pool)
	err = db.Ping()
	return db, err
}

func main() {
	_, err := NewMysqlConnection(&config)

	if err != nil {
		log.Fatal("error connecting to mysql server: ", err)
	}

	h := handler.NewHandler(config.JWTPrivateKey)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      h,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func(s *http.Server) {
		log.Printf("Efishery User API is available at %s\n", s.Addr)
		if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
			log.Fatal(serr)
		}
	}(s)

	<-sigChan

	log.Println("Efishery User API stopped.")
}

package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"projetoPadrao/internal/domain"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func ConnectionDb() (*gorm.DB, error) {
	configFile, err := ioutil.ReadFile("/home/emmanoel/projects/apiProjetoPadrao/database.yaml")
	if err != nil {
		log.Fatalf("Erro ao obter o arquivo de configuração \n error: %v", err)
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error ao decodificar o YAML \n error: %v", err)
		return nil, err
	}

	dns := fmt.Sprintf("host=%v dbname=%v user=%v password=%v port=%v sslmode=disable",
		config.Database.Host, config.Database.DBName, config.Database.User,
		config.Database.Password, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Person{})
	if err != nil {
		log.Println("[connection] Erro ao criar tabelas: " + err.Error())
	}

	return db, nil
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/domjeff/deploy-to-railway/internal/domain"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db     *sql.DB
	dbGorm *gorm.DB
}

func NewPostgresDB(connStr string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check if the database connection is working
	if err := db.Ping(); err != nil {
		return nil, err
	}

	dbGorm, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// dbGorm.Debug().AutoMigrate(domain.Student{}, domain.Product{})

	return &PostgresDB{
		db:     db,
		dbGorm: dbGorm,
	}, nil
}

func (p *PostgresDB) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *PostgresDB) Save(u *domain.User) error {
	if u.ID != 0 {
		// update part
		query := `UPDATE public.users
		SET username=$1,email=$2
		WHERE id=$3;
		`

		_, err := p.db.Query(query, u.UserName, u.Email, u.ID)
		if err != nil {
			return err
		}
		return nil
	}

	// insertion part
	query := `
	INSERT INTO users
	(username, email)
	VALUES($1, $2)
	RETURNING *
	`

	var out domain.User

	err := p.db.QueryRow(query, u.UserName, u.Email).Scan(&out.ID, &out.UserName, &out.Email)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) SaveGormStudent(gormStudent domain.Student) error {
	return p.dbGorm.Create(&gormStudent).Error
}

func (p *PostgresDB) GetGormStudent(id int) (domain.Student, error) {
	var res domain.Student
	err := p.dbGorm.Preload("Products").First(&res, id).Error
	if err != nil {
		return domain.Student{}, err
	}

	return res, nil
}

func (p *PostgresDB) UpdateGormStudent(gormStudent domain.Student) error {
	return p.dbGorm.Model(domain.Student{}).Where("id = ?", gormStudent.ID).Updates(domain.Student{
		Email: gormStudent.Email,
	}).Error
}

func (p *PostgresDB) GetUsers() ([]domain.User, error) {
	query := `
	select * from users
	`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []domain.User

	for rows.Next() {
		var currentRowUser domain.User
		err = rows.Scan(&currentRowUser.ID, &currentRowUser.UserName, &currentRowUser.Email)
		if err != nil {
			return nil, err
		}

		res = append(res, currentRowUser)
	}

	return res, nil
}

func (p *PostgresDB) FindByID(id int) (*domain.User, error) {
	// Implement the logic to find data by ID from PostgreSQL
	return nil, fmt.Errorf("Not implemented")
}

func InitPostgresDB() (*PostgresDB, error) {
	// Load database credentials from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := NewPostgresDB(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}
	return db, nil
}

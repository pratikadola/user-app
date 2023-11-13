package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/user-assignment/app/config"
	"github.com/user-assignment/app/models"
)

type UserDB struct {
	psql *sq.StatementBuilderType
	DB   *sql.DB
}

func NewUserDB(conf config.Config) *UserDB {
	return connect(conf)
}

func connect(conf config.Config) *UserDB {

	username := conf.Database.Username
	// password := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port
	dbName := conf.Database.DBName
	password := conf.Database.Password
	dbUrl := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to db")
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(3)
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &UserDB{&psql, db}
}

func (u *UserDB) CreateUser(user *models.User) (*models.User, error) {
	existing, err := u.GetUser(user.Name)
	if err == nil {
		fmt.Println(existing, err)
		return nil, errors.New("user exists")
	}
	query, args, err := u.psql.Insert("users").Columns("name").Values(user.Name).Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		log.Fatal(err)
	}
	if u.DB.QueryRow(query, args...).Scan(&user.ID); err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (u *UserDB) GetUser(name string) (*models.User, error) {
	user := models.User{}
	query, args, _ := u.psql.Select("*").From("users").Where(sq.Eq{"name": name}).ToSql()
	res, _ := u.DB.Query(query, args...)
	if !res.Next() {
		return nil, errors.New("user not found")
	}
	defer res.Close()
	for res.Next() {
		if err := res.Scan(&user.ID, &user.Name); err != nil {
			return nil, errors.New("unable to get response")
		}
	}
	return &user, nil
}

func (u *UserDB) UpdateUser(user *models.User) (*models.User, error) {
	_, err := u.GetUser(user.Name)
	if err == nil {
		return nil, errors.New("user with name exists")
	}
	query, args, _ := u.psql.Update("").Table("users").Where(sq.Eq{"id": user.ID}).
		SetMap(sq.Eq{"name": user.Name}).Suffix("RETURNING \"name\"").ToSql()
	if err = u.DB.QueryRow(query, args...).Scan(&user.Name); err != nil {
		return nil, errors.New("user does not exist with the ID")
	}
	return user, nil
}

func (u *UserDB) DeleteUser(user *models.User) error {
	_, err := u.GetUser(user.Name)
	if err != nil {
		return errors.New("user with name does not exist")
	}
	query, args, _ := u.psql.Delete("").From("users").Where("name = ?", user.Name).ToSql()
	fmt.Println("delete sql", query, args)
	if err = u.DB.QueryRow(query, args...).Err(); err != nil {
		return errors.New("cannot delete user " + err.Error())
	}
	return nil
}

func (u *UserDB) GetAllUsers() *[]models.User {
	users := []models.User{}
	query, args, _ := u.psql.Select("*").From("users").ToSql()
	res, _ := u.DB.Query(query, args...)
	defer res.Close()
	for res.Next() {
		user := models.User{}
		if err := res.Scan(&user.ID, &user.Name); err != nil {
			return &[]models.User{}
		}
		users = append(users, user)
	}
	return &users
}

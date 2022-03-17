package repository

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type RegistrationMysql struct {
	db *sql.DB
}

func NewRegistrationMysql(db *sql.DB) *RegistrationMysql {
	return &RegistrationMysql{db: db}
}

func (r *RegistrationMysql) RegisterUser(token string, uid int) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "auth_postgres.go",
			"function": "CreateUser",
			"message":  err,
		}).Errorf("error while starting transaction")
		return err
	}

	insertIntoUsersQuery := fmt.Sprintf(`INSERT INTO users(token, uid, pcode, status) 
										VALUES('%s', '%d', "LVD4060", "user")`, token, uid)
	insertIntoSettingsQuery := fmt.Sprintf(`insert into settings(uid, automine, autoinfect, autoinfect_interval, 
												autoinfect_workt, autoinfect_text, tickerinfect, prefix, tprefix)
											VALUES(%d, %d, %d, %d, %d, '%s', %d, '%s', '%s')`, uid, 0, 0, 3600, 0, "заразить р", 0, "р", "д")

	createCommandsBlackListTable := fmt.Sprintf(`CREATE TABLE %d_cbl(str text)`, uid)
	createTrustedTable := fmt.Sprintf(`CREATE TABLE %d_trusted(id int not null primary key auto_increment, usr_id text)`, uid)
	createTimersTable := fmt.Sprintf(`CREATE TABLE %d_timers(id int, workt bigint, tint int, text text, pid bigint)`, uid)
	createTickersTable := fmt.Sprintf(`CREATE TABLE %d_tickers(id int, workt bigint, tint int, text text, pid bigint)`, uid)

	_, err = tx.Exec(insertIntoUsersQuery)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "registration_mysql.go",
			"function": "CreateTables",
			"message":  err,
		}).Errorf("error while execute query")

		tx.Rollback()
		return err
	}

	_, err = tx.Exec(insertIntoSettingsQuery)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "registration_mysql.go",
			"function": "CreateTables",
			"message":  err,
		}).Errorf("error while execute query")

		tx.Rollback()
		return err
	}

	_, err = tx.Exec(createCommandsBlackListTable)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "registration_mysql.go",
			"function": "CreateTables",
			"message":  err,
		}).Errorf("error while execute query")

		tx.Rollback()
		return err
	}

	_, err = tx.Exec(createTrustedTable)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "registration_mysql.go",
			"function": "CreateTables",
			"message":  err,
		}).Errorf("error while execute query")

		tx.Rollback()
		return err
	}

	_, err = tx.Exec(createTimersTable)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "registration_mysql.go",
			"function": "CreateTables",
			"message":  err,
		}).Errorf("error while execute query")

		tx.Rollback()
		return err
	}

	_, err = tx.Exec(createTickersTable)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "registration_mysql.go",
			"function": "CreateTables",
			"message":  err,
		}).Errorf("error while execute query")

		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RegistrationMysql) UpdateUser(token string, uid int) error {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "auth_postgres.go",
			"function": "CreateUser",
			"message":  err,
		}).Errorf("error while starting transaction")
		return err
	}

	updateTokenQuery := fmt.Sprintf(`UPDATE users SET token='%s' WHERE uid='%d'`, token, uid)

	_, err = tx.Exec(updateTokenQuery)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "registration_mysql.go",
			"function": "CreateTables",
			"message":  err,
		}).Errorf("error while execute query")

		tx.Rollback()
		return err
	}

	return tx.Commit()
}

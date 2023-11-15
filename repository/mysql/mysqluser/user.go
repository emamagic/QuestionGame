package mysqluser

import (
	"database/sql"
	"game/domain"
	"game/pkg/richerror"
	"time"
)

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	op := "mysql.IsPhoneNumberUnique"
	row := d.conn.Conn().QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false,
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithMessage(richerror.DBError).
				WithErr(err)
	}
	return false, nil
}

func (d *DB) Register(u domain.User) (domain.User, error) {
	op := "sql.Register"
	res, err := d.conn.Conn().Exec(`insert into users(name, phone_number, password) values(?, ?, ?)`,
		u.Name, u.PhoneNumber, u.HashPassword)
	if err != nil {
		return domain.User{},
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithMessage(richerror.DBError).
				WithErr(err)
	}

	id, err := res.LastInsertId()
	u.ID = uint(id)

	return u, err
}

func (d *DB) GetUserByPhoneNumber(phoneNumber string) (domain.User, bool, error) {
	op := "sql.GetUserByPhoneNumber"
	row := d.conn.Conn().QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, false, nil
		}
		return domain.User{},
			false,
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithMessage(richerror.DBError).
				WithErr(err)
	}
	return user, true, nil
}

func (d *DB) GetUserByID(userID uint) (domain.User, error) {
	op := "mysql.GetUserByID"
	row := d.conn.Conn().QueryRow(`select * from users where id = ?`, userID)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{},
				richerror.New(op).
					WithCode(richerror.CodeNotFound).
					WithMessage(richerror.RecordNotFound)
		}
		return domain.User{},
			richerror.New(op).
				WithCode(richerror.CodeUnexpected).
				WithMessage(richerror.DBError).
				WithErr(err)

	}
	return user, nil
}

func scanUser(row *sql.Row) (domain.User, error) {
	var user domain.User
	var createdAt time.Time

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.HashPassword, &createdAt)
	return user, err
}

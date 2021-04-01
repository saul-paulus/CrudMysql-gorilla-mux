package controller

import (
	"crud-mysql-gorilla-mux/app/config"
	"crud-mysql-gorilla-mux/app/models"
	"fmt"
	"log"
	"time"
)

const nameTable = "mahasiswa"
const layoutFormat_time = "2006-01-02 15:04:05"

func CreateData(mhs models.Mahasiswa) (err error) {
	db := config.InitMysql()
	timeNow := time.Now()

	sqlText := fmt.Sprintf("INSERT INTO %v (nim,nama,semester,created_at,updated_at) VALUES(?,?,?,?,?)", nameTable)

	row, err := db.Prepare(sqlText)
	if err != nil {
		log.Fatal(err)
		return
	}

	resQuery, err := row.Exec(mhs.Nim, mhs.Nama, mhs.Semester, timeNow, timeNow)
	if err != nil {
		log.Fatal(err)
		return
	}
	lastId, err := resQuery.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return
	}

	mhs.ID = int(lastId)
	mhs.CreatedAt = timeNow
	mhs.UpdatedAt = timeNow

	defer db.Close()
	defer row.Close()

	return nil
}

func GetAllData() (mahasiswa []models.Mahasiswa, err error) {
	db := config.InitMysql()

	sqlText := fmt.Sprintf("SELECT * FROM %v ORDER BY id DESC", nameTable)

	resQuery, err := db.Query(sqlText)
	if err != nil {
		log.Fatal(err)
		return
	}

	for resQuery.Next() {
		mhs := models.Mahasiswa{}

		if err = resQuery.Scan(&mhs.ID,
			&mhs.Nim,
			&mhs.Nama,
			&mhs.Semester,
			&mhs.CreatedAt,
			&mhs.UpdatedAt); err != nil {
			return
		}
		mahasiswa = append(mahasiswa, mhs)
	}

	defer resQuery.Close()
	return mahasiswa, nil
}

func GetIdData(id int) (mhs models.Mahasiswa, err error) {
	db := config.InitMysql()

	sqlText := fmt.Sprintf("SELECT * FROM %v WHERE id=?", nameTable)

	resQuery, err := db.Query(sqlText, id)
	if err != nil {
		log.Fatal(err)
		return
	}

	for resQuery.Next() {
		if resQuery.Scan(&mhs.ID, &mhs.Nim, &mhs.Nama, &mhs.Semester, &mhs.CreatedAt, &mhs.UpdatedAt); err != nil {
			log.Fatal(err)
			return
		}
	}
	return mhs, nil
}

func GetUpdateData(mhs models.Mahasiswa) (err error) {
	db := config.InitMysql()
	sqlText := fmt.Sprintf("UPDATE %v SET nim=?,nama=?,semester=?, created_at=?, updated_at=? WHERE id =?", nameTable)

	timeNow := time.Now()

	row, err := db.Prepare(sqlText)
	if err != nil {
		log.Fatal(err)
		return
	}

	resQuery, err := row.Exec(mhs.Nim, mhs.Nama, mhs.Semester, timeNow, timeNow, mhs.ID)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = resQuery.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer row.Close()
	defer db.Close()

	return nil
}

func GetDeleteData(id int) (err error) {
	db := config.InitMysql()
	sqlText := fmt.Sprintf("DELETE FROM %v WHERE id=?", nameTable)

	row, err := db.Prepare(sqlText)
	if err != nil {
		log.Fatal(err)
		return
	}
	resQuery, err := row.Exec(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = resQuery.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	return nil
}

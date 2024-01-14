package respositories

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
)

func CheckRecordExists(tableName string, column string, value string, con *sql.DB) bool {
	var name string
	sqlStatement := "SELECT " + column + " FROM " + tableName + " WHERE " + column + " = ? AND status ='ACTIVE'"
	err := con.QueryRow(sqlStatement, value).Scan(&name)
	fmt.Println(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	if name != "" {
		return true
	}
	return false
}

func GenerateCaptchaCode(n int, stype string) string {
	letters := "123456789ABCDEFGHJKMNPRSTUVWXYZabcdefghkmnpqrstuvwxyz"
	if stype == "NUM" {
		letters = "0123456789"
	} else if stype == "ALPHA" {
		letters = "ABCDEFGHJKMNPRSTUVWXYZabcdefghkmnpqrstuvwxyz"
	}

	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}

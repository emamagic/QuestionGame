package mysqlaccesscontrol

import (
	"game/domain"
	"game/repository/mysql"
	"time"
)

func scanPermission(scanner mysql.Scanner) (domain.Permission, error) {
	var createdAt time.Time
	var p domain.Permission

	err := scanner.Scan(&p.ID, &p.Title, &createdAt)

	return p, err
}
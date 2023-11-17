package mysqlaccesscontrol

import (
	"game/domain"
	"game/pkg/richerror"
	"game/pkg/slice"
	"game/repository/mysql"
	"strings"
	"time"
)

func (d *DB) GetUserPermissionTitles(userID uint, role domain.Role) ([]domain.PermissionTitle, error) {
	const op = "mysql.GetUserPermissionTitles"

	roleACL := make([]domain.AccessControl, 0)

	rows, err := d.conn.Conn().Query(`select * from access_controls where actor_type = ? and actor_id = ?`,
		domain.RoleActorType, role)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
	}

	defer rows.Close()

	for rows.Next() {
		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
		}

		roleACL = append(roleACL, acl)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
	}

	userACL := make([]domain.AccessControl, 0)

	userRows, err := d.conn.Conn().Query(`select * from access_controls where actor_type = ? and actor_id = ?`,
		domain.UserActorType, userID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
	}

	defer userRows.Close()

	for userRows.Next() {
		acl, err := scanAccessControl(userRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
		}

		userACL = append(userACL, acl)
	}

	if err := userRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
	}

	// merge ACLs by permission id
	permissionIDs := make([]uint, 0)
	for _, r := range roleACL {
		if !slice.DoesExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}
	}

	if len(permissionIDs) == 0 {
		return nil, nil
	}

	// select * from permissions where id in (?,?,?,?...)
	args := make([]any, len(permissionIDs))

	for i, id := range permissionIDs {
		args[i] = id
	}

	// warning: this query works if we have one or more permission id
	query := "select * from permissions where id in (?" +
		strings.Repeat(",?", len(permissionIDs)-1) +
		")"

	pRows, err := d.conn.Conn().Query(query, args...)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
	}
	defer pRows.Close()

	permissionTitles := make([]domain.PermissionTitle, 0)

	for pRows.Next() {
		permission, err := scanPermission(pRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
		}

		permissionTitles = append(permissionTitles, permission.Title)
	}

	if err := pRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(richerror.SomethingWentWrong).WithCode(richerror.CodeUnexpected)
	}

	return permissionTitles, nil
}

func scanAccessControl(scanner mysql.Scanner) (domain.AccessControl, error) {
	var createdAt time.Time
	var acl domain.AccessControl

	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)

	return acl, err
}

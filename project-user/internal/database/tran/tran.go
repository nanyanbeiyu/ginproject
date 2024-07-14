/*
@author: NanYan
*/
package tran

import "carrygpc.com/project-user/internal/database"

type Transaction interface {
	Action(func(conn database.DbConn) error) error
}

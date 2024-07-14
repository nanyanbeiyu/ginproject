/*
@author: NanYan
*/
package repo

import (
	"carrygpc.com/project-user/internal/data/organization"
	"carrygpc.com/project-user/internal/database"
	"context"
)

type OrganizationRepo interface {
	SaveOrganization(conn database.DbConn, ctx context.Context, member *organization.Organization) error
}

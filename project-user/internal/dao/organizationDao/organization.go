/*
@author: NanYan
*/
package organizationDao

import (
	"carrygpc.com/project-user/internal/data/organization"
	"carrygpc.com/project-user/internal/database"
	"carrygpc.com/project-user/internal/database/gorms"
	"context"
)

type OrganizationDao struct {
	conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorms.New(),
	}
}

func (o OrganizationDao) SaveOrganization(conn database.DbConn, ctx context.Context, org *organization.Organization) error {
	o.conn = conn.(*gorms.GormConn)
	return o.conn.Tx(ctx).Create(org).Error
}

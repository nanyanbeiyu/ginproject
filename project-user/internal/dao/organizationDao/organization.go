/*
@author: NanYan
*/
package organizationDao

import (
	"carrygpc.com/project-user/internal/data/organization"
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

func (o OrganizationDao) SaveOrganization(ctx context.Context, org *organization.Organization) error {
	return o.conn.Default(ctx).Create(org).Error
}

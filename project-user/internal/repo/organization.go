/*
@author: NanYan
*/
package repo

import (
	"carrygpc.com/project-user/internal/data/organization"
	"context"
)

type OrganizationRepo interface {
	SaveOrganization(ctx context.Context, member *organization.Organization) error
}

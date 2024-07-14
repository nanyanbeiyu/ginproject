/*
@author: NanYan
*/
package repo

import (
	"carrygpc.com/project-user/internal/data/member"
	"carrygpc.com/project-user/internal/database"
	"context"
)

type MemberRepo interface {
	GetMemberEmail(ctx context.Context, email string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	SaveMember(conn database.DbConn, ctx context.Context, member *member.Member) error
	GetMemberByName(ctx context.Context, name string) (bool, error)
}

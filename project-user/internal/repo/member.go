/*
@author: NanYan
*/
package repo

import (
	"carrygpc.com/project-user/internal/data/member"
	"context"
)

type MemberRepo interface {
	GetMemberEmail(ctx context.Context, email string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	SaveMember(ctx context.Context, member *member.Member) error
	GetMemberByName(ctx context.Context, name string) (bool, error)
}

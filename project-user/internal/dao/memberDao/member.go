/*
@author: NanYan
*/
package memberDao

import (
	"carrygpc.com/project-user/internal/data/member"
	"carrygpc.com/project-user/internal/database"
	"carrygpc.com/project-user/internal/database/gorms"
	"context"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m MemberDao) GetMemberEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Default(ctx).Model(&member.Member{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (m MemberDao) GetMemberByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := m.conn.Default(ctx).Model(&member.Member{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

func (m MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Default(ctx).Model(&member.Member{}).Where("mobile = ?", mobile).Count(&count).Error
	return count > 0, err
}

func (m MemberDao) SaveMember(conn database.DbConn, ctx context.Context, member *member.Member) error {
	m.conn = conn.(*gorms.GormConn)
	return m.conn.Tx(ctx).Create(member).Error
}

func NewMemberDao() *MemberDao {
	return &MemberDao{conn: gorms.New()}
}

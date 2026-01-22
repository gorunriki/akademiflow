package users

import (
	"errors"
	"testing"

	serr "github.com/gorunriki/akademiflow/internal/shared/errors"
	"gorm.io/gorm"
)

type fakeRepo struct {
	updateErr error
}

func (f *fakeRepo) UpdateRole(id uint, role string) error {
	return f.updateErr
}

func (f *fakeRepo) FindByID(id uint) (*User, error) {
	return nil, nil
}

func (f *fakeRepo) Create(user *User) error {
	return nil
}

func (f *fakeRepo) ExistsByEmail(email string) (bool, error) {
	return true, nil
}

func (f *fakeRepo) ListUsers(limit int, offset int, keyword string, includeDeleted bool) ([]User, int64, error) {
	return nil, 0, nil
}

func (f *fakeRepo) Delete(id uint) error {
	return f.updateErr
}

func (f *fakeRepo) Restore(id uint) error {
	return nil
}

func (f *fakeRepo) FindByIDIncludingDeleted(id uint) (*User, error) {
	return nil, nil
}

func newTestService(repoErr error) *service {
	repo := &fakeRepo{
		updateErr: repoErr,
	}
	return &service{
		repo: repo,
	}
}

// test role invalid
func TestUpdateUserRole_InvalidRole(t *testing.T) {
	svc := newTestService(nil)
	err := svc.UpdateUserRole(2, "superadmin", 1)

	if !errors.Is(err, serr.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

// test admin self demote
func TestUpdateUserRole_SelfDemote(t *testing.T) {
	svc := newTestService(nil)
	err := svc.UpdateUserRole(1, "user", 1)
	if !errors.Is(err, serr.ErrForbidden) {
		t.Errorf("expected ErrForbidden, got %v", err)
	}
}

// test user not found
func TestUpdateUserRole_NotFound(t *testing.T) {
	svc := newTestService(gorm.ErrRecordNotFound)
	err := svc.UpdateUserRole(2, "admin", 1)
	if !errors.Is(err, serr.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

// test role update success
func TestUpdateUserRole_Success(t *testing.T) {
	svc := newTestService(nil)
	err := svc.UpdateUserRole(2, "admin", 1)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

// test delete user not found
func TestDeleteUser_NotFound(t *testing.T) {
	svc := newTestService(gorm.ErrRecordNotFound)
	err := svc.DeleteUser(2)
	if !errors.Is(err, serr.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

// test delete user success
func TestDeleteUser_Success(t *testing.T) {
	svc := newTestService(nil)
	err := svc.DeleteUser(3)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

package base

import (
	"context"
	"testing"
)

func TestRepo(t *testing.T) {
	r := NewRepo("https://git3.carpcoupon.com/go_pkg/kratos-layout.git")
	if err := r.Clone(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.CopyTo(context.Background(), "/tmp/test_repo", "git3.carpcoupon.com/go_pkg/kratos-layout", nil); err != nil {
		t.Fatal(err)
	}
}

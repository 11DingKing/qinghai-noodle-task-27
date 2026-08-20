package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask27(t *testing.T) {
	r := NewRegistry()
	require.NoError(t, r.SaveLicense(context.Background(), BrandLicense{ID: "l", RegionCodes: []string{"TJ", "JS"}}, 0))
	s := NewService(r, time.Now)
	regions, err := s.LicenseRegions(context.Background(), "l")
	require.NoError(t, err)
	require.Equal(t, []string{"TJ", "JS"}, regions)
}

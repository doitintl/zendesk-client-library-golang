package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganizationMembershipCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	org := randOrg(t, client)
	defer client.DeleteOrganization(*org.ID)

	user := randUser(t, client)
	defer client.DeleteUser(*user.ID)

	// it should create an organization membership (user can belong only to one org)
	orgMembership1 := OrganizationMembership{
		UserID:         user.ID,
		OrganizationID: org.ID,
	}

	membership, err := client.CreateOrganizationMembership(&orgMembership1)
	require.NoError(t, err)
	require.NotNil(t, membership.ID)
	require.Equal(t, *user.ID, *membership.UserID)
	require.Equal(t, *org.ID, *membership.OrganizationID)

	// it should return all organization memberships for specific user
	found, err := client.ListOrganizationMembershipsByUserID(*user.ID)
	require.NoError(t, err)
	require.Len(t, found, 1)
	found1 := isExistingMembership(*membership.UserID, *membership.OrganizationID, found)
	require.Equal(t, found1, true)

	// it should delete an organization membership
	err = client.DeleteOrganizationMembershipByID(*membership.ID)
	require.NoError(t, err)
	found, err = client.ListOrganizationMembershipsByUserID(*user.ID)
	require.NoError(t, err)
	require.Len(t, found, 0)
}

func isExistingMembership(userId, orgId int64, memberships []OrganizationMembership) bool {
	if memberships != nil {
		for _, membership := range memberships {
			if *membership.OrganizationID == orgId && *membership.UserID == userId {
				return true
			}
		}
	}

	return false
}

package major

import "testing"


func TestShouldReturnUserId(t *testing.T){
	var expecteduserID int = 0;
	username := "root"
	retunedduserID := userIdlookup(username)

	if  retunedduserID!= expecteduserID{
		t.Errorf(" Expected %d but get %d", expecteduserID, retunedduserID )
	}
}
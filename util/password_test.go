package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashPassword,err := HashPassword(password)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	err =  CheckPassword(password,hashPassword)
	require.NoError(t,err)
	
}
package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_GetAccessToken(t *testing.T) {
	logrus.Info("try to get access token")

	accessToken := GetAccessToken()
	logrus.Info(accessToken)
}

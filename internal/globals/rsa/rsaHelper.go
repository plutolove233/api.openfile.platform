// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 1:48
// @Software: GoLand

package rsa

import (
	"api.openfileplatform.com/internal/utils/rsa"
	"github.com/spf13/viper"
)

func GetRSAHelper() rsa.RSA {
	return rsa.RSA{
		PublicKeyPath:  viper.GetString("system.RSAPublic"),
		PrivateKeyPath: viper.GetString("system.RSAPrivate"),
	}
}

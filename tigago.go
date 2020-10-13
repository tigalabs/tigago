//	Package tigago is an integration package for the framework.
//
//	包 tigago 包是框架的集成包。
package tigago

import "github.com/tigateam/tigago/modules/net/tghttp"

// NewServer 新建服务
func NewServer() *tghttp.Server {
	return tghttp.NewServer()
}

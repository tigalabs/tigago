package tigago

import "fmt"

// Framework basic information
//
// 框架基本信息
const (
	Version = "v0.0.7"
	Authors = "TigaTeam"
	Email   = "misitebao@tigateam.com"
	Site    = "https://tigago.tigateam.com"
	Repo    = "https://github.com/tigateam/tigago"
)

//	Info Print frame basic information
//
//	Info 打印框架基本信息
func Info() {
	fmt.Printf(`
  Version:%v
  Authors:%v
  Email  :%v
  Site   :%v
  Repo   :%v
  `, Version, Authors, Email, Site, Repo)
}

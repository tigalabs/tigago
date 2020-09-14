package tigago

import "fmt"

// 框架基本信息
const (
	Version = "v0.0.3"
	Authors = "TigaTeam"
	Email   = "misitebao@tigateam.com"
	Site    = "https://tigago.tigateam.com"
	Repo    = "https://github.com/tigateam/tigago"
)

// Info 输出框架基本信息
func Info() {
	fmt.Printf(`
  Site   :%v
  Repo   :%v
  Email  :%v
  Authors:%v
  Version:%v
  `, Version, Authors, Email, Site, Repo)
}

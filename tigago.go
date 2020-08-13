package tigago

import "fmt"

const (
	VERSION = "v0.0.3"
	AUTHORS = "TigaTeam"
	EMAIL   = "misitebao@tigateam.com"
	SITE    = "https://tigago.tigateam.com"
	REPO    = "https://github.com/tigateam/tigago"
)

// Info 输出框架信息
func Info() {
	fmt.Printf(`
  SITE   :%v
  REPO   :%v
  EMAIL  :%v
  AUTHORS:%v
  VERSION:%v
  `, VERSION, AUTHORS, EMAIL, SITE, REPO)
}

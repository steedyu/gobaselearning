package gousualsample

import (
	"regexp"
	"fmt"
)

var uidmatch = regexp.MustCompile("^\\d{16}$")
func IsAllow(uid string) (ok bool) {

	if uidmatch.MatchString(uid) {
		ok = true
	}

	return
}
//"(?=.*[0-9].*)(?=.*[A-Z].*)(?=.*[a-z].*).{12,20}"
var(
	bidformatregmatch = regexp.MustCompile("^[a-zA-z0-9]{32}$")
	onlynummatch = regexp.MustCompile("^[0-9]{32}$")
	onlylettermatch = regexp.MustCompile("^[a-zA-z]{32}$")
)
func IsBid(bid string) (result bool){
	if bidformatregmatch.MatchString(bid) && !onlylettermatch.MatchString(bid) && !onlynummatch.MatchString(bid) {
		result = true
	}

	return
}

func Xx() {
	//\@=   (?= 顺序环视
	//查找后面是sql的my： /my\(sql\)\@=
	reg :=regexp.MustCompile("(.*[0-9].*)@(.*[A-Z].*)@(.*[a-z].*)@.{8}")

	reg1 := regexp.MustCompile(`my(?=sql)`)
	fmt.Println(reg1.MatchString("mysql"))
	//reg, _ := regexp.Compile("(.*[0-9].*)(.*[a-zA-Z].*){8}")
	fmt.Println(reg.MatchString("111"))
	fmt.Println(reg.MatchString("11z1aa1"))
	fmt.Println(reg.MatchString("z11z1Aa1"))
}
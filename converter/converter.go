package converter 

import(
	"strings"
	"unicode"
)

func StringToArt(s string)(o string){

	if s==""{return}
	for _,r:=range s{if r!='\n'&&!unicode.IsDigit(r){return""}}

	if s=="1"{return strings.Repeat("  |  \n",5)}
	for _,p:=range strings.Split(s,"\n"){

		l:=strings.NewReplacer("1","||","2","|_|").Replace(p)
		o+=strings.Repeat(strings.ReplaceAll(l,"","|")[1:len(l)+1]+"\n",5)
	}
	return
}

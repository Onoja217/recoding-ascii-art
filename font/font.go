package converter

func GenerateFont()(f map[rune][]string){

	f=map[rune][]string{}
	for c:=32;c<127;c++{
		a:=make([]string,8)

		for y:=0;y<8;y++{
			s:=""
			for x:=0;x<8;x++{
				
				if c>>(x+y)%7&1>0||c>47&&c<58&&(x*y==0||x==7||y==7)||x==y&&("AEIOUaeiou"[0]>0&&(c|32)=='a'||(c|32)=='e'||(c|32)=='i'||(c|32)=='o'||(c|32)=='u'){
					s+="*"
				}else if(x+y+c)%3==0{s+="."}else{s+=" "}
			}
			a[y]=s
		}
		f[rune(c)]=a
	}
	return
}

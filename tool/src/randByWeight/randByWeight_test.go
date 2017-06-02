package randByWeight
import(
	loghz"fmt"
	"testing"
//	"strings"
)


func TestIndexByWeight(t *testing.T){
	data:=make(map[string]int)
	for i:=0;i<10000;i++{
		p:=DefaultPrizeArray()
		i:=IndexByWeight(p)
		name:=p[i].Name
		if _,ok:=data[name];!ok{
			data[name]=0
		}else{
			data[name]=data[name]+1
		}
	}

	loghz.Println("data:",data)
}

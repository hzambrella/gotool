package randByWeight
import(
	"math/rand"
//	"loghz"
	"time"
)

func init(){
	rand.Seed(time.Now().Unix())
}

type Array interface{
	Len()int
	Swap(i,j int)
	Weight(i int)int
}

func IndexByWeight(arr Array)int{
	if  arr.Len()<=0{
		return -1
	}
//洗牌shuffle
	for i:=arr.Len()-1;i>0;i--{
		j:=rand.Int()%(i+1)
		arr.Swap(i,j)
//		loghz.Println("arr.Swap:",arr)
	}
//权重和sum of weight
	sum:=0
	for j:=0;j<arr.Len();j++{
		sum+=arr.Weight(j)
	}

	r:=rand.Intn(sum)
//	loghz.Println("r",r)

	index:=-1
//这样做的话，权重小的就机会大
	for i:=0;i<arr.Len();i++{
		w:=arr.Weight(i)
		if w<=0{
			continue
		}
		if r<w{
			index=i
			break
		}
		r-=w
//		loghz.Println("r",r)
	}
	if index<0{
		index=rand.Intn(arr.Len())
	}
	return index
}

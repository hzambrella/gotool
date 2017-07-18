package randByWeight
import (
//	"loghz"
)

type Prize struct{
	Id int
	Name string
	Weight int
}

type PrizeArray []Prize

func DefaultPrizeArray()PrizeArray{
	p1:=Prize{1,"iphone",1}
	p2:=Prize{2,"miphone",10}
	p3:=Prize{3,"huawei",30}
	p4:=Prize{4,"nuojiya",20}
	p5:=Prize{5,"mac_watch",10}
	p6:=Prize{6,"none",50}

	p:=make(PrizeArray,0)
	p=append(p,p1,p2,p3,p4,p5,p6)
	//loghz.Println(p)
	return p
}

func (p PrizeArray)Len()int{
	return len(p)
}


func (p PrizeArray)Swap(i,j int){
	p[i],p[j]=p[j],p[i]
}

func (p PrizeArray)Weight(index int)int{
	return p[index].Weight
}

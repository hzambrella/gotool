package loghz
import(
	"testing"
	"errors"
//	"fmt"
)

func TestPrintln(t  *testing.T){
	Error(errors.New("test err"))
	Println("1","2")
}

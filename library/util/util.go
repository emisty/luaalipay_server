package util

import (
	// "encoding/json"

	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"math"
	// "os"

	//lrd "shbao/util/library/redis"
	"strconv"
	"strings"
	"time"
)

func Tip(c *gin.Context, mcode int, tip string, data interface{}) {

	c.JSON(200, gin.H{"code": mcode, "tip": tip, "data": data})
	c.Abort()
}

func TipC(c *gin.Context, stime, mcode int, tip string, data interface{}) {

	c.JSON(200, gin.H{"stime": stime, "code": mcode, "tip": tip, "data": data})
	c.Abort()
}

func TipGG(c *gin.Context, data string) {

	c.JSON(200, data)
	c.Abort()
}

func Tipb(c *gin.Context, code int, mcode int, tip string, data interface{}) {

	c.JSON(code, gin.H{"code": mcode, "tip": tip, "data": data})
	c.Abort()
}

func GetUID() string {
	id, _ := uuid.NewV4()
	uid := id.String()
	return strings.Replace(uid, "-", "", -1)
}

//余额流水号   //年月日
func GetYuThirdId() string {
	now := time.Now().Local().Format("2006-01-02150405") //当前时间 年-月-日時分秒
	unix := strconv.Itoa(int(time.Now().Unix()))
	uid := now + string([]rune(unix)[3:])
	return strings.Replace(uid, "-", "", -1)
}

func CheckRequest(data []string) {
	switch data[0] {
	case "get":

	case "post":
	default:
	}

}

func GetArrayMin(toChange []float64) float64 {
	a := toChange[0]
	for _, v := range toChange {
		if a > v {
			a = v
		}
	}

	return a
}

func GetMapMin(toChange map[string]float64) float64 {
	a := float64(0)
	for _, v := range toChange {
		if a > v {
			a = v
		}
	}

	return a
}

// //检查请求参数
// public function checkRequest(array $data){

//     switch ($data[0]) {
//         case 'get':
//             $po = $this->request->get();
//             break;
//         case 'post':
//             $po = $this->request->getPost();
//             break;
//         case 'delete':
//             $po = $this->request->get();
//             break;
//         case 'put':
//             $po = $this->request->getPut();
//             break;
//         default:
//             $po = $this->request->get();
//             break;
//     }

//     unset($data[0]);
//     return $this->checkPostIsempty($po,$data);
// }

//  public function checkPostIsempty($po,$need){
//     $poo = $po;

//     $has = [];
//     foreach ($poo as $key => $va) {
//         if(is_array($va)){
//             unset($poo[$key]);
//         }
//         $has[] = $key;
//     }

//     foreach ($need as $key => $va) {
//         if ($va[0] == '?') {
//             $new_va = substr($va, 1);
//             if (!array_key_exists($new_va, $po) || trim($po[$new_va]) == '') {
//                 $po[$new_va] = null;
//                 unset($poo[$new_va]);
//                 // 时间参数特殊处理
//                 $new_va_ = $new_va . '0';
//                 if (isset($poo[$new_va_])) unset($poo[$new_va_]);
//             }
//             continue;
//         }
//         if(!in_array($va, $has)){
//             $this->tip(102,"缺少参数{$va}");exit();
//         }
//     }

//     $t = array_keys(array_map('trim', $poo), '');
//     if($t) {
//         $this->tip(103,"提交参数存在不能有空值",$t);
//     }

//     $this->safeParams($po);

//     return $po;
// }

//一维数组去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func GetMonthDay(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
			//			fmt.Fprintln(os.Stdout, "The month has 31 days")
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	//	fmt.Fprintf(os.Stdout, "The %d-%d has %d days.\n", year, month, days)
	return
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//提取20接口
func GetBcode20(barcode string) (string, string) {

	codeweight := strings.Replace(barcode, "20", "", 1)
	code := []byte(codeweight)[0:5]

	weight := []byte(codeweight)[5:10]
	weightb, _ := strconv.ParseFloat(string(weight), 64)
	weightc := strconv.FormatFloat(weightb, 'f', 0, 64)

	return string(code), weightc

}

//提取20接口
func GetBcode20b(barcode, num string) string {

	codeweight := strings.Replace(barcode, num, "", 1)
	code := []byte(codeweight)[0:5]

	return string(code)
}

func CreateBarcodeM(barcode, sweight string) string {
	newbarcode := ""
	if len(barcode) > 5 {
		return ""
	}

	lsw := len(sweight)
	for i := 0; i < 5-lsw; i++ {
		sweight = "0" + sweight

	}
	newbarcode = "21" + barcode + sweight

	//生成校驗位
	sumJ := 0
	sumO := 0
	for i, v := range newbarcode {
		if math.Mod(float64(i), 2) == 0 {
			nn, _ := strconv.Atoi(string(v))
			sumJ += nn
		} else {
			nn, _ := strconv.Atoi(string(v))
			sumO += nn
		}
	}

	sumJO := sumJ*3 + sumO
	last2 := math.Mod(float64(sumJO), 10)

	last1 := 10 - int(last2)
	if last1 == 10 {
		newbarcode = newbarcode + "0"
	} else {
		newbarcode = newbarcode + strconv.Itoa(last1)
	}

	return newbarcode

}

//wei 重量g
func CreateBarcode(barcode, sweight string) string {
	newbarcode := ""

	lsw := len(sweight)
	for i := 0; i < 5-lsw; i++ {
		sweight = "0" + sweight

	}
	newbarcode = "20" + barcode + sweight

	//生成校驗位
	sumJ := 0
	sumO := 0
	for i, v := range newbarcode {
		if math.Mod(float64(i), 2) == 0 {
			nn, _ := strconv.Atoi(string(v))
			sumJ += nn
		} else {
			nn, _ := strconv.Atoi(string(v))
			sumO += nn
		}
	}

	sumJO := sumJ*3 + sumO
	last2 := math.Mod(float64(sumJO), 10)

	last1 := 10 - int(last2)
	if last1 == 10 {
		newbarcode = newbarcode + "0"
	} else {
		newbarcode = newbarcode + strconv.Itoa(last1)
	}

	return newbarcode

}

func SubString(str string, begin, length int) string {
	// fmt.Println("Substring =", str)
	rs := []rune(str)
	lth := len(rs)
	// fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	// fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	return string(rs[begin:end])
}

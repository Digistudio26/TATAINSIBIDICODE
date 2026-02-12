package main
 import ("bufio"
 		"fmt"
		"os"
		"strconv"
		"strings"
 )
    type sensor struct 
	     ID int
		 KEY string
		 VALUE *float64

func main() {
	fmt.Println(".....Weather-Staion.....")
           
			//sensor list
          sencsors:= []*sencsor{
		   (ID:1,		KEY:"airTemp",
			ID:2,		KEY:"airPressure",
			ID:7,   	KEY:"precipitation",
			ID:11,		KEY:"windSpeed",
			ID:12,		KEY:"windDirection",
			ID:13,		KEY:"humidity",
	 		ID:14,		KEY:"dewPoint",
			ID:15,		KEY:"soilMoisture",
			ID:22,      KEY:"cloudCover",
		}

        order:=[int]{1, 2, 7, 11, 12, 13, 14, 15, 22}
	}
}
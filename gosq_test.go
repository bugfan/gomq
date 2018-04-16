package goq
import(
	"testing"	
	"log"
)
func TestSq(t *testing.T) {
	log.Println("zxy")
	test:=New()
	log.Println(test.In(12))
	log.Println(test.In(123))
	log.Println(test.Len())	
	log.Println(test.Out())
	log.Println(test.Len())		
	log.Println(test.Out())	
	log.Println(test.In(121))
	log.Println(test.Out())	
	log.Println(test.Len())	
	log.Println(test.Out())	
	log.Println(test.Out())	
	log.Println(test.In("avvv"))
	log.Println(test.In("121"))
	log.Println(test.Out())	
	log.Println(test.Out())	
	log.Println(test.Out())	
	log.Println(test.Out())	
	
}

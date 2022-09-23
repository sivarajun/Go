package handlers
import(

"net/http"
"log"
"os"
)
type Hello struct {
  l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{i}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		
		http.Error(rw,"Oops",http.StatusBadRequest)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Oops"))
		return
		}
		h.l.Printf("Data %s\n",d)		
		fmt.Fprintf(rw,"Hello %s Connection established",d)
}

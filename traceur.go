package traceur

func Init()                { fmt.Println("draw mode") }
func Forward(step float64) { fmt.Printf("forward %f\n", step) }
func Right()               { fmt.Println("right") }
func Left()                { fmt.Println("left") }
func Say(mess string)      { fmt.Printf("say %s\n", mess) }
func Up()                  { fmt.Println("color off") }
func Down(col string)      { fmt.Printf("color %s\n", col) }
func Color(col string)     { fmt.Printf("color %s\n", col) }
func Pivote(angle int)     { fmt.Printf("right %d\n", angle) }

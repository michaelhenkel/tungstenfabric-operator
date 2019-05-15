// A comment just to push the positions out

package a

<<<<<<< HEAD
=======
import "fmt"

>>>>>>> v0.0.4
type A string //@A

func Stuff() { //@Stuff
	x := 5
	Random2(x) //@godef("dom2", Random2)
	Random()   //@godef("()", Random)
<<<<<<< HEAD
=======

	var err error         //@err
	fmt.Printf("%v", err) //@godef("err", err)
>>>>>>> v0.0.4
}

package major

import (
	"fmt"
	"os/exec"
	"syscall"
)


func Run(args []string) {

	/*we expect at least two argument when run the application, the application binary and one or more argument
	Ex: ./major cat myfile.txt
	*/
	if len(args)< 2{
		panic("Not enough argument provided, -h for help")
	}


	/* To be able to work with file with root previlegy we must set the effective user to previleed user (root)
	   The Seteuid becomes inrrelevant when the application file owner uid passed to exec syscall is set to root
	*/
	if err :=syscall.Seteuid(0); err!=nil{
		panic(err.Error())
	}

	binary := exec.Command(args[1], args[2:]...);

	//to get the content sent to stdout by the binary in the exec syscall
	out, errr:= binary.Output();

	if errr != nil{
		switch e:= errr.(type){

		case *exec.ExitError:
			panic(string(e.Stderr))

		case *exec.Error:
			panic(e.Error())
		}
	}

	fmt.Println(string(out))
}
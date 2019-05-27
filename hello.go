package main //evety package needs a package main

import ( //import diferent packages
	"errors" //função propria do go para retornar erros
	"fmt"    //input and output
	"math"
)

//main function
func main() {
	//output print mensagem
	fmt.Println("coe rapaziada")

	//declarando variáveis
	fmt.Println("\ndeclarando variáveis")
	var x int //declaração sem inferêrencia
	x = 5     //referenciando um valor
	y := 4    //declaração com inferêrencia e referenciando o valor diretamente
	var sum int = x + y
	fmt.Println(sum)

	//array
	fmt.Println("\narray")
	var a [5]int                  //array com tamanho fixo
	a[2] = 7                      //referencia a partir da casa do item
	b := [6]int{5, 6, 4, 7, 9, 8} //declaração com inferêrencia
	fmt.Println(b)

	//slices
	fmt.Println("\nslices")
	s := []int{1, 2, 3} // slices tem a largura variada
	fmt.Println("eu é 3 elemento", s)
	s = append(s, 13) //adiciona um valor ao fim da lista
	fmt.Println("eu é 4 agora", s)

	//statements
	fmt.Println("\nstatements")
	if x > 6 { //if simples, sem ()
		fmt.Println("x maior que 6")
	} else if x < 2 {
		fmt.Println("x menor que 2")
	} else {
		fmt.Println("x ta no meio otário")
	}

	//maps
	fmt.Println("\nmaps")
	vertices := make(map[string]int)

	vertices["triangulo"] = 3
	vertices["quadrado"] = 4
	vertices["bola"] = 12

	fmt.Println(vertices)
	fmt.Println(vertices["triangulo"])

	delete(vertices, "quadrado")

	fmt.Println(vertices)

	//loops
	fmt.Println("\nloops")
	for i := 0; i < 5; i++ { //so existo o tipo for de loop no go
		fmt.Println("e lá vamos nós...", i)
	}
	//podemos modifica-lo para se comportar como um while
	i := 0
	for i < 5 {
		fmt.Println("e lá vamos nós... de novo", i)
		i++
	}

	//podemos modifica-lo para se comportar como um foreach
	arr := [3]string{"a", "b", "c"}
	for casa, value := range arr {
		fmt.Println("index:", casa, "value:", value)
	}

	//foreach loops podem iterar sobre um map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	//funções
	fmt.Println("\nfunções")
	resultado := soma(2, 3)
	fmt.Println("eu sabe somar otário: ", resultado)

	resultado2, err := raiz(16) //go não fornece suporte a expeptions, esse metodo substitui isso
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado2)
	}

	//structs
	fmt.Println("\nstructs")
	p := pessoa{nome: "Luiz Boça", idade: 32}
	fmt.Println(p)
	fmt.Println(p.idade)

	//ponteiros
	fmt.Println("\nponteiros")
	pont := 7
	fmt.Println("meu lugar na memória:", &pont)

	inc1(i)
	fmt.Println("pont não foi modificado:", pont)

	inc2(&i)
	fmt.Println("pont modificado com sucesso:", pont)

}

func soma(x int, y int) int { //função com retorno int
	return x + y
}

func raiz(x float64) (float64, error) { //função de retorno multiplo
	if x < 0 {
		return 0, errors.New("ai vagabundo... n tem raiz de número negativo!")
	}

	return math.Sqrt(x), nil //nil é o valor zerado de pointers, interfaces, maps, slices, channels e functions
	//nil é quase um null
}

//structs
type pessoa struct {
	nome  string
	idade int
}

//ponteiros
func inc1(x int) { //n modifica nada meu, puta função injusta meu
	x++
}

func inc2(x *int) { //modifica o valor no ponteiro
	*x++
}

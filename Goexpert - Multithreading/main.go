package main

import (
	"os"
	"sync"

	"findceps/handlers"
	argparser "findceps/utils/argsparser"
)

func main() {
	ceps, invalidCeps, outputType := argparser.ParseArgs(os.Args)

	argparser.OutputInvalidCeps(invalidCeps, outputType)

	var wg sync.WaitGroup

	for _, cep := range ceps {
		wg.Add(1)
		go handlers.ProcessCep(cep, outputType, &wg)
	}

	wg.Wait()
}

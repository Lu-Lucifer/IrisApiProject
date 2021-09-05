package validate

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

func ValidRequest(err interface{}) []string {
	var wg sync.WaitGroup
	var errs []string
	if err == nil {
		return errs
	}
	if validateErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validateErrs {
			wg.Add(1)
			go func() {
				sErr := fmt.Errorf("%s 参数错误: %v", e.Namespace(), e.Value())
				errs = append(errs, sErr.Error())

				wg.Done()
			}()
		}
		wg.Wait()
	}
	return errs
}

package test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		fmt.Println("error1111:", &err)
		result, err := tryTheThing()
		fmt.Println("error22222:", &err)
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	fmt.Println("errrrrrrrrrrrrrr",&err)
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func TestParam(t *testing.T) {
	fmt.Println(DoTheThing(true))
	//fmt.Println(DoTheThing(false))


}

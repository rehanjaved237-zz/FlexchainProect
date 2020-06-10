package PersInfo

import (
  "fmt"
)

type PersInfo struct {
  Name string
  Gender string
  DOB string
  CNIC string
  Email string
  BloodGroup string
  Nationality string
  MobileNo string
}

func (a* PersInfo) PersInfoInput() {
  fmt.Println("Enter Your Name:")
  fmt.Scanln(&a.Name)
  fmt.Println("Enter Your Gender:")
  fmt.Scanln(&a.Gender)
}

package FamInfo

type FamInfo struct {
  FamSlice []Member
}

type Member struct {
  Relation string
  Name string
  CNIC string
}

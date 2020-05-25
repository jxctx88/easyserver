package util

import "github.com/satori/go.uuid"
func GenUUID() string {
	// 创建
	u1 := uuid.NewV4()
	return u1.String()
	// 解析
	//u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	//if err != nil {
	//	fmt.Printf("Something gone wrong: %s", err)
	//	return u2
	//}
	//fmt.Printf("Successfully parsed: %s", u2)
}
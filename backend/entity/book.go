package entity

import()

type Book struct {
	ID string `valid:"required~NoID , matches(^[B]\\d{7}$)~กรุณาใส่รหัสนักศึกษา"`
	BooKTitle string `valid:"required~NoName , stringlength(4|20)~กรอกแน่"`
	BookDetail string `valid:"required~NoDetail"`
	BookTotal int `valid:"required~NOTotal ,range(5|100)~เพิ่มน้อยอย่าเพิ่ม"`
	PublisherName string `valid:"required~NoPub"`
	Email string `valid:"required~NOemail,email~ใส่แน่จ้า"`
	Url string 	`valid:"required~NoUrl ,url~ใส่ผิดหญิง"`
}
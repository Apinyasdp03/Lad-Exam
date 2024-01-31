package unit_test

import (
	"testing"

	"github.com/Apinyasdp03/Lad-Exam/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBook(t *testing.T){

	g := NewGomegaWithT(t)

	t.Run(`check ID`, func(t *testing.T) {
		book := entity.Book{
			ID:"",//ผิด
			BooKTitle:"MMmmmmmm",
			BookDetail:"BBBMMMAAAA",
			BookTotal:10, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"http://lab.com",
		}
		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NoID"))
	})

	t.Run(`check booKtitle`, func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"",
			BookDetail:"BBBMMMAAAA",
			BookTotal:10, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"http://lab.com",
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NoName"))
	})

	t.Run(`check booktitle`, func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"mm",
			BookDetail:"BBBMMMAAAA",
			BookTotal:10, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"http://lab.com",
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("กรอกแน่"))
	})

	t.Run(`check bookdetail`, func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"mmmmmmmmmmm",
			BookDetail:"",
			BookTotal:10, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"http://lab.com",
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NoDetail"))

	})

	t.Run(`check booktotal`, func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"mmmmmmmmmmm",
			BookDetail:"mmmmmmmmmmmmmmm",
			BookTotal:0, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"http://lab.com",
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("NOTotal"))
	})

	t.Run(`check bookto` , func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"mmmmmmmmmmm",
			BookDetail:"mmmmmmmmmmmmmmm",
			BookTotal:1, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"http://lab.com",
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("เพิ่มน้อยอย่าเพิ่ม"))

	})

	t.Run(`check email`, func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"mmmmmmmmmmm",
			BookDetail:"mmmmmmmmmmmmmmm",
			BookTotal:6, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mmcom", 
			Url:"http://lab.com",
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("ใส่แน่จ้า"))
	})
	
	t.Run(`check url` , func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"mmmmmmmmmmm",
			BookDetail:"mmmmmmmmmmmmmmm",
			BookTotal:6, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"labjcjhcg",
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("ใส่ผิดหญิง"))
	})

	t.Run(`check bookall` , func(t *testing.T){
		book := entity.Book{
			ID:"B6400484",
			BooKTitle:"MMmmmmmm",
			BookDetail:"BBBMMMAAAA",
			BookTotal:10, 
			PublisherName:"MMMMOMOMOM", 
			Email:"mm@gmail.com", 
			Url:"http://lab.com",		
		}

		ok,err := govalidator.ValidateStruct(book)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

}

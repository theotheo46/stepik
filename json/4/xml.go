package main

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"
)

// начало решения

type Employee struct {
	Id     int    `xml:"id,attr"`
	Name   string `xml:"name"`
	City   string `xml:"city"`
	Salary int    `xml:"salary"`
}

type Department struct {
	Code      string     `xml:"code"`
	Employees []Employee `xml:"employees>employee"`
}

type Organization struct {
	Departments []Department `xml:"department"`
}

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	var organization Organization
	buf := make([]byte, 4000)
	inXML.Read(buf)
	err := xml.Unmarshal(buf, &organization)
	if err != nil {
		return err
	}
	w := csv.NewWriter(outCSV)
	//fmt.Println(organization.Departments[0].Employees[1])
	var s = []string{"id", "name", "city", "department", "salary"}
	w.Write(s)
	for _, dep := range organization.Departments {
		emps := dep.Employees
		for _, emp := range emps {
			var s []string
			s = append(s, strconv.Itoa(emp.Id))
			s = append(s, emp.Name)
			s = append(s, emp.City)
			s = append(s, dep.Code)
			s = append(s, strconv.Itoa(emp.Salary))
			w.Write(s)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil { // (2)
		return err
	}
	return nil
}

// конец решения

func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
        </employees>
    </department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	ConvertEmployees(out, in)
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}

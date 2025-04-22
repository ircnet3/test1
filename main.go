package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type student struct {
	Fio           string
	Age           int
	Birthday      string
	Payment       string
	OrderNumber   string
	OrderDate     string
	Course        string
	Snils         string
	Status        string
	Specialty     string
	Academic_year string
	Material      string
	Count         uint
}

type Inventory struct {
	Material string
	Count    uint
}

var studentslist map[string]student
var studentslist_by_ID map[string]student

// новый импорт
// новый импорт

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа

func home1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой заметки..."))
}

func main() {
	studentslist = make(map[string]student)
	var cols map[string]int = make(map[string]int)

	in := "D:\\Списки для справки.csv"

	file, err1 := os.Open(in)

	if err1 != nil {
		log.Fatal("Error while reading the file", err1)
	}

	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file

	r := csv.NewReader(file)
	r.Comma = ';'
	// records, err := reader.ReadAll()
	names1, err := r.Read()

	if err != nil {
		log.Fatal(err)
	}

	for index, value := range names1 {
		cols[value] = index
	}

	fmt.Println(cols)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		m := student{}
		m.Fio = record[cols["ФИО"]]
		m.Birthday = record[cols["Дата рождения"]]
		m.Age = 100
		m.OrderDate = record[cols["Канцелярская дата приказа на зачисления"]]
		m.OrderNumber = record[cols["Канцелярский номер приказа на зачисления"]]
		m.Payment = record[cols["Источник финансирования"]]
		m.Academic_year = record[cols["Учебный год"]]
		m.Course = record[cols["Курс"]]
		m.Snils = record[cols["СНИЛС"]]
		m.Specialty = record[cols["Специальность/направление"]]
		m.Status = record[cols["Состояние"]]

		studentslist[record[cols["ФИО"]]] = m
		//studentslist_by_ID[record[cols["СНИЛС"]]] = m

		fmt.Println(record)
	}
	fmt.Println(studentslist)
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера.
	// Мы передаем два параметра: TCP-адрес сети для прослушивания (в данном случае это "localhost:4000")
	// и созданный рутер. Если вызов http.ListenAndServe() возвращает ошибку
	// мы используем функцию log.Fatal() для логирования ошибок. Обратите внимание
	// что любая ошибка, возвращаемая от http.ListenAndServe(), всегда non-nil.
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	errx := http.ListenAndServe(":4000", mux)
	log.Fatal(errx)

	// Обработчик главной страницы.

	fmt.Println("Hello, World!")
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles("./html/homepage2.tmpl")
	//ts, err := template.ParseFiles("./html/ext.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	e := studentslist["Яхьяев Магомед Гюльмагомедович"]
	fmt.Println(e.Snils)
	fmt.Println("123-321")
	err = ts.Execute(w, e)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// sweaters := Inventory{"wool", 17}
	// e.Count = 7
	// e.Material = sweaters.Material
	// e.Fio = "1111111111111111111"
	// tmpl, err := template.New("test").Parse("{{.}}  {{.Count}} items are made of {{.Fio}}")
	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpl.Execute(w, e)
	// if err != nil {
	// 	panic(err)
	// }
}

func no_ops() {

}

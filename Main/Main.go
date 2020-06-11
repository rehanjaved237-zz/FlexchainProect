package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	//  "net"
	"html/template"
	"net/http"

	//  a1 "../Account"
	at1 "../Attendance"
	b1 "../Block"
	buff1 "../BlockBuffer"
	c1 "../Blockchain"
	cour1 "../Course"
	n1 "../Network"
	m1 "../Marks"
	p1 "../PersInfo"
	std1 "../Student"
	w1 "../Web"
	g1 "../Grades"
)
// IMPORTANT PROBLEMS TO SOLVE

// Correct GenerateBlockHash function in ../Block
// Remove blk.PrevHash = ""
// Blockchain printing issue

var ownAddr = flag.String("addr", ":10000", "Own Address")
var defaultPeer = flag.String("dpeer", ":10000", "Default Peer")
var frontEnd = flag.String("fend", ":12000", "Front End")
var fileName = flag.String("filename", "blockchain.json", "File where blockchain data is saved")
var defaultDatabase = flag.String("db", "localhost:8000", "Database is running on this IP")
var userName = flag.String("uname", "Sir Ehtesham", "This the user name of the user who is logging in")
var privateKey = flag.String("pkey", "123456", "This is the private key of the user who is logging in")

var registeredCourses []cour1.Course
var unRegisteredCourses []cour1.Course

var (
	templates = template.Must(
		template.ParseFiles(
			"index.html",
			"myHtml.html",
			"Home.html",
			"Login.html",
			"Teacher.html",
			"Student.html",
			"HOD.html",
			"HODAddStudent.html",
			"HODAddInstructor.html",
			"HODOfferCourse.html",
			"HODMineHistory.html",
			"BlockList.html",
			"TeacherCourses.html",
			"StudentEnrolledCourses.html",
			"StudentEnrollCourse.html",
			"TeacherGradeStudents.html",
			"StudentRegCourse.html",
			"TeacherSelectOption.html",
			"TeacherUploadMarks.html",
			"TeacherUploadGrades.html",
			"TeacherUploadAttendance.html",
			"StudentOfferedCourses.html",
			"StudentRegisteredCourses.html",
			"StudentViewCourse.html",
			"StudentViewAttendance.html",
			"StudentViewMarks.html",
			"StudentViewGrade.html",
		))
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "Login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		OwnAddr   string
		LoginAddr string
		UserName string
	}{
		w1.FrontEnd + "/Home/",
		w1.FrontEnd + "/Login/",
		b1.UserName,
	}

	err := templates.ExecuteTemplate(w, "Home.html", &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func myHtmlHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "myHtml.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HODHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port      string
		OwnAddr   string
		LoginAddr string
		UserName string
	}{
		w1.FrontEnd,
		w1.FrontEnd + "/Home/",
		w1.FrontEnd + "/Login/",
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "HOD.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port      string
		OwnAddr   string
		LoginAddr string
		UserName string
	}{
		w1.FrontEnd,
		w1.FrontEnd + "/Home/",
		w1.FrontEnd + "/Login/",
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "Teacher.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port      string
		OwnAddr   string
		LoginAddr string
		UserName string
	}{
		w1.FrontEnd,
		w1.FrontEnd + "/Home/",
		w1.FrontEnd + "/Login/",
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "Student.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HodAddStudentHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port string
		UserName string
	}{
		w1.FrontEnd,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "HODAddStudent.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HodAddInstructorHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port string
		UserName string
	}{
		w1.FrontEnd,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "HODAddInstructor.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherSelectOptionHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port string
		UserName string
	}{
		w1.FrontEnd,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "TeacherSelectOption.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HodOfferCourseHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port string
		UserName string
	}{
		w1.FrontEnd,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "HODOfferCourse.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HodAddInstructor1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
		r.ParseForm()
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		email := r.FormValue("email")
		specialization := r.FormValue("specialization")

		newStudent := std1.Student{
			Name:       name,
			Phone:      phone,
			Department: specialization,
			Email:      email,
		}

		blk := b1.GenerateBlock("Teacher", newStudent)
		blk.Status = true
		n1.BroadCastBlock(blk)
	}
	http.Redirect(w, r, "/hod/", http.StatusSeeOther)
}

func HodAddStudent1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
		r.ParseForm()

		photo := r.FormValue("Photo")
		rollNo := r.FormValue("RollNo")
		name := r.FormValue("Name")
		fatherName := r.FormValue("FatherName")
		cnic := r.FormValue("CNIC")
		phone := r.FormValue("Phone")
		department := r.FormValue("Department")
		email := r.FormValue("Email")

		fmt.Println("Photo:", photo)
		fmt.Println("Rollno:", rollNo)
		fmt.Println("Name:", name)
		fmt.Println("Father Name:", fatherName)
		fmt.Println("CNIC:", cnic)
		fmt.Println("Phone:", phone)
		fmt.Println("Department:", department)
		fmt.Println("Email:", email)

		newStudent := std1.Student{
			Photo:      photo,
			RollNo:     rollNo,
			Name:       name,
			FatherName: fatherName,
			CNIC:       cnic,
			Phone:      phone,
			Department: department,
			Email:      email,
		}

		blk := b1.GenerateBlock("Student", newStudent)
		blk.Status = true
		n1.BroadCastBlock(blk)
	}
	http.Redirect(w, r, "/hod/", http.StatusSeeOther)
}

func HodOfferCourse1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
		r.ParseForm()

		courseCode := r.FormValue("CourseCode")
		courseName := r.FormValue("CourseName")
		courseDescription := r.FormValue("CourseDescription")
		creditHrs := r.FormValue("CreditHrs")
		semester := r.FormValue("Semester")
		assignedTeacher := r.FormValue("AssignedTeacher")

		fmt.Println("Course Code:", courseCode)
		fmt.Println("Course Name:", courseName)
		fmt.Println("Course Description:", courseDescription)
		fmt.Println("Credit Hours:", creditHrs)
		fmt.Println("Semester:", semester)
		fmt.Println("Assigned Teacher:", assignedTeacher)

		newCourse := cour1.Course{
			CourseCode:        courseCode,
			CourseName:        courseName,
			CourseDescription: courseDescription,
			CreditHours:       creditHrs,
			Semester:          semester,
			AssignedTeacher:   assignedTeacher,
		}

		blk := b1.GenerateBlock("Course", newCourse)
		blk.Status = true
		n1.BroadCastBlock(blk)
	}
	http.Redirect(w, r, "/hod/", http.StatusSeeOther)
}

func StudentRegCourseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentRegCourseHandler Successfully")
	data := struct {
		Port string
		UserName string
	}{
		w1.FrontEnd,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentRegCourse.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//this is a new function created on 3/19/2020
func StudentRegCourse1Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
		r.ParseForm()
		regCourse := r.FormValue("regCourse")
		res1 := strings.Split(regCourse, " ")
		fmt.Println("Result 1: ", res1)
		courseCode := res1[0]
		courseName := res1[1]
		courseDescription := res1[2]
		creditHrs := res1[3]
		semester := res1[4]
		assignedTeacher := res1[5]

		fmt.Println("Course Code:", courseCode)
		fmt.Println("Course Name:", courseName)
		fmt.Println("Course Description:", courseDescription)
		fmt.Println("Credit Hours:", creditHrs)
		fmt.Println("Semester:", semester)
		fmt.Println("Assigned Teacher:", assignedTeacher)

		newCourse := cour1.Course{
			CourseCode:        courseCode,
			CourseName:        courseName,
			CourseDescription: courseDescription,
			CreditHours:       creditHrs,
			Semester:          semester,
			AssignedTeacher:   assignedTeacher,
		}

		blk := b1.GenerateBlock("Course", newCourse)
		n1.BroadCastBlock(blk)
	}
	http.Redirect(w, r, "/student/", http.StatusSeeOther)

}

func HodMineHistoryHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Port    string
		BBuffer buff1.BlockBuffer
		UserName string
	}{
		w1.FrontEnd,
		buff1.BlkBuffer,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "HODMineHistory.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherCoursesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BlockListHandler Run Successfully")
	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Course"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "TeacherCourses.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddToBlockchainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddToBlockchainHandler Run")
	hash := r.URL.Path[len("/addtoblockchain/"):]

	fmt.Println("Hashed Value", hash)

	status, index := buff1.BlkBuffer.FindBlock(hash)
	if status == true {

		_, blk := buff1.BlkBuffer.RemoveBlock(index)
		n1.BroadCastRemoveBlockBuffer(hash)
		blk.Status = true
		n1.BroadCastBlock(blk)
	} else {
		fmt.Println("Block Not Found")
	}

	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Course"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "HOD.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterStudentCourseHandler(w http.ResponseWriter, r *http.Request) {
	courseCode := r.URL.Path[len("/registerstudentcourse/"):]
	fmt.Println()

	allCourseBlks := c1.Chain1.FilterBlockchain("Course")
	course := filterCourseByCode(courseCode, allCourseBlks)
	if (course.CourseCode != "") {
		blk := b1.GenerateBlock("RegCourse", course)
		blk.Status = true
		n1.BroadCastBlock(blk)
	}

	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Course"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentOfferedCourses.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func filterCourseByCode(courseCode string, allBlks []b1.Block) cour1.Course {
	for _, blk := range allBlks {
		content := blk.Content.(cour1.Course)
		if content.CourseCode == courseCode {
			return content
		}
	}
	return cour1.Course{}
}

func RemoveFromBlockBufferHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RemoveFromBlockBuffer Run")
	hash := r.URL.Path[len("/removefromblockbuffer/"):]

	status, index := buff1.BlkBuffer.FindBlock(hash)
	if status == true {
		fmt.Println("Block Found")
		_, _ = buff1.BlkBuffer.RemoveBlock(index)
		n1.BroadCastRemoveBlockBuffer(hash)
	} else {
		fmt.Println("Block Not Found")
	}

	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Course"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "HOD", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherSelectOption0Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teacher Select Option 0 Handler")
	courseCode := r.URL.Path[len("/teacherselectoption0/"):]

	data := struct {
		Port  string
		CourseCode string
		UserName string
	}{
		w1.FrontEnd,
		courseCode,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "TeacherSelectOption.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func BlocklistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BlockListHandler Run Successfully")
	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.SliceBlockchain(),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "BlockList.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StudentEnrollCourseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentEnrollCourseHandler Successfully")
	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Course"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentEnrollCourse.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherUploadGradesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TeacherUploadGradesHandler Executed")
	courseCode := r.URL.Path[len("/teacheruploadgrades/"):]
	data := struct {
		Port  string
		Chain []b1.Block
		CourseCode string
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Student"),
		courseCode,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "TeacherUploadGrades.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherUploadMarksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TeacherUploadMarksHandler Executed")
	courseCode := r.URL.Path[len("/teacheruploadmarks/"):]
	data := struct {
		Port  string
		Chain []b1.Block
		CourseCode string
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Student"),
		courseCode,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "TeacherUploadMarks.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func TeacherUploadAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TeacherUploadAttendanceHandler Executed")
	courseCode := r.URL.Path[len("/teacheruploadattendance/"):]

	// chain1 := c1.Chain1.FilterBlockchain("RegCourse")
	// var owners []string
	// for _, val := range chain1 {
	// 	temp1 := val.Content.(cour1.Course)
	//  	if (courseCode == temp1.CourseCode) {
	//  		owners = append(owners, val.Owner)
	//  	}
	// }
	 chain2 := c1.Chain1.FilterBlockchain("Student")
	// var chain3 []b1.Block
	// for _, val := range chain2 {
	// 	temp2 := val.Content.(std1.Student)
	// 	if (contains(owners, temp2.RollNo) == true) {
	// 		chain3 = append(chain3, val)
	// 	}
	// }

	data := struct {
		Port  string
		Chain []b1.Block
		CourseCode string
		UserName string
	}{
		w1.FrontEnd,
		chain2,
		courseCode,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "TeacherUploadAttendance.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StudentEnrolledCoursesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentEnrolledCoursesHandler Successfully")
	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Course"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentEnrolledCourses.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Yahan tak chala hai")
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
		r.ParseForm()
		role := r.FormValue("role")
		username := r.FormValue("username")
		password := r.FormValue("password")

		var privateKeys []string
		privateKeys = append(privateKeys, "88ba4fd15402e3730715288c1efb4b0f4877c860fefc73074df104b620532c69")
		privateKeys = append(privateKeys, "7049841e760f87a8938cf9162a85455c82bd5949867f66c56d04695131b51db8")
		privateKeys = append(privateKeys, "28764925f935273afaf91b3a77cac7e7ff0c00069a2c58334eef7cbca4d830e7")
		privateKeys = append(privateKeys, "eff2bbd2682078bddd1ddbe14ae2cb2f157dbb5d75cfa2ac2f063ae213697e73")

		if (contains(privateKeys, password) == true) {
			b1.UserName = username
			b1.PrivateKey = password
			if (role == "Student" || role == "student") {
				data := struct {
					Port      string
					OwnAddr   string
					LoginAddr string
					UserName string
				}{
					w1.FrontEnd,
					w1.FrontEnd + "/Home/",
					w1.FrontEnd + "/Login/",
					b1.UserName,
				}
				err := templates.ExecuteTemplate(w, "Student.html", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if (role == "Teacher" || role == "teacher") {
				data := struct {
					Port      string
					OwnAddr   string
					LoginAddr string
					UserName string
				}{
					w1.FrontEnd,
					w1.FrontEnd + "/Home/",
					w1.FrontEnd + "/Login/",
					b1.UserName,
				}
				err := templates.ExecuteTemplate(w, "Teacher.html", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if (role == "HOD" || role == "hod") {
				data := struct {
					Port      string
					OwnAddr   string
					LoginAddr string
					UserName string
				}{
					w1.FrontEnd,
					w1.FrontEnd + "/Home/",
					w1.FrontEnd + "/Login/",
					b1.UserName,
				}
				err := templates.ExecuteTemplate(w, "HOD.html", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}

func TeacherUploadAttendance1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TeacherUploadAttendance1 Run Successfully")
	courseCode := r.URL.Path[len("/teacheruploadattendance1/"):]
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {

		if err := r.ParseForm(); err != nil {
		    // handle error
		}

		var attend at1.Attendance
		attend.CourseCode = courseCode

		for key, values := range r.PostForm {
		    fmt.Println(key, values)
				if (key == "attendance") {
					attend.StudentsAttendanceStatus = values
				} else if (key == "uploaddate") {
					attend.AttendanceDate = values[0]
				} else if (key == "rollno") {
					attend.StudentsRollNo = values
				} else if (key == "name") {
					attend.StudentsName = values
				}
		}

		blk := b1.GenerateBlock("UploadAttendance", attend)
		blk.Status = true
		n1.BroadCastBlock(blk)

		data := struct {
			Port  string
			Chain []b1.Block
			UserName string
		}{
			w1.FrontEnd,
			c1.Chain1.FilterBlockchain("Course"),
			b1.UserName,
		}
		err := templates.ExecuteTemplate(w, "Teacher.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func TeacherUploadGrades1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TeacherUploadGrades1 Run Successfully")
	courseCode := r.URL.Path[len("/teacheruploadgrades1/"):]
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {

		if err := r.ParseForm(); err != nil {
		    // handle error
		}

		var graded g1.Grades
		graded.CourseCode = courseCode

		for key, values := range r.PostForm {
		    if (key == "rollno") {
					graded.StudentsRollNo = values
				} else if (key == "name") {
					graded.StudentsName = values
				} else if (key == "grades") {
					graded.StudentsGrades = values
				}
		}

		blk := b1.GenerateBlock("UploadGrade", graded)
		blk.Owner = "HOD"
		blk.Status = false
		n1.BroadCastBlock(blk)

		data := struct {
			Port  string
			Chain []b1.Block
			UserName string
		}{
			w1.FrontEnd,
			c1.Chain1.FilterBlockchain("Course"),
			b1.UserName,
		}
		err := templates.ExecuteTemplate(w, "Teacher.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func StudentOfferedCoursesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentOfferedCourses Run Successfully")
	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("Course"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentOfferedCourses.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StudentRegisteredCoursesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentRegisteredCourses Run Successfully")
	data := struct {
		Port  string
		Chain []b1.Block
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("RegCourse"),
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentRegisteredCourses.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StudentViewCourse0Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentViewCourse0Handler Run Successfully")
	courseCode := r.URL.Path[len("/studentviewcourse0/"):]
	fmt.Println(courseCode)

	data := struct {
		Port  string
		CourseCode string
		UserName string
	}{
		w1.FrontEnd,
		courseCode,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentViewCourse.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type AttendUpload struct {
	AttendanceDate string
	AttendanceStatus string
}

func StudentViewAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentViewAttendance Run Successfully")
	courseCode := r.URL.Path[len("/studentviewattendance/"):]

	chain := c1.Chain1.FilterBlockchain("UploadAttendance")

	var attend []AttendUpload

	for _, sol := range chain {
		val := sol.Content.(at1.Attendance)
    if (val.CourseCode == courseCode) {
			for i, s := range val.StudentsRollNo {
				if (s == b1.UserName) {
					attend = append(attend, AttendUpload{val.AttendanceDate, val.StudentsAttendanceStatus[i]})
				}
			}
		}
	}

	fmt.Println("Attend")
	fmt.Println(attend)

	data := struct {
		Port  string
		Chain []b1.Block
		CourseCode string
		Attend []AttendUpload
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("UploadAttendance"),
		courseCode,
		attend,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentViewAttendance.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type MarksUpload struct {
	EvalName string
	TotalMarks string
	ObtainedMarks string
	Dated string
}

func StudentViewMarksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentViewMarks Run Successfully")
	courseCode := r.URL.Path[len("/studentviewmarks/"):]

	chain := c1.Chain1.FilterBlockchain("UploadMarks")

	var attend []MarksUpload

	for _, sol := range chain {
		val := sol.Content.(m1.Marks)
    if (val.CourseCode == courseCode) {
			for i, s := range val.StudentsRollNo {
				if (s == b1.UserName) {
					attend = append(attend, MarksUpload{val.Title, val.TotalMarks, val.StudentsMarks[i], val.UploadDate})
				}
			}
		}
	}

	data := struct {
		Port  string
		Chain []b1.Block
		CourseCode string
		Marks []MarksUpload
		UserName string
	}{
		w1.FrontEnd,
		c1.Chain1.FilterBlockchain("RegCourse"),
		courseCode,
		attend,
		b1.UserName,
	}
	err := templates.ExecuteTemplate(w, "StudentViewMarks.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StudentViewGradeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentViewGrade Run Successfully")
	courseCode := r.URL.Path[len("/studentviewgrade/"):]

	var grade string
	chain := c1.Chain1.FilterBlockchain("UploadGrade")

	for _, sol := range chain {
		val := sol.Content.(g1.Grades)
		if (val.CourseCode == courseCode) {
			for i, s := range val.StudentsRollNo {
				if (s == b1.UserName) {
					grade = val.StudentsGrades[i]
				}
			}
		}
	}

	data := struct {
		Port  string
		CourseCode string
		Chain []b1.Block
		Grade string
		UserName string
	}{
		w1.FrontEnd,
		courseCode,
		c1.Chain1.FilterBlockchain("RegCourse"),
		grade,
		b1.UserName,
	}

	err := templates.ExecuteTemplate(w, "StudentViewGrade.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherUploadMarks1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TeacherUploadMarks1 Run Successfully")
	courseCode := r.URL.Path[len("/teacheruploadmarks1/"):]
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {

		if err := r.ParseForm(); err != nil {
		    // handle error
		}

		var marks m1.Marks

		marks.CourseCode = courseCode

		for key, values := range r.PostForm {
		    fmt.Println(key, values)
				if (key == "EntryTitle") {
					marks.Title = values[0]
				} else if (key == "TotalMarks") {
					marks.TotalMarks = values[0]
				} else if (key == "UploadDate") {
					marks.UploadDate = values[0]
				} else if (key == "rollno") {
					marks.StudentsRollNo = values
				} else if (key == "name") {
					marks.StudentsName = values
				} else if (key == "marks") {
					marks.StudentsMarks = values
				}
		}

		blk := b1.GenerateBlock("UploadMarks", marks)
		blk.Status = true
		n1.BroadCastBlock(blk)

		data := struct {
			Port  string
			Chain []b1.Block
			UserName string
		}{
			w1.FrontEnd,
			c1.Chain1.FilterBlockchain("Course"),
			b1.UserName,
		}
		err := templates.ExecuteTemplate(w, "Teacher.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func runHandlers() {

	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/Home/", HomeHandler)
	http.HandleFunc("/Login/", LoginHandler)
	http.HandleFunc("/loginRequest/", loginRequestHandler)
	http.HandleFunc("/myHtml/", myHtmlHandler)
	http.HandleFunc("/teacher/", TeacherHandler)
	http.HandleFunc("/student/", StudentHandler)
	http.HandleFunc("/hod/", HODHandler)
	http.HandleFunc("/hodaddstudent/", HodAddStudentHandler)
	http.HandleFunc("/hodaddstudent1/", HodAddStudent1Handler)
	http.HandleFunc("/hodaddinstructor/", HodAddInstructorHandler)
	http.HandleFunc("/hodoffercourse/", HodOfferCourseHandler)
	http.HandleFunc("/hodoffercourse1/", HodOfferCourse1Handler)
	http.HandleFunc("/hodminehistory/", HodMineHistoryHandler)
	http.HandleFunc("/addtoblockchain/", AddToBlockchainHandler)
	http.HandleFunc("/removefromblockbuffer/", RemoveFromBlockBufferHandler)
	http.HandleFunc("/blocklist/", BlocklistHandler)
	http.HandleFunc("/teachercourses/", TeacherCoursesHandler)
	http.HandleFunc("/studentenrollcourse/", StudentEnrollCourseHandler)
	http.HandleFunc("/studentenrolledcourses/", StudentEnrolledCoursesHandler)
	http.HandleFunc("/studentregcourse/", StudentRegCourseHandler)
	http.HandleFunc("/studentregcourse1/", StudentRegCourse1Handler)
	http.HandleFunc("/teacherselectoption/", TeacherSelectOptionHandler)
	http.HandleFunc("/teacherselectoption0/", TeacherSelectOption0Handler)
	http.HandleFunc("/teacheruploadattendance/", TeacherUploadAttendanceHandler)
	http.HandleFunc("/teacheruploadmarks/", TeacherUploadMarksHandler)
	http.HandleFunc("/teacheruploadgrades/", TeacherUploadGradesHandler)
	http.HandleFunc("/studentregisteredcourses/", StudentRegisteredCoursesHandler)
	http.HandleFunc("/studentofferedcourses/", StudentOfferedCoursesHandler)
	http.HandleFunc("/registerstudentcourse/", RegisterStudentCourseHandler)
	http.HandleFunc("/studentviewcourse0/", StudentViewCourse0Handler)
	http.HandleFunc("/studentviewattendance/", StudentViewAttendanceHandler)
	http.HandleFunc("/studentviewmarks/", StudentViewMarksHandler)
	http.HandleFunc("/studentviewgrade/", StudentViewGradeHandler)
	http.HandleFunc("/teacheruploadattendance1/", TeacherUploadAttendance1Handler)
	http.HandleFunc("/teacheruploadmarks1/", TeacherUploadMarks1Handler)
	http.HandleFunc("/teacheruploadgrades1/", TeacherUploadGrades1Handler)
	http.HandleFunc("/hodaddinstructor1/", HodAddInstructor1Handler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Fatal(http.ListenAndServe(w1.FrontEnd, nil))
}

func main() {
	flag.Parse()

	n1.OwnAddress = *ownAddr
	n1.DefaultPeer = *defaultPeer
	w1.FrontEnd = *frontEnd
	n1.DefaultDatabase = *defaultDatabase
	c1.FileName = *fileName
	b1.UserName =  *userName
	b1.PrivateKey = *privateKey

	fmt.Println("Own Address:", n1.OwnAddress)
	fmt.Println("Default Peer:", n1.DefaultPeer)
	fmt.Println("Front End:", w1.FrontEnd)
	fmt.Println("Default Database:", n1.DefaultDatabase)

	go runHandlers()
	n1.StartServer()

	input := 0
	for {
		fmt.Println("Choose one of the following scenario:")
		fmt.Println("Enter 1 to view add a new node")
		fmt.Println("Enter 2 to view all known nodes")
		fmt.Println("Enter 3 to add a new PersInfo Block")
		fmt.Println("Enter 4 to view PersChain")
		fmt.Println("Enter 5 to view the BlockBuffer")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var ip string
			fmt.Println("Enter the new ip")
			fmt.Scanln(&ip)
			n1.AddToKnownNode(ip)
		case 2:
			n1.PrintKnownNodes()
		case 3:
			var input p1.PersInfo
			input.PersInfoInput()
			blk := b1.GenerateBlock("PersInfo", input)
			//      blk.PrintBlock()
			n1.BroadCastBlock(blk)
		case 4:
			c1.PrintBlockchain(c1.Chain1)
		case 5:
			buff1.BlkBuffer.PrintBlockBuffer()
		default:
			fmt.Println("Invalid Command Entered")
		}
	}

}

/*type Fun1 struct {
  A int
}

type Fun2 struct {
  A string
  B float64
}

func main() {
  blk1 := b1.GenerateBlock("Rehan", Fun1{A: 5})
  blk1.PrintBlock()
  blk2 := b1.GenerateBlock("Rehan", Fun1{A: 6})
  blk2.PrintBlock()
  blk3 := b1.GenerateBlock("Rehan", Fun2{A: "Alhamdulillah", B: 4.7})
  blk3.PrintBlock()
}*/

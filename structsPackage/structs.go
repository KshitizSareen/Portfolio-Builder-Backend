package structsPackage

type User struct {
	ID int `json:"user_id"`
}
type Header struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	LinkedIn    string `json:"LinkedIn"`
	Profile_URL string `json:"Profile_URL"`
	Summary     string `json:"Summary"`
	User_id     string `json:"user_id"`
}

type Course struct {
	ID           int    `json:"ID"`
	CourseName   string `json:"CourseName"`
	Education_ID int    `json:"Education_id"`
}

type Education struct {
	ID        int      `json:"ID"`
	Name      string   `json:"Name"`
	Degree    string   `json:"Degree"`
	Location  string   `json:"Location"`
	Major     string   `json:"Major"`
	StartDate string   `json:"StartDate"`
	EndDate   string   `json:"EndDate"`
	GPA       float32  `json:"GPA"`
	User_id   int      `json:"user_id"`
	Courses   []Course `json:"Courses"`
}

type UserCode struct {
	User_Code string `json:"User_Code"`
}

type Experience struct {
	ID                      int                      `json:"ID"`
	Company_Name            string                   `json:"Company_Name"`
	Position                string                   `json:"Position"`
	Location                string                   `json:"Location"`
	StartDate               string                   `json:"Start_Date"`
	EndDate                 string                   `json:"End_Date"`
	User_id                 int                      `json:"User_ID"`
	Experience_Descriptions []Experience_Description `json:"Experience_Descriptions"`
}

type Experience_Description struct {
	ID            int    `json:"ID"`
	Description   string `json:"Description"`
	Experience_id int    `json:"Experience_id"`
}

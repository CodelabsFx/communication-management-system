package users

type User struct {
    ID        string   `json:"id"`
    Email     string   `json:"email"`
    FirstName string   `json:"first_name"`
    LastName  string   `json:"last_name"`
    CompanyID string   `json:"company_id"`
    Roles     []string `json:"roles"`
}

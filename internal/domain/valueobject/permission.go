package valueobject

type Permission struct {
	Name string
}

var P = struct {
	User struct {
		Read        Permission
		Write       Permission
		Delete      Permission
		Update      Permission
		ChangeRole  Permission
		ChangeSatus Permission
	}
}{
	User: struct {
		Read        Permission
		Write       Permission
		Delete      Permission
		Update      Permission
		ChangeRole  Permission
		ChangeSatus Permission
	}{
		Read:        Permission{Name: "user:read"},
		Write:       Permission{Name: "user:write"},
		Update:      Permission{Name: "user:update"},
		Delete:      Permission{Name: "user:delete"},
		ChangeRole:  Permission{Name: "user:change_role"},
		ChangeSatus: Permission{Name: "user:change_status"},
	},
}

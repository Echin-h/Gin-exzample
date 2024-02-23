package casbin

func (e *MyEnforce) AddBasedPolicies() {
	e.AddPolicy("user", "users", "read")
	e.AddPolicy("user", "users", "write")
}

func (e *MyEnforce) LinkUserWithPolicy(name string) error {
	_, err := e.AddGroupingPolicy(name, "user")
	return err
}

func (e *MyEnforce) UnLinkUserWithPolicy(name string) error {
	_, err := e.RemoveGroupingPolicy(name, "user")
	return err
}

func (e *MyEnforce) CheckUserPolicyForRead(name, data, action string) bool {
	ok, _ := Enforce.Enforce(name, data, action)
	return ok
}

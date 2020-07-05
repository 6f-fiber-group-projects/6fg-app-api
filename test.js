await fetch("/users", {
	method: "POST",
	headers: {
			"content-type": "application/json",
	},
	body: JSON.stringify({
      auth_id: 1,
      google_id: null,
      name: "test",
      email: "test@test.com",
      password: null,
	})
});
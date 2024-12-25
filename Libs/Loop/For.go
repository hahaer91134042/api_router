package Loop

func DoWhile(fun func() bool) {
	for {
		if !fun() {
			break
		}
	}
}

func While(fun func() (bool, func())) {

	for {
		con, run := fun()
		if !con {
			break
		}
		run()
	}
}

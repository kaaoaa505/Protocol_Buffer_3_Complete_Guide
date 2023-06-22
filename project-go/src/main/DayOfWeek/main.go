package main

import "local.dev/m/v2/build/local"

func main() {
	println("\n\n\n################# Day of week - START.")

	message := local.TargetWeekObj{}
	message.Id = 123
	message.DayOfWeek = local.DayOfWeek_FRIDAY

	println(message.DayOfWeek);
	println(message.DayOfWeek.String());

	message.DayOfWeek = local.DayOfWeek_SATURDAY

	println(message.DayOfWeek);
	println(message.DayOfWeek.String());

	println("################# Day of week - END.\n\n\n")
}


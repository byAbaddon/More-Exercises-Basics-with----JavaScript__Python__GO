function numbersDividedBy3WithoutReminder() {
    let loop = 1

    while (loop <= 100) {
        if (loop % 3 == 0) {
            console.log(loop)
        }

        loop++
    }
}

// numbersDividedBy3WithoutReminder()
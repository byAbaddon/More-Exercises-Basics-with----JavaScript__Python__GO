function grades(arg) {
    let students = Number(arg.shift())
    let grades = [...arg.map(parseFloat)]
    console.log(`Top students: ${(grades.filter(el => el >= 5 ).length * 100 / students).toFixed(2)}%`)
    console.log(`Between 4.00 and 4.99: ${(grades.filter(el => el >= 4  && el <= 4.99 ).length * 100 / students).toFixed(2)}%`)
    console.log(`Between 3.00 and 3.99: ${(grades.filter(el => el >= 3 && el <= 3.99 ).length * 100 / students).toFixed(2)}%`)
    console.log(`Fail: ${(grades.filter(el => el <= 2.99).length * 100 / students).toFixed(2)}%`)
    console.log(`Average: ${(grades.filter(el => el < 7).reduce((a, b) => a + b) / students).toFixed(2)}`)   
}

// grades([10, 3.00, 2.99, 5.68, 3.01, 4, 4, 6.00, 4.50, 2.44, 5, ])
// grades([6, 2, 3, 4, 5, 6, 2.2])
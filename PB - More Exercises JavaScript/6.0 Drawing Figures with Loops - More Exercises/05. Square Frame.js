function squareFrame(n) {
    //top
    console.log(`+${' -'.repeat(n - 2)} +`)
    
    //middle
    for (let i = 0; i < n - 2; i++) {
        console.log(`|${' -'.repeat(n - 2)} |`)
    }

    //bottom
    console.log(`+${' -'.repeat(n - 2)} +`)
}

// squareFrame(5)
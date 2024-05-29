class Library {
    books: Book[]

    constructor() {
        this.books = []
    }

    borrowBook(name: string, by: Student) {
        let borrowedBook = null;

        for (const book of this.books) {
            if (name === book.name && !book.borrowedBy) {
                borrowedBook = book
            }
        }

        if (!borrowedBook) {
            console.log("this book already was borrowed")
            return
        }

        borrowedBook.borrowedBy = by
        console.log(`here you book: ${borrowedBook.name}`)
    }

    returnBook(name: string) {
        let borrowedBook = null;

        for (const book of this.books) {
            if (name === book.name && !!book.borrowedBy) {
                borrowedBook = book
            }
        }

        if (!borrowedBook) {
            console.log("this book is not in our lib or was not borrowed")
            return
        }

        borrowedBook.borrowedBy = null
        console.log(`Thank you for return: ${borrowedBook.name}`)
    }
}

class Book {
    name: string
    #author: string
    #year: number
    borrowedBy: Student | null

    constructor(name: string, author: string = "", year: number ) {
        this.name = name;
        this.#author = author;
        this.#year = year;
        this.borrowedBy = null
    }
}

class Student {
    name: string

    constructor(name: string) {
        this.name = name
    }
}

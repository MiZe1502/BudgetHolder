import {Category} from "./types";

export const mockData: Category[] = [
    {
        id: 1,
        name: "Products",
        categories: [
            {
                id: 4,
                name: "Milk"
            },
            {
                id: 5,
                name: "Fish"
            },
            {
                id: 7,
                name: "Snack"
            },
            {
                id: 8,
                name: "Groceries"
            }
        ]
    },
    {
        id: 2,
        name: "Sport"
    },
    {
        id: 3,
        name: "Entertainment",
        categories: [
            {
                id: 6,
                name: "Movies"
            },
            {
                id: 9,
                name: "Games",
                categories: [
                    {
                        id: 17,
                        name: "PS4",
                    },
                    {
                        id: 18,
                        name: "PC"
                    }
                ]
            }
        ]
    },
    {
        id: 10,
        name: "Electronics",
        categories: [
            {
                id: 11,
                name: "TV"
            }
        ]
    },
    {
        id: 12,
        name: "Furniture"
    },
    {
        id: 13,
        name: "Clothes",
        categories: [
            {
                id: 14,
                name: "Children"
            },
            {
                id: 15,
                name: "Women"
            },
            {
                id: 16,
                name: "Men"
            }
        ]
    }
]
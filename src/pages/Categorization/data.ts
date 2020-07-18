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
            }
        ]
    }
]
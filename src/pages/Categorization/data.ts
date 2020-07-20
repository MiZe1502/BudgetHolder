import {Category} from "./types";

export const mockData: Category[] = [
    {
        id: 1,
        name: "Products",
        parentId: null,
        categories: [
            {
                id: 4,
                name: "Milk",
                parentId: 1,
            },
            {
                id: 5,
                name: "Fish",
                parentId: 1,
            },
            {
                id: 7,
                name: "Snack",
                parentId: 1,
            },
            {
                id: 8,
                name: "Groceries",
                parentId: 1,
            }
        ]
    },
    {
        id: 2,
        name: "Sport",
        parentId: null,
    },
    {
        id: 3,
        name: "Entertainment",
        parentId: null,
        categories: [
            {
                id: 6,
                name: "Movies",
                parentId: 3,
            },
            {
                id: 9,
                name: "Games",
                parentId: 3,
                categories: [
                    {
                        id: 17,
                        name: "PS4",
                        parentId: 9,
                    },
                    {
                        id: 18,
                        name: "PC",
                        parentId: 9,
                    }
                ]
            }
        ]
    },
    {
        id: 10,
        name: "Electronics",
        parentId: null,
        categories: [
            {
                id: 11,
                name: "TV",
                parentId: 10,
            }
        ]
    },
    {
        id: 12,
        name: "Furniture",
        parentId: null,
    },
    {
        id: 13,
        name: "Clothes",
        parentId: null,
        categories: [
            {
                id: 14,
                name: "Children",
                parentId: 13,
            },
            {
                id: 15,
                name: "Women",
                parentId: 13,
            },
            {
                id: 16,
                name: "Men",
                parentId: 13,
            }
        ]
    }
]
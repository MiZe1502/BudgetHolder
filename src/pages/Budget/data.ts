import {GoodsData, Purchase} from "./types";
import {SimpleCategory} from "../Categorization/types";

export const mockPurchases: Purchase[] = [
    {
        id: 1,
        totalPrice: 300,
        comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
        shop: {
            id: 1,
            name: "Dicksy",
            url: "https://dicksy.ru",
            address: "Moscow, 2-ya Vladimirskaya, 20",
            comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        },
        date: new Date(),
        goods: [
            {
                id: 1,
                name: "Burger - the most powerful burger in the entire universe with bacon and chips and chocolate",
                category: {
                    id: 1,
                    value: "Products",
                },
                comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
                price: 200,
                amount: 4,
                purchaseId: 1,
            },
            {
                id: 2,
                name: "Milk with juice",
                category: {
                    id: 1,
                    value: "Products",
                },
                comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
                price: 100,
                amount: 1,
                purchaseId: 1,
            },
            {
                id: 4,
                name: "God of war",
                category: {
                    id: 17,
                    value: "PS4",
                },
                comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
                price: 100,
                amount: 1,
                purchaseId: 1,
            },
        ],
    },
    {
        id: 2,
        totalPrice: 100,
        comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        shop: {
            id: 3,
            name: "Karusel",
            url: "https://karusel.ru",
            address: "Moscow, 2-ya Vladimirskaya, 30",
        },
        date: new Date(),
        goods: [
            {
                id: 1,
                name: "Burger",
                category: {
                    id: 1,
                    value: "Products",
                },
                comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
                price: 100,
                amount: 1,
                purchaseId: 2,
            },
        ],
    },
    {
        id: 3,
        totalPrice: 15000,
        comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
        shop: {
            id: 3,
            name: "Karusel",
            url: "https://karusel.ru",
            address: "Moscow, 2-ya Vladimirskaya, 30",
        },
        date: new Date(),
        goods: [
            {
                id: 6,
                name: "Without category",
                category: null,
                comment: "Test with no category",
                price: 10000,
                purchaseId: 3,
            },
            {
                id: 5,
                name: "Table",
                category: {
                    id: 12,
                    value: "Furniture",
                },
                comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
                price: 5000,
                amount: 1,
                purchaseId: 3,
            },
        ],
    },
]

export const mockGoods: GoodsData[] = [
    {
        id: 1,
        name: "Burger",
        category: {
            id: 1,
            value: "Products",
        },
        comment: "Test burger",
    },
    {
        id: 1123,
        name: "Buter",
        category: {
            id: 1,
            value: "Products",
        },
        comment: "Test burger",
    },
    {
        id: 123,
        name: "Busido",
        category: {
            id: 1,
            value: "Products",
        },
        comment: "Test burger",
    },
    {
        id: 1222,
        name: "labuten",
        category: {
            id: 1,
            value: "Products",
        },
        comment: "Test burger",
    },
    {
        id: 121413,
        name: "Burittoger",
        category: {
            id: 1,
            value: "Products",
        },
        comment: "Test burger",
    },
    {
        id: 199,
        name: "ansburg",
        category: {
            id: 1,
            value: "Products",
        },
        comment: "Test burger",
    },
    {
        id: 2,
        name: "Milk",
        category: {
            id: 1,
            value: "Products",
        },
        comment: "Test milk",
    },
    {
        id: 3,
        name: "God of war",
        category: {
            id: 17,
            value: "PS4",
        },
        comment: "Test ps4 game",
    },
    {
        id: 4,
        name: "LG OLED",
        category: {
            id: 11,
            value: "TV",
        },
        comment: "Test tv",
    },
    {
        id: 5,
        name: "Table",
        category: {
            id: 12,
            value: "Furniture",
        },
        comment: "Test table",
    },
    {
        id: 6,
        name: "Without category",
        category: {} as SimpleCategory,
        comment: "Test with no category",
    }
]
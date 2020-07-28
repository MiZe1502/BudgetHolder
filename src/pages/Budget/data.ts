import {GoodsData, Purchase} from "./types";

export const mockPurchases: Purchase[] = [
    {
        id: 1,
        totalPrice: 300,
        comment: "Test purchase 1",
        shop: {
            id: 1,
            name: "Dicksy",
            url: "https://dicksy.ru",
            address: "Moscow, 2-ya Vladimirskaya, 20",
            comment: "Good shop",
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
                comment: "Test burger",
                price: 200,
                amount: 4
            },
            {
                id: 2,
                name: "Milk",
                category: {
                    id: 1,
                    value: "Products",
                },
                comment: "Test milk",
                price: 100,
                amount: 1
            },
        ],
    },
    {
        id: 2,
        totalPrice: 100,
        comment: "Test purchase 2",
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
                comment: "Test burger",
                price: 100,
                amount: 1
            },
        ],
    },
    {
        id: 3,
        totalPrice: 15000,
        comment: "Test purchase 3",
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
                price: 10000
            },
            {
                id: 5,
                name: "Table",
                category: {
                    id: 12,
                    value: "Furniture",
                },
                comment: "Test table",
                price: 5000,
                amount: 1,
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
        category: null,
        comment: "Test with no category",
    }
]
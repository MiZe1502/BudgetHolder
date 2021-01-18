import {UserData} from "./auth";

export type MockedUserData = UserData & {password: string}

export const mockedUsers: MockedUserData[]   =[
    {
        login: "test",
        name: "testName",
        surname: "testSurname",
        password: "123",
    },
    {
        login: "testWithImage",
        name: "testName",
        surname: "testSurname",
        image: "pic.png",
        password: "qwerty"
    }
]
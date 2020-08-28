export default {
    common: {
        messages: {
            errors: {
                fetching: "Error fetching data"
            },
            no_data: {
                not_found: "No data found"
            }
        },
        components: {
            map: {
                title: "Map"
            },
            menu: {
                home: "Home",
                budget: "Budget",
                shops: "Shops",
                statistics: "Statistics",
                categorization: "Categorization"
            },
            buttons: {
                ok: "OK",
                save: "Save",
                cancel: "Cancel"
            }
        },
        labels: {
            name: "Name",
            url: "Url",
            address: "Address",
            comment: "Comment"
        },
        titles: {
            edit: "Edit: ",
            remove: "Remove: ",
            details: "Details",
        }
    },
    categories: {
        buttons: {
            new: "Add new category"
        },
        titles: {
            categories: "Categories",
            new: "New category"
        },
        messages: {
            remove_part1: "Do You really want to remove",
            remove_part2: "? All children categories will be erased too."
        }
    },
    shops: {
        buttons: {
            new: "Add new shop"
        },
        titles: {
            shops: "Shops",
            new: "New shop"
        },
        messages: {
            remove_part1: "Do You really want to remove",
            remove_part2: "?",
        }
    },
    budget: {
        messages: {
            remove_part1: "Do You really want to remove purchase with all connected goods details",
            remove_part2: "?",
        },
        buttons: {
            new: "Add new purchase",
            clear: "Clear all data",
            save: "Save purchase"
        },
        titles: {
            purchases: "Purchases",
            edit: "Edit purchase at: ",
        },
        labels: {
            date: "Date",
            price: "Total price",
            shop: "Shop",
            category: "Category",
            amount: "Amount",
            single_price: "Price",
        },
    },
    goods: {
        messages: {
            remove_part1: "Do You really want to remove goods item from this purchase",
            remove_part2: "?",
        },
        titles: {
            edit: "Edit goods details"
        }
    },
    auth: {
        labels: {
            login: "Login",
            new_password: "New password",
            password: "Password",
            name: "Name",
            surname: "Surname"
        },
        buttons: {
            login: "Login",
            save_and_login: "Save and login",
            register: "Register",
            auth: "Authorization",
        }
    },
    user: {
        menu: {
            settings: "Settings",
            logout: "Logout",
        },
        titles: {
            edit: "Edit user",
        }
    }
}
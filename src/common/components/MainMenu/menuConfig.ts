interface SingleMenuBlock {
    title: string;
    url: string;
    children?: SingleMenuBlock[];
}

export default [
    {
        title: "Home",
        url: "/"
    },
    {
        title: "Budget",
        url: "/budget",
    },
    {
        title: "Shops",
        url: "/shops",
    },
    {
        title: "Statistics",
        url: "/stats",
    }
];
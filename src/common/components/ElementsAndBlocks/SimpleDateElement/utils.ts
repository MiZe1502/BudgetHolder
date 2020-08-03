export const dateToDMY = (date: Date): string => {
    let parsedDate = typeof date === "string" ? new Date(date) : date;
    const d = parsedDate.getDate();
    const m = parsedDate.getMonth() + 1;
    const y = parsedDate.getFullYear();
    return "" + (d <= 9 ? "0" + d : d) + "." + (m<=9 ? "0" + m : m) + "." + y;
}
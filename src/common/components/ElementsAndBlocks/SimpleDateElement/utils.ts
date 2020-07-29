export const dateToDMY = (date: Date): string => {
    const d = date.getDate();
    const m = date.getMonth() + 1;
    const y = date.getFullYear();
    return "" + (d <= 9 ? "0" + d : d) + "." + (m<=9 ? "0" + m : m) + "." + y;
}
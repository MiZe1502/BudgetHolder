const datePattern = "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$";

export const dateToDMY = (date: Date | string): string => {
    if (!date) {
        return "";
    }

    if (typeof date === "string") {
        if ((new RegExp(datePattern)).test(date)) {
            const dateParts = date.split("-");
            return "" + dateParts[2] + "." + dateParts[1] + "." + dateParts[0];
        } else {
            return date;
        }
    }

    let parsedDate = typeof date === "string" ? new Date(date) : date;

    const d = parsedDate.getDate();
    const m = parsedDate.getMonth() + 1;
    const y = parsedDate.getFullYear();
    return "" + (d <= 9 ? "0" + d : d) + "." + (m<=9 ? "0" + m : m) + "." + y;
}
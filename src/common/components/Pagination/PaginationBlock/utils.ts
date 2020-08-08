
    export const middleRange = 5;
    export const maxRecordsPerPage = 10;

    export enum Delimeters {
        Left= "Left",
        Right= "Right",
    }

    export const formPagesArray = (totalPages: number, currentPage: number): number[] => {
        const pagesInMiddle = [];
        if (totalPages > 2) {
            const anchorPage = Math.max(2, Math.min(currentPage - 2, totalPages - middleRange));
            if (anchorPage > 2) {
                pagesInMiddle.push(Delimeters.Left);
            }
            for (let i = 0; i <= middleRange - 1; i += 1) {
                const page = anchorPage + i;
                if (page < totalPages) {
                    pagesInMiddle.push(page);
                }
            }
            if (anchorPage + middleRange < totalPages) {
                pagesInMiddle.push(Delimeters.Right);
            }
        }

        return pagesInMiddle;
    }
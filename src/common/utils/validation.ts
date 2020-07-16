export interface ValidationRule {
    fieldName: string;
    checkLength?: boolean;
    length?: number;
    checkRegexp?: boolean;
    regexp?: RegExp;
    validator: (data: Record<string, any>) => boolean;
    message: string;
}

export const notInvalidByDefault = false;

export const validURL = (str: string): boolean => {
    const pattern = new RegExp('^(https?:\\/\\/)?' + // protocol
        '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|' + // domain name
        '((\\d{1,3}\\.){3}\\d{1,3}))' + // OR ip (v4) address
        '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*' + // port and path
        '(\\?[;&a-z\\d%_.~+=-]*)?' + // query string
        '(\\#[-a-z\\d_]*)?$', 'i'); // fragment locator
    return !!pattern.test(str);
}

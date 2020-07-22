export interface Category {
    id: number;
    name: string;
    categories?: Category[];
    parentId: number | null;
}

export interface SimpleCategory {
    id: number;
    value: string;
}
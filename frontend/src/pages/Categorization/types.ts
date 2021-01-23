export interface Category {
    id: number;
    name: string;
    categories?: Category[];
    parent_id: number | null;
}

export interface SimpleDataItem {
    id: number;
    value?: string;
}
